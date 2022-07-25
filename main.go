package main

import (
	"github.com/flagship-io/code-analyzer/pkg/handler"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.WarnLevel)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	if os.Getenv("ENABLE_LOGS") == "1" {
		log.SetLevel(log.InfoLevel)
	}

	err := handler.AnalyzeCode()
	if err != nil {
		log.Fatalf("Analyser failed with error: %v", err)
	}
}
