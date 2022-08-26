package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/flagship-io/code-analyzer/pkg/config"
	"github.com/flagship-io/code-analyzer/pkg/handler"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func extractConfig() *config.Config {
	cfg := &config.Config{
		FlagshipAPIURL:        "https://api.flagship.io",
		FlagshipAuthAPIURL:    "https://auth.flagship.io",
		FlagshipClientID:      os.Getenv("FLAGSHIP_CLIENT_ID"),
		FlagshipClientSecret:  os.Getenv("FLAGSHIP_CLIENT_SECRET"),
		FlagshipAPIToken:      os.Getenv("FS_AUTH_ACCESS_TOKEN"),
		FlagshipAccountID:     os.Getenv("ACCOUNT_ID"),
		FlagshipEnvironmentID: os.Getenv("ENVIRONMENT_ID"),
		Directory:             ".",
		RepositoryURL:         os.Getenv("REPOSITORY_URL"),
		RepositoryBranch:      "master",
		FilesToExcludes:       []string{".git"},
		NbLineCodeEdges:       1,
		SearchCustomRegex:     os.Getenv("CUSTOM_REGEX_JSON"),
	}

	if os.Getenv("FS_API") != "" {
		cfg.FlagshipAPIURL = os.Getenv("FS_API")
	}

	if os.Getenv("FS_AUTH_API") != "" {
		cfg.FlagshipAuthAPIURL = os.Getenv("FS_AUTH_API")
	}

	if os.Getenv("REPOSITORY_BRANCH") != "" {
		cfg.RepositoryBranch = os.Getenv("REPOSITORY_BRANCH")
	}

	if os.Getenv("FILES_TO_EXCLUDE") != "" {
		cfg.FilesToExcludes = strings.Split(os.Getenv("FILES_TO_EXCLUDE"), ",")
	}

	if os.Getenv("DIRECTORY") != "" {
		cfg.Directory = os.Getenv("DIRECTORY")
	}

	if os.Getenv("NB_CODE_LINES_EDGES") != "" {
		cfg.NbLineCodeEdges, _ = strconv.Atoi(os.Getenv("NB_CODE_LINES_EDGES"))
	}

	return cfg
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetLevel(log.WarnLevel)

	if err := godotenv.Load(); err != nil {
		log.Info("Fail to load .env file")
	}

	levelString := os.Getenv("LOG_LEVEL")
	if levelString != "" {
		lvl, err := log.ParseLevel(levelString)
		if err != nil {
			log.Fatalf("log level %s is invalid", levelString)
		}
		log.SetLevel(lvl)
	}

	cfg := extractConfig()
	cfg.Validate()

	err := handler.AnalyzeCode(cfg)
	if err != nil {
		log.Fatalf("Analyser failed with error: %v", err)
	}
}
