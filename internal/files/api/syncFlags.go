package api

import (
	"gitlab/canarybay/aws/integration/code-analyser/internal/files/model"
	"os"
)

// FlagInfo represent a flag code info sent to the Flagship API
type FlagInfo struct {
	FlagKey       string `json:"flagKey"`
	RepositoryURL string `json:"repositoryUrl"`
	FilePath      string `json:"filePath"`
	LineNumber    int    `json:"lineNumber"`
	Code          string `json:"code"`
}

// SendFlagsToAPI takes file search result & sends flag info to the API
func SendFlagsToAPI(results []model.FileSearchResult) (err error) {
	flagInfos := []FlagInfo{}
	for _, fr := range results {
		for _, r := range fr.Results {
			flagInfos = append(flagInfos, FlagInfo{
				FlagKey:       r.FlagKey,
				RepositoryURL: os.Getenv("REPOSITORY_URL"),
				FilePath:      fr.File,
				LineNumber:    r.LineNumber,
				Code:          r.CodeLines,
			})
		}
	}

	// TODO : send call to API to route PUT /account_environments/{id}/tokens
	// {id} should come from environment variable

	return err
}
