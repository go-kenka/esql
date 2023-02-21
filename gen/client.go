package gen

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Client struct {
	Tables []*Table
	Pkg    string
}

//go:embed template/*
var tmpl embed.FS

func GenClient(base, pkg string, tbs []*Table) error {

	err := os.MkdirAll(base, os.ModePerm)
	if err != nil {
		return err
	}

	genFile := filepath.Join(base, "client.go")

	os.Remove(genFile)

	for i, tb := range tbs {
		fmt.Printf("正在生成第%d【%s】个表的数据\n", i+1, tb.Name)
		if err := genData(base, tb); err != nil {
			return err
		}
		if err := genCreate(base, tb); err != nil {
			return err
		}
		if err := genDelete(base, tb); err != nil {
			return err
		}
		if err := genQuery(base, tb); err != nil {
			return err
		}
		if err := genUpdate(base, tb); err != nil {
			return err
		}
		fmt.Printf("正在生成第%d个表的数据生成成功\n", i+1)
	}

	data := Client{
		Tables: tbs,
		Pkg:    pkg,
	}
	fmt.Printf("正在生成migrate的数据\n")
	if err := genMigrate(base, &data); err != nil {
		return err
	}
	if err := genSchema(base, &data); err != nil {
		return err
	}
	fmt.Printf("生成migrate的数据成功\n")

	fmt.Printf("正在生成client的数据\n")

	tmp := template.New("client.tmpl")
	tmp.Funcs(template.FuncMap{
		"camelCase": CamelCase,
		"goType":    GoType,
		"lower":     Lower,
	})
	tmp, err = tmp.ParseFS(tmpl, "template/client.tmpl")
	if err != nil {
		return err
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fs.Close()

	err = tmp.Execute(fs, &data)
	if err != nil {
		return err
	}

	fmt.Printf("生成client的数据成功\n")
	return nil
}
