package gen

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
)

type Client struct {
	Tables []*Table
}

//go:embed template/*
var tmpl embed.FS

func GenClient(base string, tbs []*Table) error {

	err := os.MkdirAll(base, os.ModePerm)
	if err != nil {
		return err
	}

	genFile := filepath.Join(base, "client.go")

	os.Remove(genFile)

	fmt.Printf("生成client的数据成功\n")

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

	data := Client{
		Tables: tbs,
	}

	fs, err := os.OpenFile(genFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	err = tmp.Execute(fs, &data)
	if err != nil {
		return err
	}

	return nil
}

func getPkg() string {
	_, pkg := filepath.Split(reflect.TypeOf(Client{}).PkgPath())
	return pkg
}
