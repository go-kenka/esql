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

		tb := astReadFile(entry.Name(), string(defBytes))

		var id *gen.Field
		var index int
		for i, field := range tb.Fields {
			if field.TypeInfo == gen.TypeString && field.Size == 0 {
				field.Size = 255
			}
			if field.TypeInfo == gen.TypeEnum && field.Size == 0 {
				field.Size = 50
			}
			if field.Name == "id" {
				id = field
				index = i
			}
		}

		// 如果没有ID，需要添加ID
		if id == nil {
			tb.Fields = append([]*gen.Field{
				{
					Name:     "id",
					TypeInfo: gen.TypeInt,
					Comment:  "Primary key",
				},
			}, tb.Fields...)
		} else {
			//	有ID的，需要将ID调整到首位
			fields := append(tb.Fields[:index:index], tb.Fields[index+1:]...)
			tb.Fields = append([]*gen.Field{id}, fields...)
		}

		tbs = append(tbs, tb)
	}

	return tbs
}
