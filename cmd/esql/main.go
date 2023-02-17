package main

import (
	"context"
	"fmt"
	sql "github.com/go-kenka/esql/cmd/esql/data"
	"github.com/go-kenka/esql/gen"
	"github.com/go-kenka/esql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//g()
	use()
}

func g() {
	tbs := schema.ReadDir("cmd/esql/data/schema")
	for _, tb := range tbs {
		fmt.Println(tb.Name)
	}

	err := gen.GenClient("cmd/esql/data", tbs)
	if err != nil {
		panic(err)
	}
}

func use() {
	db, err := sql.Open("mysql", "root:Jianshu2020!@tcp(192.168.0.219:3306)/data-factory?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	data, err := db.User.Query().WithRole().First(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", data)
}
