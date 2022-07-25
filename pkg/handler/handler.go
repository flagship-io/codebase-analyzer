package handler

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/flagship-io/code-analyzer/internal/files/api"
	"github.com/flagship-io/code-analyzer/internal/files/model"
)

// AnalyzeCode loads and checks environment variables, extract flags from code and send flag infos to Flagship API
func AnalyzeCode() error {
	// Environment variables for the project

	if os.Getenv("FS_API") == "" {
		_ = os.Setenv("FS_API", "https://api.flagship.io")
	}

	if os.Getenv("FS_AUTH_API") == "" {
		_ = os.Setenv("FS_AUTH_API", "https://auth.flagship.io")
	}

	// Environment variables to set by the client

	if os.Getenv("REPOSITORY_URL") == "" {
		log.WithFields(log.Fields{"variable": "REPOSITORY_URL"}).Fatal("Missing required environment variable")
	}

	if os.Getenv("ENVIRONMENT_ID") == "" {
		log.WithFields(log.Fields{"variable": "ENVIRONMENT_ID"}).Fatal("Missing required environment variable")
	}

	if os.Getenv("REPOSITORY_BRANCH") == "" {
		_ = os.Setenv("REPOSITORY_BRANCH", "master")
	}

	if os.Getenv("FILES_TO_EXCLUDE") == "" {
		_ = os.Setenv("FILES_TO_EXCLUDE", ".git")
	}

	toExclude := strings.Split(os.Getenv("FILES_TO_EXCLUDE"), ",")

	dir := os.Getenv("DIRECTORY")
	if dir == "" {
		dir = "."
	}

	if os.Getenv("NB_CODE_LINES_EDGES") == "" {
		_ = os.Setenv("NB_CODE_LINES_EDGES", "1")
	}

	if os.Getenv("CUSTOM_REGEX_JSON") != "" {
		model.AddCustomRegexes(os.Getenv("CUSTOM_REGEX_JSON"))
	}

	results, err := ExtractFlagsInfo(dir, toExclude)

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

	err = api.SendFlagsToAPI(results, os.Getenv("ENVIRONMENT_ID"))
	return err
}
