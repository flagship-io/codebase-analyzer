package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/flagship-io/codebase-analyzer/internal/model"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
)

type SendFlagsParameters struct {
	FlagshipAPIURL       string
	FlagshipAuthAPIURL   string
	FlagshipClientID     string
	FlagshipClientSecret string
	RepositoryURL        string
	RepositoryBranch     string
}

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
func SendFlagsToAPI(cfg *config.Config, results []model.FileSearchResult) (err error) {
	flagUsageRequest := FlagUsageRequest{
		RepositoryURL:    cfg.RepositoryURL,
		RepositoryBranch: cfg.RepositoryBranch,
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

	err = callAPI(cfg, flagUsageRequest)

	return err
}

func generateAuthenticationToken(cfg *config.Config) (string, error) {

	authRequest := AuthRequest{
		GrantType:    "client_credentials",
		Scope:        "*",
		ClientId:     cfg.FlagshipClientID,
		ClientSecret: cfg.FlagshipClientSecret,
	}

	body, err := json.Marshal(authRequest)

	if err != nil {
		log.Fatal("Error while marshal json", err.Error())
	}

	route := fmt.Sprintf("%s/%s/token?expires_in=0", cfg.FlagshipAuthAPIURL, cfg.FlagshipAccountID)

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
		return "", err
	}

	if resp.StatusCode == 200 {
		if err != nil {
			return "", err
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

		return result.AccessToken, nil
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error when calling Flagship authentication API. Status: %s, body: %s", resp.Status, string(body))
	}
}

func callAPI(cfg *config.Config, flagInfos FlagUsageRequest) error {

	if cfg.FlagshipAPIToken == "" {
		token, err := generateAuthenticationToken(cfg)
		if err != nil {
			return err
		}
		cfg.FlagshipAPIToken = token
	}

	route := fmt.Sprintf("%s/v1/accounts/%s/account_environments/%s/flags_usage", cfg.FlagshipAPIURL, cfg.FlagshipAccountID, cfg.FlagshipEnvironmentID)

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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.FlagshipAPIToken))

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
