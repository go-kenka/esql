// Code generated by esql, DO NOT EDIT.
package role

import (
	"entgo.io/ent/dialect/sql"
	"github.com/go-kenka/esql"
)

const (
	TableName      = "role"
	ColumnId       = "id"
	ColumnRoleName = "role_name"
	// EdgeUserTableName user
	EdgeUserTableName       = "user"
	EdgeUserLinkField       = "id"
	EdgeUserRefField        = "role_id"
	EdgeUserDisplayNikeName = "nike_name"
)

var (
	roleTable     = sql.Table(TableName).As("t1")
	edgeUserTable = sql.Table(EdgeUserTableName).As("t2")
)

var Columns = []string{
	ColumnId,
	ColumnRoleName,
}

type RoleClient struct {
	direct string
	db     esql.Driver
}

type RoleData struct {
	UserList []*RoleEdgeUserData

	Id       int    `db:"id"`        // 角色ID
	RoleName string `db:"role_name"` // 角色名称
}

func (d *RoleData) HasUser() bool {
	return d.UserList != nil
}

type RoleEdgeUserData struct {
	RoleId   int    `db:"role_id"`   // role_id
	NikeName string `db:"nike_name"` //
}

func NewRoleClient(db esql.Driver) *RoleClient {
	return &RoleClient{
		direct: db.DriverName(),
		db:     db,
	}
}

func (c *RoleClient) Query() *RoleQuery {
	var cols []string
	for _, column := range Columns {
		cols = append(cols, roleTable.C(column))
	}
	return &RoleQuery{
		selector: sql.Dialect(c.direct).Select(cols...).From(roleTable),
		db:       c.db,
		with:     map[string]struct{}{},
	}
}

func (c *RoleClient) Create() *RoleCreate {
	return &RoleCreate{
		builder: sql.Dialect(c.direct).Insert(TableName),
		db:      c.db,
		data:    &RoleData{},
	}
}

func (c *RoleClient) CreateBulk(data ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{
		db:   c.db,
		data: data,
	}
}

func (c *RoleClient) Update() *RoleUpdate {
	return &RoleUpdate{
		builder: sql.Dialect(c.direct).Update(TableName),
		db:      c.db,
		data:    &RoleData{},
	}
}

func (c *RoleClient) UpdateOne(id int) *RoleUpdateOne {
	return &RoleUpdateOne{
		builder: sql.Dialect(c.direct).Update(TableName).Where(sql.EQ(ColumnId, id)),
		db:      c.db,
		data:    &RoleData{},
	}
}

func (c *RoleClient) Delete() *RoleDelete {
	return &RoleDelete{
		builder: sql.Dialect(c.direct).Delete(TableName),
		db:      c.db,
	}
}

func (c *RoleClient) DeleteOne(id int) *RoleDeleteOne {
	return &RoleDeleteOne{
		builder: sql.Dialect(c.direct).Delete(TableName).Where(sql.EQ(ColumnId, id)),
		db:      c.db,
	}
}
