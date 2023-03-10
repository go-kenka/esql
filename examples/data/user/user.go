// Code generated by esql, DO NOT EDIT.
package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/go-kenka/esql"
)

const (
	TableName      = "user"
	ColumnId       = "id"
	ColumnUsername = "username"
	ColumnNikeName = "nike_name"
	ColumnRoleId   = "role_id"
	// EdgeRoleTableName role
	EdgeRoleTableName       = "role"
	EdgeRoleLinkField       = "role_id"
	EdgeRoleRefField        = "id"
	EdgeRoleDisplayRoleName = "role_name"
	// RoleEdgeAccessTableName access
	RoleEdgeAccessTableName         = "access"
	RoleEdgeAccessLinkField         = "access_id"
	RoleEdgeAccessRefField          = "id"
	RoleEdgeAccessDisplayAccessName = "access_name"
)

var (
	userTable           = sql.Table(TableName).As("t1")
	edgeRoleTable       = sql.Table(EdgeRoleTableName).As("t2")
	roleEdgeAccessTable = sql.Table(RoleEdgeAccessTableName)
)

var Columns = []string{
	ColumnId,
	ColumnUsername,
	ColumnNikeName,
	ColumnRoleId,
}

type UserClient struct {
	direct string
	db     esql.Driver
}

type UserData struct {
	*UserEdgeRoleData

	Id       int    `db:"id"`        // ID
	Username string `db:"username"`  // 用户账号
	NikeName string `db:"nike_name"` // 用户名称
	RoleId   int    `db:"role_id"`   // 角色ID
}

func (d *UserData) HasRole() bool {
	return d.UserEdgeRoleData != nil
}

type UserEdgeRoleData struct {
	*RoleEdgeAccessData
	Id       int    `db:"id"`        // id
	RoleName string `db:"role_name"` //
}

func (d *UserEdgeRoleData) HasAccess() bool {
	return d.RoleEdgeAccessData != nil
}

type RoleEdgeAccessData struct {
	AccessName string `db:"access_name"` //
}

func NewUserClient(db esql.Driver) *UserClient {
	return &UserClient{
		direct: db.DriverName(),
		db:     db,
	}
}

func (c *UserClient) Query() *UserQuery {
	var cols []string
	for _, column := range Columns {
		cols = append(cols, userTable.C(column))
	}
	return &UserQuery{
		selector: sql.Dialect(c.direct).Select(cols...).From(userTable),
		db:       c.db,
		with:     map[string]struct{}{},
	}
}

func (c *UserClient) Create() *UserCreate {
	return &UserCreate{
		builder: sql.Dialect(c.direct).Insert(TableName),
		db:      c.db,
		data:    &UserData{},
	}
}

func (c *UserClient) CreateBulk(data ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{
		db:   c.db,
		data: data,
	}
}

func (c *UserClient) Update() *UserUpdate {
	return &UserUpdate{
		builder: sql.Dialect(c.direct).Update(TableName),
		db:      c.db,
		data:    &UserData{},
	}
}

func (c *UserClient) UpdateOne(id int) *UserUpdateOne {
	return &UserUpdateOne{
		builder: sql.Dialect(c.direct).Update(TableName).Where(sql.EQ(ColumnId, id)),
		db:      c.db,
		data:    &UserData{},
	}
}

func (c *UserClient) Delete() *UserDelete {
	return &UserDelete{
		builder: sql.Dialect(c.direct).Delete(TableName),
		db:      c.db,
	}
}

func (c *UserClient) DeleteOne(id int) *UserDeleteOne {
	return &UserDeleteOne{
		builder: sql.Dialect(c.direct).Delete(TableName).Where(sql.EQ(ColumnId, id)),
		db:      c.db,
	}
}
