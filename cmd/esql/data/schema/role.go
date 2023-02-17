package schema

import (
	. "github.com/go-kenka/esql/dsl"
)

var _ = Table("role",
	Desc("角色表"),
	Fields(
		Field("id",
			Tag("db:\"id\""),
			TypeInfo(TypeInt),
			Unique(true),
			Nillable(false),
			Default(0),
			Comment("角色ID"),
		),
		Field("role_name",
			Tag("db:\"role_name\""),
			TypeInfo(TypeString),
			Size(20),
			Unique(false),
			Nillable(false),
			Default(""),
			Comment("角色名称"),
		),
	),
	Edges(
		Edge("角色与用户关系（1对n）",
			Link("id"),
			From("user"),
			Ref("role_id"),
			Display(
				Field("nike_name",
					Tag("db:\"nike_name\""),
					TypeInfo(TypeString),
				),
			),
		),
	),
)
