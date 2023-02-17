// Code generated by esql, DO NOT EDIT.
package role

import (
	"context"
	"github.com/go-kenka/esql"
	"github.com/jmoiron/sqlx"
)

type RoleCreate struct {
	builder  *esql.InsertBuilder
	selector *esql.Selector
	db       *sqlx.DB
	data     *RoleData
}

func (c *RoleCreate) Set(column string, v any) *RoleCreate {
	c.builder.Set(column, v)
	return c
}

func (c *RoleCreate) Save(ctx context.Context) (*RoleData, error) {
	id, err := c.sqlSave(ctx)
	if err != nil {
		return nil, err
	}
	return c.get(ctx, id)
}

func (c *RoleCreate) sqlSave(ctx context.Context) (int, error) {
	query, args := c.builder.Query()
	result, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}

func (c *RoleCreate) sql() (string, []any) {
	return c.builder.Query()
}

func (c *RoleCreate) get(ctx context.Context, id int) (*RoleData, error) {
	query, args := c.selector.Where(esql.EQ(ColumnId, id)).Query()
	var data RoleData

	err := c.db.GetContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type RoleCreateBulk struct {
	db       *sqlx.DB
	selector *esql.Selector
	data     []*RoleCreate
}

func (cb RoleCreateBulk) Save(ctx context.Context) ([]*RoleData, error) {
	ids, err := cb.sqlSave(ctx)
	if err != nil {
		return nil, err
	}
	return cb.find(ctx, ids)
}
func (cb RoleCreateBulk) sqlSave(ctx context.Context) ([]any, error) {
	var ids []any
	var stmt *sqlx.Stmt
	var err error
	for i, d := range cb.data {
		query, args := d.sql()
		if i == 0 {
			stmt, err = cb.db.Preparex(query)
			if err != nil {
				return nil, err
			}
		}
		if stmt != nil {
			result, err := stmt.ExecContext(ctx, args...)
			if err != nil {
				return nil, err
			}
			id, _ := result.LastInsertId()
			ids = append(ids, id)
		}
	}
	return ids, nil
}

func (cb RoleCreateBulk) find(ctx context.Context, ids []any) ([]*RoleData, error) {
	query, args := cb.selector.Where(esql.In(ColumnId, ids...)).Query()
	var data []*RoleData
	err := cb.db.SelectContext(ctx, data, query, args...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

