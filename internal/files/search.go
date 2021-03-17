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
	"strconv"
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
	".go": {
		flagRegexes: []string{
			`(?s)\.GetModification(String|Number|Bool|Object|Array)\(.*?\)`,
		},
		flagKeyRegex: `s*['"](.*?)['"]`,
	},
	".py": {
		flagRegexes: []string{
			`(?s)\.get_modification\(.*?\)`,
		},
		flagKeyRegex: `s*['"](.*?)['"]`,
	},
	".java": {
		flagRegexes: []string{
			`(?s)\.getModification\(.*?\)`,
		},
		flagKeyRegex: `s*['"](.*?)['"]`,
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
	fileContentStr := strings.ReplaceAll(string(fileContent), "\r\n", "\n")

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
		flagResults = append(flagResults, flagRegex.FindAllStringIndex(fileContentStr, -1)...)
	}

	for _, flagResult := range flagResults {
		keyRegex := regexp.MustCompile(regexes.flagKeyRegex)

		// Extract the flag code part
		submatch := fileContentStr[flagResult[0]:flagResult[1]]
		// Extract the code with a certain number of lines
		firstLineIndex := getSurroundingLineIndex(fileContentStr, flagResult[0], true);
		lastLineIndex := getSurroundingLineIndex(fileContentStr, flagResult[1], false);
		code := strings.TrimSpace(fileContentStr[firstLineIndex:lastLineIndex])

		// Find the key name in the flag code part
		flagKeyResults := keyRegex.FindStringSubmatch(submatch)
		if len(flagKeyResults) < 2 {
			log.Printf("Did not find the flag key in file %s. Code : %s", path, submatch)
			continue
		}

		lineNumber := getLineFromPos(fileContentStr, flagResult[0])
		results = append(results, model.SearchResult{
			FlagKey:     flagKeyResults[1],
			CodeLines:   code,
			CodeLineHighlight: getLineFromPos(code, strings.Index(code, submatch)),
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

func getSurroundingLineIndex(input string, indexPosition int, topDirection bool) int {
	i := indexPosition
	n := 0
	e, _ := strconv.Atoi(os.Getenv("NB_CODE_LINES_EDGES"))
	for {
		// to top or bottom
		if topDirection {
			i--
		} else {
			i++
		}

		if n == e && topDirection {
			return i
		} else if n == e {
			return i-1
		}
		
		// edge cases
		if i <= 0 {
			return 0
		}
		if i >= len(input) - 1 {
			return len(input)
		}

		// if new line
		if input[i] == '\n' {
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
