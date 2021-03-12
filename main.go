package main

import (
	"gitlab/canarybay/aws/integration/code-analyser/pkg/handler"
	"log"
)

func main() {
	err := handler.AnalyzeCode()
	if err != nil {
		log.Fatalf("Analyser failed with error: %v", err)
	}
}
