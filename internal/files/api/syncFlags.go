package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/flagship-io/code-analyzer/internal/files/model"
)

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type FlagUsageRequest struct {
	RepositoryURL    string `json:"repositoryUrl"`
	RepositoryBranch string `json:"repositoryBranch"`
	Flags            []Flag `json:"flags"`
}

// Flag represent a flag code info sent to the Flagship API
type Flag struct {
	FlagKey           string `json:"flagKey"`
	FilePath          string `json:"filePath"`
	LineNumber        int    `json:"lineNumber"`
	Code              string `json:"code"`
	CodeLineHighlight int    `json:"codeLineHighlight"`
}

// SendFlagsToAPI takes file search result & sends flag info to the API
func SendFlagsToAPI(results []model.FileSearchResult, envId string) (err error) {
	flagUsageRequest := FlagUsageRequest{
		RepositoryURL:    os.Getenv("REPOSITORY_URL"),
		RepositoryBranch: os.Getenv("REPOSITORY_BRANCH"),
	}
	var flags []Flag
	for _, fr := range results {
		for _, r := range fr.Results {
			flags = append(flags, Flag{
				FlagKey:           r.FlagKey,
				FilePath:          fr.File,
				LineNumber:        r.LineNumber,
				Code:              r.CodeLines,
				CodeLineHighlight: r.CodeLineHighlight,
			})
		}
	}
	flagUsageRequest.Flags = flags

	err = callAPI(envId, flagUsageRequest)

	return err
}

func generateAuthenticationToken() error {
	if os.Getenv("FLAGSHIP_CLIENT_ID") == "" {
		log.WithFields(log.Fields{"variable": "FLAGSHIP_CLIENT_ID"}).Fatal("Missing required environment variable")
	}

	if os.Getenv("FLAGSHIP_CLIENT_SECRET") == "" {
		log.WithFields(log.Fields{"variable": "FLAGSHIP_CLIENT_SECRET"}).Fatal("Missing required environment variable")
	}

	if os.Getenv("ACCOUNT_ID") == "" {
		log.WithFields(log.Fields{"variable": "ACCOUNT_ID"}).Fatal("Missing required environment variable")
	}

	if os.Getenv("FS_API") == "" {
		log.WithFields(log.Fields{"variable": "FS_API"}).Fatal("Missing required environment variable")
	}

	authRequest := AuthRequest{
		GrantType:    "client_credentials",
		Scope:        "*",
		ClientId:     os.Getenv("FLAGSHIP_CLIENT_ID"),
		ClientSecret: os.Getenv("FLAGSHIP_CLIENT_SECRET"),
	}

	body, err := json.Marshal(authRequest)

	if err != nil {
		log.Fatal("Error while marshal json", err.Error())
	}

	route := fmt.Sprintf("%s/%s/token?expires_in=0", os.Getenv("FS_AUTH_API"), os.Getenv("ACCOUNT_ID"))

	req, err := http.NewRequest("POST", route, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error in request", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	log.WithFields(log.Fields{
		"method": "POST",
		"route":  route,
	}).Info("Calling Flagship authentication api")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		var result AuthResponse
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal("Error while reading body", err.Error())
		}

		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}

		_ = os.Setenv("FS_AUTH_ACCESS_TOKEN", result.AccessToken)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		log.WithFields(log.Fields{
			"status": resp.Status,
			"body":   bytes.NewBuffer(body),
		}).Fatal("Error when calling Flagship authentication API")
	}

	return nil
}

func callAPI(envID string, flagInfos FlagUsageRequest) error {

	if os.Getenv("FS_AUTH_ACCESS_TOKEN") == "" {
		err := generateAuthenticationToken()
		if err != nil {
			return err
		}
	}

	route := fmt.Sprintf("%s/v1/accounts/%s/account_environments/%s/flags_usage", os.Getenv("FS_API"), os.Getenv("ACCOUNT_ID"), envID)

	body, _ := json.Marshal(flagInfos)

	log.WithFields(log.Fields{
		"method": "POST",
		"route":  route,
	}).Info("Calling Flagship api")

	req, err := http.NewRequest("POST", route, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error in request", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("FS_AUTH_ACCESS_TOKEN")))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal("Error while reading body", err.Error())
		}

		log.WithFields(log.Fields{
			"status": resp.Status,
			"body":   bytes.NewBuffer(body),
		}).Fatal("Error when calling Flagship API")
	} else {
		log.Info("Synchronisation done with success")
	}

	defer resp.Body.Close()

	return nil
}
