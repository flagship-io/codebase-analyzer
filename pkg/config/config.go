package config

import log "github.com/sirupsen/logrus"

type Config struct {
	FlagshipAPIURL        string
	FlagshipAuthAPIURL    string
	FlagshipClientID      string
	FlagshipClientSecret  string
	FlagshipAPIToken      string
	FlagshipAccountID     string
	FlagshipEnvironmentID string
	Directory             string
	RepositoryURL         string
	RepositoryBranch      string
	FilesToExclude        []string
	NbLineCodeEdges       int
	SearchCustomRegex     string
}

func (cfg *Config) Validate() {
	// Validate required options to be set by the client
	if cfg.RepositoryURL == "" {
		log.WithFields(log.Fields{"variable": "REPOSITORY_URL"}).Fatal("Missing required environment variable")
	}

	if cfg.FlagshipEnvironmentID == "" {
		log.WithFields(log.Fields{"variable": "ENVIRONMENT_ID"}).Fatal("Missing required environment variable")
	}

	if cfg.FlagshipClientID == "" {
		log.WithFields(log.Fields{"variable": "FLAGSHIP_CLIENT_ID"}).Fatal("Missing required environment variable")
	}

	if cfg.FlagshipClientSecret == "" {
		log.WithFields(log.Fields{"variable": "FLAGSHIP_CLIENT_SECRET"}).Fatal("Missing required environment variable")
	}

	if cfg.FlagshipAccountID == "" {
		log.WithFields(log.Fields{"variable": "ACCOUNT_ID"}).Fatal("Missing required environment variable")
	}
}
