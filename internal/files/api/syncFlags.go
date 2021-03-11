package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab/canarybay/aws/integration/code-analyser/internal/files/model"
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

	callApi(envId, flagInfos)

	return err
}

func callApi(envId string, flagInfos []FlagInfo) {
	syncFlagsRoute := os.Getenv("FS_API") + "/account_environments/" + envId + "/flag_usages"

	json, _ := json.Marshal(flagInfos)

	fmt.Println("Calling PUT", syncFlagsRoute)

	req, err := http.NewRequest("PUT", syncFlagsRoute, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Error in request", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("FLAGSHIP_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error in request execution", err.Error())
	}

	if resp.StatusCode == 200 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error", resp.StatusCode)
	}

	defer resp.Body.Close()
}
