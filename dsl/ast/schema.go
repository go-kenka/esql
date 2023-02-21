package ast

import (
	"github.com/go-kenka/esql/gen"
	"os"
	"path/filepath"
	"strings"
)

func ReadDir(path string) []*gen.Table {
	var tbs []*gen.Table

	files, _ := os.ReadDir(path)

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}

		defBytes, err := os.ReadFile(filepath.Join(path, entry.Name()))
		if err != nil {
			return nil
		}
		tbs = append(tbs, astReadFile(entry.Name(), string(defBytes)))
	}

	return tbs
}
