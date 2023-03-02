// Code generated by esql, DO NOT EDIT.
package user

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kenka/esql"
)

type UserQuery struct {
	*sql.Selector
	db   esql.Driver
	with map[string]struct{}
}

func (q *UserQuery) First(ctx context.Context) (*UserData, error) {
	query, args := q.Limit(1).Query()
	var data UserData
	err := q.db.GetContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}

	err = q.queryWith(ctx, []*UserData{&data})
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (q *UserQuery) FirstID(ctx context.Context) (int, error) {
	query, args := q.Select(ColumnId).Limit(1).Query()
	var id int
	err := q.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (q *UserQuery) IDs(ctx context.Context) ([]int, error) {
	query, args := q.Select(ColumnId).Limit(1).Query()
	rows, err := q.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		data = append(data, id)
	}

	return data, nil
}

func (q *UserQuery) ScanX(ctx context.Context, dist any) error {
	query, args := q.Query()
	err := q.db.SelectContext(ctx, &dist, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQuery) AllX(ctx context.Context) ([]*UserData, error) {
	query, args := q.Query()
	var data []*UserData
	err := q.db.SelectContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}

	err = q.queryWith(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q *UserQuery) CountX(ctx context.Context) (int, error) {
	query, args := q.Count(ColumnId).Query()
	var count int
	err := q.db.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *UserQuery) ExistX(ctx context.Context) (bool, error) {
	query, args := q.Count(ColumnId).Query()
	var count int
	err := q.db.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (q *UserQuery) WithRole() *UserQuery {
	// 添加Display字段
	q.AppendSelect(edgeRoleTable.C(EdgeRoleDisplayRoleName))
	// 添加关系（左连接）
	q.LeftJoin(edgeRoleTable).
		On(
			q.C(EdgeRoleLinkField),
			edgeRoleTable.C(EdgeRoleRefField),
		)
	// 添加Display字段
	q.AppendSelect(roleEdgeAccessTable.C(RoleEdgeAccessDisplayAccessName))
	// 添加关系（左连接）
	q.LeftJoin(roleEdgeAccessTable).
		On(
			edgeRoleTable.C(RoleEdgeAccessLinkField),
			roleEdgeAccessTable.C(RoleEdgeAccessRefField),
		)
	return q
}

func (q *UserQuery) queryWith(ctx context.Context, data []*UserData) error {
	return nil
}
