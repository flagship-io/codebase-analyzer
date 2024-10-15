package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flagship-io/codebase-analyzer/internal/model"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type MatchVariableFlag struct {
	Variable          string
	FlagKey           string
	CodeLines         string
	CodeLineHighlight int
	CodeLineURL       string
	LineNumber        int
}

type MatchVariableDefaultValue struct {
	Variable         string
	FlagDefaultValue string
	FlagType         string
}

type RegexStruct struct {
	isFlag         bool
	isDefaultValue bool
	regex          []string
}

type FlagIndexesStruct struct {
	isFlag         bool
	isDefaultValue bool
	FlagIndexes    [][]int
}

func GetFlagType(defaultValue string) (string, string) {

	var flagType string = "string"
	var flagTypeInterface interface{}

	r, _ := regexp.Compile(`[\{\}\[\]]`)

	json.Unmarshal([]byte(defaultValue), &flagTypeInterface)

	if match := r.MatchString(defaultValue); match {
		flagType = "unknown"
	}

	if len(defaultValue) > 0 && (defaultValue[0:1] == "\"" || defaultValue[0:1] == "'") && (defaultValue[len(defaultValue)-1:] == "\"" || defaultValue[len(defaultValue)-1:] == "'") {
		flagType = "string"
	}

	if funk.ContainsString([]string{"TRUE", "YES", "True"}, defaultValue) {
		defaultValue = "true"
		flagType = "boolean"
	}
	if funk.ContainsString([]string{"FALSE", "NO", "False"}, defaultValue) {
		defaultValue = "false"
		flagType = "boolean"
	}

	if _, isNumber := flagTypeInterface.(float64); isNumber {
		flagType = "number"
	}

	if _, isBool := flagTypeInterface.(bool); isBool {
		flagType = "boolean"
	}

	return flagType, defaultValue
}

