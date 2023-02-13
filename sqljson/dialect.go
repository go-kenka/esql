// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqljson

import (
	"fmt"
	"reflect"

	"github.com/go-kenka/esql"
)

type sqlite struct{}

// Append implements the driver.Append method.
func (d *sqlite) Append(u *esql.UpdateBuilder, column string, elems []any, opts ...Option) {
	setCase(u, column, when{
		Cond: func(b *esql.Builder) {
			typ := func(b *esql.Builder) *esql.Builder {
				return b.WriteString("JSON_TYPE").Wrap(func(b *esql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).mysqlPath(b)
				})
			}
			typ(b).WriteOp(esql.OpIsNull)
			b.WriteString(" OR ")
			typ(b).WriteOp(esql.OpEQ).WriteString("'null'")
		},
		Then: func(b *esql.Builder) {
			if len(opts) > 0 {
				b.WriteString("JSON_SET").Wrap(func(b *esql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).mysqlPath(b)
					b.Comma().Argf("JSON(?)", marshalArg(elems))
				})
			} else {
				b.Arg(marshalArg(elems))
			}
		},
		Else: func(b *esql.Builder) {
			b.WriteString("JSON_INSERT").Wrap(func(b *esql.Builder) {
				b.Ident(column).Comma()
				// If no path was provided the top-level value is
				// a JSON array. i.e. JSON_INSERT(c, '$[#]', ?).
				path := func(b *esql.Builder) { b.WriteString("'$[#]'") }
				if len(opts) > 0 {
					p := identPath(column, opts...)
					p.Path = append(p.Path, "[#]")
					path = p.mysqlPath
				}
				for i, e := range elems {
					if i > 0 {
						b.Comma()
					}
					path(b)
					b.Comma()
					d.appendArg(b, e)
				}
			})
		},
	})
}

func (d *sqlite) appendArg(b *esql.Builder, v any) {
	switch {
	case !isPrimitive(v):
		b.Argf("JSON(?)", marshalArg(v))
	default:
		b.Arg(v)
	}
}

type mysql struct{}

// Append implements the driver.Append method.
func (d *mysql) Append(u *esql.UpdateBuilder, column string, elems []any, opts ...Option) {
	setCase(u, column, when{
		Cond: func(b *esql.Builder) {
			typ := func(b *esql.Builder) *esql.Builder {
				b.WriteString("JSON_TYPE(JSON_EXTRACT(")
				b.Ident(column).Comma()
				identPath(column, opts...).mysqlPath(b)
				return b.WriteString("))")
			}
			typ(b).WriteOp(esql.OpIsNull)
			b.WriteString(" OR ")
			typ(b).WriteOp(esql.OpEQ).WriteString("'NULL'")
		},
		Then: func(b *esql.Builder) {
			if len(opts) > 0 {
				b.WriteString("JSON_SET").Wrap(func(b *esql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).mysqlPath(b)
					b.Comma().WriteString("JSON_ARRAY(").Args(d.marshalArgs(elems)...).WriteByte(')')
				})
			} else {
				b.WriteString("JSON_ARRAY(").Args(d.marshalArgs(elems)...).WriteByte(')')
			}
		},
		Else: func(b *esql.Builder) {
			b.WriteString("JSON_ARRAY_APPEND").Wrap(func(b *esql.Builder) {
				b.Ident(column).Comma()
				for i, e := range elems {
					if i > 0 {
						b.Comma()
					}
					identPath(column, opts...).mysqlPath(b)
					b.Comma()
					d.appendArg(b, e)
				}
			})
		},
	})
}

func (d *mysql) marshalArgs(args []any) []any {
	vs := make([]any, len(args))
	for i, v := range args {
		if !isPrimitive(v) {
			v = marshalArg(v)
		}
		vs[i] = v
	}
	return vs
}

func (d *mysql) appendArg(b *esql.Builder, v any) {
	switch {
	case !isPrimitive(v):
		b.Argf("CAST(? AS JSON)", marshalArg(v))
	default:
		b.Arg(v)
	}
}

type postgres struct{}

// Append implements the driver.Append method.
func (*postgres) Append(u *esql.UpdateBuilder, column string, elems []any, opts ...Option) {
	setCase(u, column, when{
		Cond: func(b *esql.Builder) {
			valuePath(b, column, append(opts, Cast("jsonb"))...)
			b.WriteOp(esql.OpIsNull)
			b.WriteString(" OR ")
			valuePath(b, column, append(opts, Cast("jsonb"))...)
			b.WriteOp(esql.OpEQ).WriteString("'null'::jsonb")
		},
		Then: func(b *esql.Builder) {
			if len(opts) > 0 {
				b.WriteString("jsonb_set").Wrap(func(b *esql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).pgArrayPath(b)
					b.Comma().Arg(marshalArg(elems))
					b.Comma().WriteString("true")
				})
			} else {
				b.Arg(marshalArg(elems))
			}
		},
		Else: func(b *esql.Builder) {
			if len(opts) > 0 {
				b.WriteString("jsonb_set").Wrap(func(b *esql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).pgArrayPath(b)
					b.Comma()
					path := identPath(column, opts...)
					path.value(b)
					b.WriteString(" || ").Arg(marshalArg(elems))
					b.Comma().WriteString("true")
				})
			} else {
				b.Ident(column).WriteString(" || ").Arg(marshalArg(elems))
			}
		},
	})
}

// driver groups all dialect-specific methods.
type driver interface {
	Append(u *esql.UpdateBuilder, column string, elems []any, opts ...Option)
}

func newDriver(name string) (driver, error) {
	switch name {
	case esql.SQLite:
		return (*sqlite)(nil), nil
	case esql.MySQL:
		return (*mysql)(nil), nil
	case esql.Postgres:
		return (*postgres)(nil), nil
	default:
		return nil, fmt.Errorf("sqljson: unknown driver %q", name)
	}
}

type when struct{ Cond, Then, Else func(*esql.Builder) }

// setCase sets the column value using the "CASE WHEN" statement.
// The x defines the condition/predicate, t is the true (if) case,
// and 'f' defines the false (else).
func setCase(u *esql.UpdateBuilder, column string, w when) {
	u.Set(column, esql.ExprFunc(func(b *esql.Builder) {
		b.WriteString("CASE WHEN ").Wrap(func(b *esql.Builder) {
			w.Cond(b)
		})
		b.WriteString(" THEN ")
		w.Then(b)
		b.WriteString(" ELSE ")
		w.Else(b)
		b.WriteString(" END")
	}))
}

func isPrimitive(v any) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct, reflect.Ptr, reflect.Interface:
		return false
	}
	return true
}
