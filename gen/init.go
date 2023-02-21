package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func Init(base, name string) error {
	dir := filepath.Join(base, "ast")
	genFile := filepath.Join(fmt.Sprintf("%s/%s.go", dir, name))

	// 生成之前，先删除文件
	os.Remove(genFile)

	tmp := template.New("init.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"goType":    GoType,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/init.tmpl")
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
	defer fs.Close()

	return tmp.Execute(fs, name)
}