// SearchFiles search code pattern in files and return results and error
func SearchFiles(cfg *config.Config, path string, resultChannel chan model.FileSearchResult) {
	// Read file contents
	fileContent, err := os.ReadFile(path)
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
	var regexesNotSplit []string
	var regexesSplit []RegexStruct
	for _, extRegex := range model.LanguageRegexes {
		if !extRegex.Split {
			regxp := regexp.MustCompile(extRegex.FileExtension)
			if regxp.Match([]byte(ext)) {
				regexesNotSplit = append(regexesNotSplit, extRegex.Regexes...)
			}
		}

		if extRegex.Split {
			regxp := regexp.MustCompile(extRegex.FileExtension)
			if regxp.Match([]byte(ext)) {
				regexesSplit = append(regexesSplit, RegexStruct{isFlag: extRegex.ForFlag, isDefaultValue: extRegex.ForDefaultValue, regex: extRegex.Regexes})
			}
		}
	}

	if len(regexesNotSplit) == 0 && len(regexesSplit) == 0 {
		resultChannel <- model.FileSearchResult{
			File:    path,
			Results: nil,
			Error:   fmt.Errorf("file extension %s not handled", ext),
		}
		return
	}

	// Add default regex for flags in commentaries
	regexesNotSplit = append(regexesNotSplit,
		`fe:flag:\s*(\w+)\s*[,]\s*(\w+)\s*`,
	)

	resultsNotSplit := []model.SearchResult{}
	flagIndexesNotSplit := [][]int{}
	flagIndexesSplit := flagIndexSplitter(fileContentStr, path, regexesSplit)

	for _, regex := range regexesNotSplit {
		regxp := regexp.MustCompile(regex)
		flagLineIndexes := regxp.FindAllStringIndex(fileContentStr, -1)

		for _, flagLineIndex := range flagLineIndexes {
			submatch := fileContentStr[flagLineIndex[0]:flagLineIndex[1]]

			submatchIndexes := regxp.FindAllStringSubmatchIndex(submatch, -1)

			for _, submatchIndex := range submatchIndexes {
				if len(submatchIndex) < 3 {
					log.WithFields(log.Fields{
						"reason": fmt.Sprintf("Did not find the flag key in file %s. Code: %s", path, submatch),
					}).Error("Key not found")
					continue
				}

				if len(submatchIndex) < 6 {
					log.WithFields(log.Fields{
						"reason": fmt.Sprintf("Did not find the flag default value in file %s. Code: %s", path, submatch),
					}).Warn("Type unknown")
					flagIndexesNotSplit = append(flagIndexesNotSplit, []int{
						flagLineIndex[0] + submatchIndex[2],
						flagLineIndex[0] + submatchIndex[3],
					})
					continue
				}

				flagIndexesNotSplit = append(flagIndexesNotSplit, []int{
					flagLineIndex[0] + submatchIndex[2],
					flagLineIndex[0] + submatchIndex[3],
					flagLineIndex[0] + submatchIndex[4],
					flagLineIndex[0] + submatchIndex[5],
				})
			}
		}
	}

	var flagMatches []MatchVariableFlag
	var defaultValueMatches []MatchVariableDefaultValue

	for _, flagIndexes := range flagIndexesSplit {
		for _, flagIndex := range flagIndexes.FlagIndexes {
			variable := fileContentStr[flagIndex[0]:flagIndex[1]]
			if len(flagIndex) < 3 {
				log.WithFields(log.Fields{
					"reason": fmt.Sprintf("Did not find the flag key or default value in file %s", path),
				}).Error("Key not found")
				continue
			}

			keyOrDefaultValue := fileContentStr[flagIndex[2]:flagIndex[3]]

			if flagIndexes.isFlag {
				firstLineIndex := getSurroundingLineIndex(fileContentStr, flagIndex[0], true, cfg.NbLineCodeEdges)
				lastLineIndex := getSurroundingLineIndex(fileContentStr, flagIndex[1], false, cfg.NbLineCodeEdges)
				code := fileContentStr[firstLineIndex:lastLineIndex]
				keyWrapped := keyWrapper(keyOrDefaultValue, fileContentStr, flagIndex)
				lineNumber := getLineFromPos(fileContentStr, flagIndex[2])
				codeLineHighlight := getLineFromPos(code, strings.Index(code, keyWrapped))
				_ = codeLineHighlight

				flagMatches = append(flagMatches, MatchVariableFlag{
					Variable:          variable,
					FlagKey:           keyOrDefaultValue,
					CodeLines:         code,
					CodeLineHighlight: codeLineHighlight,
					CodeLineURL:       getCodeURL(cfg, path, &lineNumber),
					LineNumber:        lineNumber,
				})
			}

			if flagIndexes.isDefaultValue {
				flagType, defaultValue_ := GetFlagType(keyOrDefaultValue)

				if variable == "" || keyOrDefaultValue == "" || defaultValue_ == "" {
					flagType = "unknown"
				}

				if variable != "" {
					defaultValueMatches = append(defaultValueMatches, MatchVariableDefaultValue{
						Variable:         variable,
						FlagDefaultValue: keyOrDefaultValue,
						FlagType:         flagType,
					})
				}

			}
		}
	}

	for _, flagIndex := range flagIndexesNotSplit {
		// Extract the code with a certain number of lines
		defaultValue_ := ""
		flagType := "unknown"
		firstLineIndex := getSurroundingLineIndex(fileContentStr, flagIndex[0], true, cfg.NbLineCodeEdges)
		lastLineIndex := getSurroundingLineIndex(fileContentStr, flagIndex[1], false, cfg.NbLineCodeEdges)
		code := fileContentStr[firstLineIndex:lastLineIndex]
		key := fileContentStr[flagIndex[0]:flagIndex[1]]

		if len(flagIndex) >= 3 {
			defaultValue := fileContentStr[flagIndex[2]:flagIndex[3]]
			flagType, defaultValue_ = GetFlagType(defaultValue)
		}

		// Better value wrapper for code highlighting (5 chars wrapping)
		keyWrapper := key
		nbCharsWrapping := 5
		if flagIndex[0] > nbCharsWrapping && flagIndex[1] < len(fileContentStr)-nbCharsWrapping {
			keyWrapper = fileContentStr[flagIndex[0]-nbCharsWrapping : flagIndex[1]+nbCharsWrapping]
		}

		if key == "" || defaultValue_ == "" {
			flagType = "unknown"
		}

		lineNumber := getLineFromPos(fileContentStr, flagIndex[0])
		codeLineHighlight := getLineFromPos(code, strings.Index(code, keyWrapper))

		if key != "" {
			resultsNotSplit = append(resultsNotSplit, model.SearchResult{
				FlagKey:           key,
				FlagDefaultValue:  defaultValue_,
				FlagType:          flagType,
				CodeLines:         code,
				CodeLineHighlight: codeLineHighlight,
				CodeLineURL:       getCodeURL(cfg, path, &lineNumber),
				LineNumber:        lineNumber,
			})
		}

	}

	resultsSplit := matchFlagWithDefaultValue(flagMatches, defaultValueMatches)
	combinedResults := append(resultsNotSplit, resultsSplit...)

	res, _ := json.Marshal(combinedResults)
	fmt.Println(string(res))

	resultChannel <- model.FileSearchResult{
		File:    path,
		FileURL: getCodeURL(cfg, path, nil),
		Results: combinedResults,
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

func keyWrapper(key, fileContentStr string, flagIndex []int) string {
	keyWrapper := key
	nbCharsWrapping := 5
	if flagIndex[2] > nbCharsWrapping && flagIndex[3] < len(fileContentStr)-nbCharsWrapping {
		keyWrapper = fileContentStr[flagIndex[2]-nbCharsWrapping : flagIndex[3]+nbCharsWrapping]
	}
	return keyWrapper
}

func flagIndexSplitter(fileContentStr, path string, regexesSplit []RegexStruct) []FlagIndexesStruct {
	flagIndexesSplits := []FlagIndexesStruct{}
	for _, regexSplit := range regexesSplit {
		for _, regex := range regexSplit.regex {
			var indexSplit [][]int
			regxp := regexp.MustCompile(regex)
			flagLineIndexes := regxp.FindAllStringIndex(fileContentStr, -1)
			for _, flagLineIndex := range flagLineIndexes {
				submatch := fileContentStr[flagLineIndex[0]:flagLineIndex[1]]
				submatchIndexes := regxp.FindAllStringSubmatchIndex(submatch, -1)

				for _, submatchIndex := range submatchIndexes {
					if len(submatchIndex) < 3 {
						log.WithFields(log.Fields{
							"reason": fmt.Sprintf("Did not find the variable in file %s. Code: %s", path, submatch),
						}).Error("Key not found")
						continue
					}

					if len(submatchIndex) < 6 {
						log.WithFields(log.Fields{
							"reason": fmt.Sprintf("Did not find the flag key or default value in file %s. Code: %s", path, submatch),
						}).Warn("Type unknown")
						indexSplit = append(indexSplit, []int{
							flagLineIndex[0] + submatchIndex[2],
							flagLineIndex[0] + submatchIndex[3],
						})
						continue
					}

					indexSplit = append(indexSplit, []int{
						flagLineIndex[0] + submatchIndex[2],
						flagLineIndex[0] + submatchIndex[3],
						flagLineIndex[0] + submatchIndex[4],
						flagLineIndex[0] + submatchIndex[5],
					})
				}
			}

			flagIndexesSplits = append(flagIndexesSplits, FlagIndexesStruct{
				isFlag:         regexSplit.isFlag,
				isDefaultValue: regexSplit.isDefaultValue,
				FlagIndexes:    indexSplit,
			})
		}
	}

	return flagIndexesSplits
}

func matchFlagWithDefaultValue(flagMatches []MatchVariableFlag, defaultValueMatches []MatchVariableDefaultValue) []model.SearchResult {
	nameMap := make(map[string]MatchVariableDefaultValue)
	results := []model.SearchResult{}

	for _, p2 := range defaultValueMatches {
		nameMap[p2.Variable] = p2
	}

	for _, p1 := range flagMatches {

		searchResult := model.SearchResult{
			FlagKey:           p1.FlagKey,
			FlagDefaultValue:  "",
			FlagType:          "unknown",
			CodeLines:         p1.CodeLines,
			CodeLineHighlight: p1.CodeLineHighlight,
			CodeLineURL:       p1.CodeLineURL,
			LineNumber:        p1.LineNumber,
		}

		if p2, found := nameMap[p1.Variable]; found {
			searchResult.FlagType = p2.FlagType
			searchResult.FlagDefaultValue = p2.FlagDefaultValue
		}

		if p1.FlagKey != "" {
			results = append(results, searchResult)
		}

	}

	return results
}
