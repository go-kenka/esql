// Code generated by esql, DO NOT EDIT.
package user

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kenka/esql"
)

type UserDelete struct {
	builder *sql.DeleteBuilder
	db      esql.Driver
}

func (d *UserDelete) Where(p *sql.Predicate) *UserDelete {
	d.builder.Where(p)
	return d
}

func (d *UserDelete) Exec(ctx context.Context) (int, error) {
	return d.sqlSave(ctx)
}

func (d *UserDelete) sqlSave(ctx context.Context) (int, error) {
	query, args := d.builder.Query()
	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	aff, _ := result.RowsAffected()
	return int(aff), nil
}

type UserDeleteOne struct {
	builder *sql.DeleteBuilder
	db      esql.Driver
}

func (d *UserDeleteOne) Save(ctx context.Context) error {
	return d.sqlSave(ctx)
}

func (d *UserDeleteOne) sqlSave(ctx context.Context) error {
	query, args := d.builder.Query()
	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
