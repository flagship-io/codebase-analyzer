package files

import (
	"fmt"
	"testing"

	"github.com/flagship-io/codebase-analyzer/internal/model"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/stretchr/testify/assert"
)

var testCfg = &config.Config{
	RepositoryURL:    "wwww.toto.com",
	RepositoryBranch: "master",
	NbLineCodeEdges:  5,
}

func TestSearchFiles(t *testing.T) {

	type flag struct {
		name              string
		lineNumber        int
		codeLineHighlight int
	}

	type testCase struct {
		filePath string
		flags    []flag
	}

	var cases = []testCase{
		{
			filePath: "../../example/src/go/SDK_V2/sample.go",
			flags: []flag{
				{name: "btnColor", lineNumber: 31, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 32, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 33, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/ios/SDK_V2/sample.m",
			flags: []flag{
				{name: "btnColor", lineNumber: 21, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 22, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 23, codeLineHighlight: 6},
				{name: "btnSizeFloat", lineNumber: 24, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/ios/SDK_V3/sample.m",
			flags: []flag{
				{name: "btnColor", lineNumber: 8, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 9, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 10, codeLineHighlight: 6},
				{name: "btnSizeFloat", lineNumber: 11, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/ios/SDK_V2/sample.swift",
			flags: []flag{
				{name: "freeDelivery", lineNumber: 7, codeLineHighlight: 6},
				{name: "btnColor", lineNumber: 20, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 23, codeLineHighlight: 6},
				{name: "btnSizeFloat", lineNumber: 24, codeLineHighlight: 6},
				{name: "btnSizeInt", lineNumber: 25, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/ios/SDK_V3/sample.swift",
			flags: []flag{
				{name: "btnColor", lineNumber: 9, codeLineHighlight: 6},
				{name: "displayVipFeature", lineNumber: 10, codeLineHighlight: 6},
				{name: "vipFeature", lineNumber: 11, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/java/SDK_V2/sample.java",
			flags: []flag{
				{name: "btnColor", lineNumber: 36, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 37, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 38, codeLineHighlight: 6},
				{name: "showBackground", lineNumber: 39, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/java/SDK_V3/sample.java",
			flags: []flag{
				{name: "btnColor", lineNumber: 4, codeLineHighlight: 4},
				{name: "backgroundSize", lineNumber: 5, codeLineHighlight: 5},
				{name: "showBackground", lineNumber: 6, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 7, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/java/SDK_V2/sample.kt",
			flags: []flag{
				{name: "btnColor", lineNumber: 20, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 21, codeLineHighlight: 6},
				{name: "displayBtn", lineNumber: 22, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/java/SDK_V3/sample.kt",
			flags: []flag{
				{name: "btnColor", lineNumber: 10, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 11, codeLineHighlight: 6},
				{name: "showBackground", lineNumber: 12, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/js/SDK_V2/sample.js",
			flags: []flag{
				{name: "btnColor", lineNumber: 12, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 17, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 21, codeLineHighlight: 6},
				{name: "showbackground", lineNumber: 25, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/js/SDK_V3/sample.js",
			flags: []flag{
				{name: "btnColor", lineNumber: 15, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 16, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 17, codeLineHighlight: 6},
				{name: "showBackground", lineNumber: 18, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/js/SDK_V2/sample.ts",
			flags: []flag{
				{name: "btnColor", lineNumber: 12, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 17, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 21, codeLineHighlight: 6},
				{name: "showbackground", lineNumber: 25, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/js/SDK_V3/sample.ts",
			flags: []flag{
				{name: "btnColor", lineNumber: 15, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 16, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 17, codeLineHighlight: 6},
				{name: "showBackground", lineNumber: 18, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/net/SDK_V1/sample.fs",
			flags: []flag{
				{name: "btnColor", lineNumber: 11, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 12, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 13, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/net/SDK_V3/sample.fs",
			flags: []flag{
				{name: "btnColor", lineNumber: 13, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 14, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 15, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/net/SDK_V1/sample.cs",
			flags: []flag{
				{name: "btnColor", lineNumber: 12, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 13, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 14, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/net/SDK_V3/sample.cs",
			flags: []flag{
				{name: "btnColor", lineNumber: 15, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 16, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 17, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/net/SDK_V1/sample.vb",
			flags: []flag{
				{name: "btnColor", lineNumber: 11, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 12, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 13, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/net/SDK_V3/sample.vb",
			flags: []flag{
				{name: "btnColor", lineNumber: 12, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 13, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 14, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/python/SDK_V2/sample.py",
			flags: []flag{
				{name: "showBtn", lineNumber: 30, codeLineHighlight: 6},
				{name: "btnColor", lineNumber: 31, codeLineHighlight: 6},
				{name: "btnSize", lineNumber: 32, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/react/SDK_V2/sample.jsx",
			flags: []flag{
				{name: "backgroundColor", lineNumber: 16, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 21, codeLineHighlight: 6},
				{name: "showBackground", lineNumber: 26, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/react/SDK_V3/sample.jsx",
			flags: []flag{
				{name: "backgroundColor", lineNumber: 5, codeLineHighlight: 5},
				{name: "btnColor", lineNumber: 6, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 7, codeLineHighlight: 6},
				{name: "showBtn", lineNumber: 8, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/php/SDK_V1/sample.php",
			flags: []flag{
				{name: "displayVipFeature", lineNumber: 7, codeLineHighlight: 6},
				{name: "vipFeatureSize", lineNumber: 8, codeLineHighlight: 6},
				{name: "vipFeatureColor", lineNumber: 9, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/php/SDK_V2/sample.php",
			flags: []flag{
				{name: "displayVipFeature", lineNumber: 8, codeLineHighlight: 6},
				{name: "vipFeatureSize", lineNumber: 9, codeLineHighlight: 6},
				{name: "vipFeatureColor", lineNumber: 10, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/php/SDK_V3/sample.php",
			flags: []flag{
				{name: "displayVipFeature", lineNumber: 8, codeLineHighlight: 6},
				{name: "vipFeatureSize", lineNumber: 9, codeLineHighlight: 6},
				{name: "vipFeatureColor", lineNumber: 10, codeLineHighlight: 6},
			},
		},
		{
			filePath: "../../example/src/flutter/sample.dart",
			flags: []flag{
				{name: "displayVipFeature", lineNumber: 15, codeLineHighlight: 6},
				{name: "backgroundColor", lineNumber: 16, codeLineHighlight: 6},
				{name: "backgroundSize", lineNumber: 17, codeLineHighlight: 6},
			},
		},
	}

	resultChannel := make(chan model.FileSearchResult)
	var r model.FileSearchResult

	for _, c := range cases {
		go SearchFiles(testCfg, c.filePath, resultChannel)
		r = <-resultChannel
		assert.Equal(t, len(c.flags), len(r.Results), fmt.Sprintf("File : %s", c.filePath))
		for i, result := range r.Results {
			assert.Equal(t, c.flags[i].name, result.FlagKey, fmt.Sprintf("File : %s", c.filePath))
			assert.Equal(t, c.flags[i].lineNumber, result.LineNumber, fmt.Sprintf("File : %s", c.filePath))
			assert.Equal(t, c.flags[i].codeLineHighlight, result.CodeLineHighlight, fmt.Sprintf("File : %s", c.filePath))
			assert.Equal(t,
				fmt.Sprintf(
					"%s/-/blob/%s/%s#L%d",
					testCfg.RepositoryURL,
					testCfg.RepositoryBranch,
					c.filePath,
					c.flags[i].lineNumber,
				),
				result.CodeLineURL,
				fmt.Sprintf("File : %s", c.filePath),
			)
		}
	}
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
	start := getSurroundingLineIndex(code, 24, true, testCfg.NbLineCodeEdges)
	end := getSurroundingLineIndex(code, 24, false, testCfg.NbLineCodeEdges)

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
	start := getSurroundingLineIndex(code, 8, true, testCfg.NbLineCodeEdges)
	end := getSurroundingLineIndex(code, 8, false, testCfg.NbLineCodeEdges)

	assert.Equal(t, 0, start)
	assert.Equal(t, 17, end)
	assert.Equal(t, code, code[start:end])
}
