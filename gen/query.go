package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func genQuery(base string, t *Table) error {
	dir := filepath.Join(base, t.Name)
	genFile := filepath.Join(fmt.Sprintf("%s/%s_query.go", dir, t.Name))

	// 生成之前，先删除文件
	os.Remove(genFile)

	tmp := template.New("query.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"goType":    GoType,
		"lower":     Lower,
		"withCheck": WithCheck,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/query.tmpl")
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fs.Close()

	return tmp.Execute(fs, t)
}
