package main

import (
	"gitlab/canarybay/aws/integration/code-analyser/pkg/handler"
)

var toExclude = []string{".git/*", "go.mod", "go.sum", "main.go", "internal/*", "example/*"}

func main() {
	handler.AnalyzeCode()
}
