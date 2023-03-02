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
		Edge("user",
			Link("id"),
			From("user"),
			Ref("role_id"),
			EType(TypeO2M),
			Display(
				Field("nike_name",
					Tag("db:\"nike_name\""),
					TypeInfo(TypeString),
				),
			),
		),
	),
)
