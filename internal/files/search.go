package files

import (
	"fmt"
	"github/flagship-io/code-analyzer/internal/files/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type languageRegex struct {
	flagRegexes  []string
	flagKeyRegex string
}

var regexes = map[string]languageRegex{
	".js": {
		flagRegexes: []string{
			`(?s)useFsModifications\(.*?\)`,
			`(?s)\.getModifications\(.*?\)`,
		},
		flagKeyRegex: `['"]?key['"]?\s*\:\s*['"](.*)['"]`,
	},
	// TODO: add all languages
}

// SearchFiles search code pattern in files and return results and error
func SearchFiles(path string, resultChannel chan model.FileSearchResult) {
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

	// Get file extension to choose matching regex
	ext := filepath.Ext(path)
	regexes, ok := regexes[ext]
	if !ok {
		resultChannel <- model.FileSearchResult{
			File:    path,
			Results: nil,
			Error:   fmt.Errorf("File extension %s not handled", ext),
		}
		return
	}

	results := []model.SearchResult{}

	flagResults := [][]int{}
	for _, flagRegexString := range regexes.flagRegexes {
		// Find all flag code occurences within file
		flagRegex := regexp.MustCompile(flagRegexString)
		flagResults = append(flagResults, flagRegex.FindAllStringIndex(string(fileContent), -1)...)
	}

	for _, flagResult := range flagResults {
		keyRegex := regexp.MustCompile(regexes.flagKeyRegex)

		// Extract the flag code part
		submatch := string(fileContent[flagResult[0]:flagResult[1]])

		// Find the key name in the flag code part
		flagKeyResults := keyRegex.FindStringSubmatch(submatch)
		if len(flagKeyResults) < 2 {
			log.Printf("Did not find the flag key in file %s. Code : %s", path, submatch)
			continue
		}

		lineNumber := getLineFromPos(string(fileContent), flagResult[0])
		results = append(results, model.SearchResult{
			FlagKey:     flagKeyResults[1],
			CodeLines:   strings.TrimSpace(submatch),
			CodeLineURL: getCodeURL(path, &lineNumber),
			// Get line number of the code
			LineNumber: lineNumber,
		})
	}

	resultChannel <- model.FileSearchResult{
		File:    path,
		FileURL: getCodeURL(path, nil),
		Results: results,
		Error:   err,
	}
}

func getCodeURL(filePath string, line *int) string {
	repositoryURL := os.Getenv("REPOSITORY_URL")
	repositoryBranch := os.Getenv("REPOSITORY_BRANCH")

	repositorySep := ""
	if repositoryURL[len(repositoryURL)-1:] != "/" {
		repositorySep = "/"
	}
	lineAnchor := ""
	if line != nil {
		lineAnchor = fmt.Sprintf("#L%d", *line)
	}
	return fmt.Sprintf("%s%s-/blob/%s/%s%s", repositoryURL, repositorySep, repositoryBranch, filePath, lineAnchor)
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
