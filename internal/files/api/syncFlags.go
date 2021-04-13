package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/flagship-io/code-analyzer/internal/files/model"
)

type FlagUsageRequest struct {
	RepositoryURL    string `json:"repositoryUrl"`
	RepositoryBranch string `json:"repositoryBranch"`
	Flags            []Flag `json:"flags"`
}

// Fl represent a flag code info sent to the Flagship API
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
	flags := []Flag{}
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

func callAPI(envID string, flagInfos FlagUsageRequest) error {
	syncFlagsRoute := fmt.Sprintf("%s/account_environments/%s/flag_usages", os.Getenv("FS_API"), envID)

	json, _ := json.Marshal(flagInfos)

	fmt.Println("Calling PUT", syncFlagsRoute)

	req, err := http.NewRequest("PUT", syncFlagsRoute, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Error in request", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", os.Getenv("FLAGSHIP_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Error when calling Flagship API. Got status %s, response: %s", resp.Status, body)
	}

	defer resp.Body.Close()
	return nil
}
