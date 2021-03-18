package files

import (
	"github/flagship-io/code-analyzer/internal/files/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("REPOSITORY_URL", "wwww.toto.com")
	os.Setenv("REPOSITORY_BRANCH", "master")
	os.Setenv("NB_CODE_LINES_EDGES", "5")
}

func TestSearchFiles(t *testing.T) {
	resultChannel := make(chan model.FileSearchResult)
	var r model.FileSearchResult

	go SearchFiles("../../example/src/sample.js", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "btnColor", r.Results[0].FlagKey)
	assert.Equal(t, "customLabel", r.Results[1].FlagKey)
	assert.Equal(t, "key", r.Results[2].FlagKey)
	assert.Equal(t, 3, len(r.Results))

	go SearchFiles("../../example/src/sample.jsx", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "backgroundColor", r.Results[0].FlagKey)
	assert.Equal(t, "btnColor", r.Results[1].FlagKey)
	assert.Equal(t, 2, len(r.Results))

	go SearchFiles("../../example/src/sample.go", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "btnColor", r.Results[0].FlagKey)
	assert.Equal(t, 5, len(r.Results))

	go SearchFiles("../../example/src/sample.py", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "btnColor", r.Results[0].FlagKey)
	assert.Equal(t, 1, len(r.Results))

	go SearchFiles("../../example/src/java/sample.java", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "btnColor", r.Results[0].FlagKey)
	assert.Equal(t, 1, len(r.Results))

	go SearchFiles("../../example/src/java/sample.kt", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "btnColor", r.Results[0].FlagKey)
	assert.Equal(t, 1, len(r.Results))

	go SearchFiles("../../example/src/swift/sample.swift", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "freeDelivery", r.Results[0].FlagKey)
	assert.Equal(t, "btnColor", r.Results[1].FlagKey)
	assert.Equal(t, 2, len(r.Results))

	go SearchFiles("../../example/src/swift/sample.m", resultChannel)
	r = <-resultChannel
	assert.Equal(t, "btnColor", r.Results[0].FlagKey)
	assert.Equal(t, 1, len(r.Results))
}

func TestGetSurroundingLineIndex(t *testing.T) {
	code := `	1;
2;
	3;
4;
5;
6;
	7;
8;
9;
	10;
	11;
12;
	13;
14;
15;`
	codeSampleExpected := `	3;
4;
5;
6;
	7;
8;
9;
	10;
	11;
12;
	13;`
	start := getSurroundingLineIndex(code, 24, true)
	end := getSurroundingLineIndex(code, 24, false)

	assert.Equal(t, 7, start)
	assert.Equal(t, 48, end)
	assert.Equal(t, codeSampleExpected, code[start:end])
}

func TestGetSurroundingLineIndexEdgeCases(t *testing.T) {
	code := `	1;
2;
	3;
	4;
5;`
	start := getSurroundingLineIndex(code, 8, true)
	end := getSurroundingLineIndex(code, 8, false)

	assert.Equal(t, 0, start)
	assert.Equal(t, 17, end)
	assert.Equal(t, code, code[start:end])
}
