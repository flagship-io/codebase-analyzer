package handler

import (
	log "github.com/sirupsen/logrus"

	"github.com/flagship-io/codebase-analyzer/internal/api"
	"github.com/flagship-io/codebase-analyzer/internal/model"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
)

// AnalyzeCode loads and checks environment variables, extract flags from code and send flag infos to Flagship API
func AnalyzeCode(cfg *config.Config) error {
	// Environment variables for the project

	if cfg.SearchCustomRegex != "" {
		model.AddCustomRegexes(cfg.SearchCustomRegex)
	}

	results, err := ExtractFlagsInfo(cfg)

	if err != nil {
		log.Fatalf("Error occured when parsing code files: %v", err)
	}

	for _, r := range results {
		if len(results) > 0 {
			log.WithFields(log.Fields{
				"fileName":  r.File,
				"flagUsage": len(r.Results),
			}).Info("Scanned file")
		}
	}

	err = api.SendFlagsToAPI(cfg, results)
	return err
}
