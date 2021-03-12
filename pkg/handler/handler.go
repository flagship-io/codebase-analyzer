package handler

import (
	"gitlab/canarybay/aws/integration/code-analyser/internal/files/api"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// AnalyzeCode loads and checks environment variables, extract flags from code and send flag infos to Flagship API
func AnalyzeCode() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file : %v", err)
	}

	// Environment variables for the prcject

	if os.Getenv("FS_API") == "" {
		os.Setenv("FS_API", "https://api.flagship.io")
	}

	// Environment variables to set by the client

	fsToken := os.Getenv("FLAGSHIP_TOKEN")
	if fsToken == "" {
		log.Fatal("Missing required environment variable FLAGSHIP_TOKEN")
	}

	repoURL := os.Getenv("REPOSITORY_URL")
	if repoURL == "" {
		log.Fatal("Missing required environment variable REPOSITORY_URL")
	}

	envID := os.Getenv("ENVIRONMENT_ID")
	if envID == "" {
		log.Fatal("Missing required environment variable ENVIRONMENT_ID")
	}

	repoBranch := os.Getenv("REPOSITORY_BRANCH")
	if repoBranch == "" {
		repoBranch = "master"
	}

	if os.Getenv("FILES_TO_EXCLUDE") == "" {
		os.Setenv("FILES_TO_EXCLUDE", ".git")
	}
	toExclude := strings.Split(os.Getenv("FILES_TO_EXCLUDE"), ",")

	dir := os.Getenv("DIRECTORY")
	if dir == "" {
		dir = "."
	}

	results, err := ExtractFlagsInfo(dir, toExclude)

	if err != nil {
		log.Fatalf("Error occured when parsing code files: %v", err)
	}

	for _, r := range results {
		log.Printf("Scanned file %s and found %d flag usages", r.File, len(r.Results))
	}

	err = api.SendFlagsToAPI(results, envID)
	return err
}
