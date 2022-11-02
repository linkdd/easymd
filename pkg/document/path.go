package document

import (
	"path/filepath"
	"strings"
)

func GetTestPaths(rootDocument string, urlPath string) []string {
	documentPath := strings.TrimSuffix(urlPath, "/")
	var testPaths []string

	if documentPath == "" {
		testPaths = []string{
			filepath.Join(rootDocument, "_index.md"),
		}
	} else {
		testPaths = []string{
			filepath.Join(rootDocument, documentPath, "_index.md"),
			filepath.Join(rootDocument, documentPath+".md"),
		}
	}

	return testPaths
}
