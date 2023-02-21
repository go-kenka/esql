package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func genMigrate(base string, data *Client) error {
	dir := filepath.Join(base, "migrate")
	genFile := filepath.Join(fmt.Sprintf("%s/migrate.go", dir))

	// 生成之前，先删除文件
	os.Remove(genFile)

	tmp := template.New("migrate.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"goType":    GoType,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/migrate.tmpl")
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

	return tmp.Execute(fs, data)
}

func genSchema(base string, data *Client) error {
	dir := filepath.Join(base, "migrate")
	genFile := filepath.Join(fmt.Sprintf("%s/schema.go", dir))

	// 生成之前，先删除文件
	os.Remove(genFile)

	tmp := template.New("schema.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"dbType":    DBType,
		"isString":  IsString,
		"lower":     Lower,
	})
	tmp, err := tmp.ParseFS(tmpl, "template/schema.tmpl")
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

	return tmp.Execute(fs, data)
}
