# `Esql`是一个简单的SQL生成工具

通过内置的`dsl`语法，实现数据库`sql`和实体之间的映射和相关`CRUD`的实现
实现本身通过解析`dsl`语法树，解析数据表的相关定义，然后使用`template`实现模板的生成。

抽离了[ent](entgo.io/ent)的builder，并结合[sqlx](github.com/jmoiron/sqlx)实现数据库的相关操作。

# 快速开始
```shell
go install github.com/go-kenka/esql/cmd/esql@latest
```

# 生成schema文件
```shell
esql init user --target ./data
```

## 示例
```go
package schema

import (
	. "github.com/go-kenka/esql/dsl"
)

var _ = Table("user",
	Desc("用户表"),
	Fields(
		Field("id",
			Tag("db:\"id\""),
			TypeInfo(TypeInt),
			Unique(true),
			Nillable(false),
			Default(0),
			Comment("ID"),
		),
		Field("username",
			Tag("db:\"username\""),
			TypeInfo(TypeString),
			Unique(true),
			Nillable(false),
			Default(0),
			Comment("用户账号"),
		),
		Field("nike_name",
			Tag("db:\"nike_name\""),
			TypeInfo(TypeString),
			Unique(false),
			Nillable(false),
			Default([]string{"aaa"}),
			Comment("用户名称"),
		),
		Field("role_id",
			Tag("db:\"role_id\""),
			TypeInfo(TypeInt),
			Unique(false),
			Nillable(false),
			Default([]string{"aaa"}),
			Comment("角色ID"),
		),
	),
	Edges(
		Edge("用户与角色关系",
			Link("role_id"),
			From("role"),
			Ref("id"),
			EType(TypeM2O),
			Display(
				Field("role_name",
					Tag("db:\"role_name\""),
					TypeInfo(TypeString),
				),
			),
		),
	),
)

```


# 生成CRUD文件
```shell
esql gen ./data/schema --target ./data
```

## 示例结果

### 目录结构
```text
│  client.go
│  
├─role
│      role.go
│      role_create.go
│      role_delete.go
│      role_query.go
│      role_update.go
│      
├─schema
│      role.go
│      user.go
│      
└─user
        user.go
        user_create.go
        user_delete.go
        user_query.go
        user_update.go
```

### client.go

```go
// Code generated by esql, DO NOT EDIT.
package sql

import (
	"fmt"
	"github.com/go-kenka/esql"
	"github.com/go-kenka/esql/examples/data/role"
	"github.com/go-kenka/esql/examples/data/user"
	"github.com/jmoiron/sqlx"
)

// Client .
type Client struct {
	Role *role.RoleClient
	User *user.UserClient
}

// NewClient .
func NewClient(db *sqlx.DB) *Client {
	return &Client{
		Role: role.NewRoleClient(db),
		User: user.NewUserClient(db),
	}
}

// Open .
func Open(driverName, dataSourceName string) (*Client, error) {
	switch driverName {
	case esql.MySQL, esql.Postgres, esql.SQLite:
		db, err := sqlx.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(db), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

```

### 使用示例
```go
package examples

import (
	"context"
	"fmt"
	sql "github.com/go-kenka/esql/examples/data"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := sql.Open("mysql", "root:xxxx!@tcp(192.168.0.12:3306)/risk-sensor?parseTime=true")
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
```