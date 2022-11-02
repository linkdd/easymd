package server

import (
	"os"
	"path/filepath"
	"strings"
)

func hasStaticFile(rootDocument string, urlPath string) (string, bool) {
	documentPath := filepath.Join(rootDocument, strings.TrimSuffix(urlPath, "/"))

	if _, err := os.Stat(documentPath); !os.IsNotExist(err) {
		return documentPath, true
	}

	return "", false
}
