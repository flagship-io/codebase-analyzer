package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/flagship-io/code-analyzer/internal/files/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FlagInfo represent a flag code info sent to the Flagship API
type FlagInfo struct {
	FlagKey          string `json:"flagKey"`
	RepositoryURL    string `json:"repositoryUrl"`
	RepositoryBranch string `json:"repositoryBranch"`
	FilePath         string `json:"filePath"`
	LineNumber       int    `json:"lineNumber"`
	Code             string `json:"code"`
}

// SendFlagsToAPI takes file search result & sends flag info to the API
func SendFlagsToAPI(results []model.FileSearchResult, envId string) (err error) {
	flagInfos := []FlagInfo{}
	for _, fr := range results {
		for _, r := range fr.Results {
			flagInfos = append(flagInfos, FlagInfo{
				FlagKey:          r.FlagKey,
				RepositoryURL:    os.Getenv("REPOSITORY_URL"),
				RepositoryBranch: os.Getenv("REPOSITORY_BRANCH"),
				FilePath:         fr.File,
				LineNumber:       r.LineNumber,
				Code:             r.CodeLines,
			})
		}
	}

	err = callAPI(envId, flagInfos)

	return err
}

func callAPI(envID string, flagInfos []FlagInfo) error {
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
