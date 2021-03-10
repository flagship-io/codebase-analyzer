package handler

import (
	"gitlab/canarybay/aws/integration/code-analyser/internal/files"
	"gitlab/canarybay/aws/integration/code-analyser/internal/files/model"
	"log"
)

// ExtractFlagsInfo extract all flag usage information for code
func ExtractFlagsInfo(dir string, toExclude []string) ([]model.FileSearchResult, error) {
	// List all files within the current directory
	filePaths, err := files.ListFiles(dir, toExclude)

	if err != nil {
		log.Panicf("Error occured when listing files : %v", err)
	}

	results := []model.FileSearchResult{}
	resultsChan := make(chan model.FileSearchResult)

	for _, f := range filePaths {
		go files.SearchFiles(f, resultsChan)
	}

	for range filePaths {
		r := <-resultsChan
		results = append(results, r)
	}

	return results, err
}
