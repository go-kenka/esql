// Code generated by esql, DO NOT EDIT.
package role

import (
	"context"
	"github.com/go-kenka/esql"
	"github.com/jmoiron/sqlx"
)

type RoleDelete struct {
	builder *esql.DeleteBuilder
	db      *sqlx.DB
}

func (d *RoleDelete) Where(p *esql.Predicate) *RoleDelete {
	d.builder.Where(p)
	return d
}

func (d *RoleDelete) Exec(ctx context.Context) (int, error) {
	return d.sqlSave(ctx)
}

func (d *RoleDelete) sqlSave(ctx context.Context) (int, error) {
	query, args := d.builder.Query()
	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	aff, _ := result.RowsAffected()
	return int(aff), nil
}

type RoleDeleteOne struct {
	builder *esql.DeleteBuilder
	db      *sqlx.DB
}

func (d *RoleDeleteOne) Save(ctx context.Context) error {
	return d.sqlSave(ctx)
}

func (d *RoleDeleteOne) sqlSave(ctx context.Context) error {
	query, args := d.builder.Query()
	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}