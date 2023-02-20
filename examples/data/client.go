// Code generated by esql, DO NOT EDIT.
package sql

import (
	"fmt"
	"github.com/go-kenka/esql"
	"github.com/go-kenka/esql/data/role"
	"github.com/go-kenka/esql/data/user"
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