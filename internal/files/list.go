package files

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/thoas/go-funk"
)

// ListFiles list all the files in the selected directory that does not match exluded patterns
func ListFiles(dir string, toExclude []string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			// Check if any shouldExclude expression matches file path
			shouldExclude := funk.Find(toExclude, func(exclude string) bool {
				matched, _ := regexp.Match(exclude, []byte(path))
				return matched
			})

			if shouldExclude == nil {
				files = append(files, path)
			}
			return nil
		})

	return files, err
}
