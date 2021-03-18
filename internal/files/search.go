package files

import (
	"fmt"
	"github/flagship-io/code-analyzer/internal/files/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
			`(?s)\.getModifications\(.*?\].*?\)`,
		},
		flagKeyRegex: `['"]?key['"]?\s*\:\s*['"](.*?)['"]`,
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

	// [$start $end]
	flagIndexes := [][]int{}
	for _, flagRegexString := range regexes.flagRegexes {
		regxp := regexp.MustCompile(flagRegexString)
		flagIndexes = append(flagIndexes, regxp.FindAllStringIndex(fileContentStr, -1)...)
	}

	// [$start $end $flagStart $flagEnd]
	flagKeyIndexes := [][]int{}
	for _, flagIndex := range flagIndexes {
		submatch := fileContentStr[flagIndex[0]:flagIndex[1]]
		regxp := regexp.MustCompile(regexes.flagKeyRegex)

		submatchIndexes := regxp.FindAllStringSubmatchIndex(submatch, -1)

		for _, submatchIndex := range submatchIndexes {
			if len(submatchIndex) < 4 {
				log.Printf("Did not find the flag key in file %s. Code : %s", path, submatch)
				continue
			}

			flagKeyIndexes = append(flagKeyIndexes, []int{
				flagIndex[0] + submatchIndex[2],
				flagIndex[0] + submatchIndex[3],
			})
		}
	}

	for _, flagKeyIndex := range flagKeyIndexes {
		// Extract the code with a certain number of lines
		firstLineIndex := getSurroundingLineIndex(fileContentStr, flagKeyIndex[0], true)
		lastLineIndex := getSurroundingLineIndex(fileContentStr, flagKeyIndex[1], false)
		code := fileContentStr[firstLineIndex:lastLineIndex]
		value := fileContentStr[flagKeyIndex[0]:flagKeyIndex[1]]

		lineNumber := getLineFromPos(fileContentStr, flagKeyIndex[0])
		codeLineHighlight := getLineFromPos(code, strings.Index(code, value))
		fmt.Println(value, codeLineHighlight)
		results = append(results, model.SearchResult{
			FlagKey:           value,
			CodeLines:         code,
			CodeLineHighlight: codeLineHighlight,
			CodeLineURL:       getCodeURL(path, &lineNumber),
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

		// edge cases
		if i <= 0 {
			return 0
		}
		if i >= len(input)-1 {
			return len(input)
		}

		// if new line, in the top direction we don't wan't the first \n in the code
		if input[i] == '\n' {
			if n == e {
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
