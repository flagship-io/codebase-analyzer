package handler

import (
	"log"

	"github.com/flagship-io/codebase-analyzer/internal/files"
	"github.com/flagship-io/codebase-analyzer/internal/model"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
)

// ExtractFlagsInfo extract all flag usage information for code
func ExtractFlagsInfo(cfg *config.Config) ([]model.FileSearchResult, error) {
	if cfg.SearchCustomRegex != "" {
		model.AddCustomRegexes(cfg.SearchCustomRegex)
	}

	// List all files within the current directory
	filePaths, err := files.ListFiles(cfg.Directory, cfg.FilesToExcludes)

	if err != nil {
		log.Panicf("Error occured when listing files : %v", err)
	}

	results := []model.FileSearchResult{}
	resultsChan := make(chan model.FileSearchResult)

	for _, f := range filePaths {
		go files.SearchFiles(cfg, f, resultsChan)
	}

	for range filePaths {
		r := <-resultsChan
		results = append(results, r)
	}

	return results, err
}
