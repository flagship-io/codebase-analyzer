package main

import (
	"github/flagship-io/code-analyzer/pkg/handler"
	"log"
)

func main() {
	err := handler.AnalyzeCode()
	if err != nil {
		log.Fatalf("Analyser failed with error: %v", err)
	}
}
