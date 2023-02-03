package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flagship-io/codebase-analyzer/internal/model"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
)

// SearchFiles search code pattern in files and return results and error
func SearchFiles(cfg *config.Config, path string, resultChannel chan model.FileSearchResult) {
	// Read file contents
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		resultChannel <- model.FileSearchResult{
			File:    path,
			Results: nil,
			Error:   err,
		}
		return
	}
	fileContentStr := strings.ReplaceAll(string(fileContent), "\r\n", "\n")

	// Get file extension to choose matching regex
	ext := filepath.Ext(path)
	var flagRegexes []model.FlagRegex
	for _, extRegex := range model.LanguageRegexes {
		regxp := regexp.MustCompile(extRegex.ExtensionRegex)
		if regxp.Match([]byte(ext)) {
			flagRegexes = extRegex.FlagRegexes
		}
	}
	if len(flagRegexes) == 0 {
		resultChannel <- model.FileSearchResult{
			File:    path,
			Results: nil,
			Error:   fmt.Errorf("file extension %s not handled", ext),
		}
		return
	}

	// Add default regex for flags in commentaries
	flagRegexes = append(flagRegexes, model.FlagRegex{
		FunctionRegex: `(?s)fs:flag:(\w+)`,
		FieldRegex:    `fs:flag:(.+)`,
	})

	results := []model.SearchResult{}

	flagIndexes := [][]int{}
	for _, flagRegex := range flagRegexes {
		regxp := regexp.MustCompile(flagRegex.FunctionRegex)
		flagLineIndexes := regxp.FindAllStringIndex(fileContentStr, -1)

		for _, flagLineIndex := range flagLineIndexes {
			submatch := fileContentStr[flagLineIndex[0]:flagLineIndex[1]]
			regxp := regexp.MustCompile(flagRegex.FieldRegex)

			submatchIndexes := regxp.FindAllStringSubmatchIndex(submatch, -1)

			for k, submatchIndex := range submatchIndexes {
				if len(submatchIndex) < 6 {
					log.Printf("Did not find the flag key in file %s. Code : %s", path, submatch)
					continue
				}
				if !flagRegex.HasMultipleKeys && k > 0 {
					break
				}

				flagIndexes = append(flagIndexes, []int{
					flagLineIndex[0] + submatchIndex[2],
					flagLineIndex[0] + submatchIndex[3],
					flagLineIndex[0] + submatchIndex[4],
					flagLineIndex[0] + submatchIndex[5],
				})
			}
		}
	}

	for _, flagIndex := range flagIndexes {
		// Extract the code with a certain number of lines
		firstLineIndex := getSurroundingLineIndex(fileContentStr, flagIndex[0], true, cfg.NbLineCodeEdges)
		lastLineIndex := getSurroundingLineIndex(fileContentStr, flagIndex[1], false, cfg.NbLineCodeEdges)
		code := fileContentStr[firstLineIndex:lastLineIndex]
		key := fileContentStr[flagIndex[0]:flagIndex[1]]
		defaultValue := fileContentStr[flagIndex[2]:flagIndex[3]]
		// Better value wrapper for code highlighting (5 chars wrapping)
		keyWrapper := key
		nbCharsWrapping := 5
		if flagIndex[0] > nbCharsWrapping && flagIndex[1] < len(fileContentStr)-nbCharsWrapping {
			keyWrapper = fileContentStr[flagIndex[0]-nbCharsWrapping : flagIndex[1]+nbCharsWrapping]
		}

		lineNumber := getLineFromPos(fileContentStr, flagIndex[0])
		codeLineHighlight := getLineFromPos(code, strings.Index(code, keyWrapper))

		var flagType string = "text"
		var flagTypeInterface interface{}

		r, _ := regexp.Compile(`[^\w#]`)

		json.Unmarshal([]byte(defaultValue), &flagTypeInterface)

		if _, isNumber := flagTypeInterface.(float64); isNumber {
			flagType = "number"
		}

		if _, isBool := flagTypeInterface.(bool); isBool {
			flagType = "boolean"
		}

		if match := r.MatchString(defaultValue); match {
			flagType = "unknown"
		}

		results = append(results, model.SearchResult{
			FlagKey:           key,
			FlagDefaultValue:  defaultValue,
			FlagType:          flagType,
			CodeLines:         code,
			CodeLineHighlight: codeLineHighlight,
			CodeLineURL:       getCodeURL(cfg, path, &lineNumber),
			// Get line number of the code
			LineNumber: lineNumber,
		})

	}

	resultChannel <- model.FileSearchResult{
		File:    path,
		FileURL: getCodeURL(cfg, path, nil),
		Results: results,
		Error:   err,
	}
}

func getCodeURL(cfg *config.Config, filePath string, line *int) string {
	repositoryURL := strings.TrimSuffix(cfg.RepositoryURL, "/")
	lineAnchor := ""
	if line != nil {
		lineAnchor = fmt.Sprintf("#L%d", *line)
	}
	return fmt.Sprintf("%s/-/blob/%s/%s%s", repositoryURL, cfg.RepositoryBranch, filePath, lineAnchor)
}

func getSurroundingLineIndex(input string, indexPosition int, topDirection bool, nbLineCodeEdges int) int {
	i := indexPosition
	n := 0
	for {
		// to top or bottom
		if topDirection {
			i--
		} else {
			i++
		}

		// edge cases
		if i <= 0 {
			return 0
		}
		if i >= len(input)-1 {
			return len(input)
		}

		// if new line, in the top direction we don't wan't the first \n in the code
		if input[i] == '\n' {
			if n == nbLineCodeEdges {
				if topDirection {
					return i + 1
				}
				return i
			}
			n++
		} else {
			continue
		}
	}
}

func getLineFromPos(input string, indexPosition int) int {
	lineNumber := 1
	for i := 0; i <= indexPosition; i++ {
		if input[i] == '\n' {
			lineNumber++
		}
	}
	return lineNumber
}
