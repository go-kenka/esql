package main

import (
	"context"
	"fmt"
	sql "github.com/go-kenka/esql/examples/data"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	client, err := sql.Open("mysql", "root:Jianshu2020!@tcp(192.168.0.219:3306)/data-factory?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	// 查询单个
	user, err := client.User.Query().First(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Printf("result %+v", user)
}
