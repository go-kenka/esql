package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func genCreate(base string, t *Table) error {
	dir := filepath.Join(base, t.Name)
	genFile := filepath.Join(fmt.Sprintf("%s/%s_create.go", dir, t.Name))

	// 生成之前，先删除文件
	os.Remove(genFile)

	tmp := template.New("create.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"goType":    GoType,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/create.tmpl")
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	return tmp.Execute(fs, t)
}
