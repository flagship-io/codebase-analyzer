package main

import (
	"log"

	"github.com/flagship-io/code-analyzer/pkg/handler"
)

func main() {
	err := handler.AnalyzeCode()
	if err != nil {
		log.Fatalf("Analyser failed with error: %v", err)
	}
}
