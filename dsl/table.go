package dsl

type TableExpr struct {
	Name   string       // 表名称
	Fields []*FieldExpr // 表字段集合
	Desc   string       // 备注
	Edges  []*EdgeExpr  // 关系
}

type TableFn func(t *TableExpr)

func Table(name string, fns ...TableFn) *TableExpr {
	t := &TableExpr{
		Name: name,
	}

	for _, fn := range fns {
		fn(t)
	}
	return t
}

func Desc(desc string) TableFn {
	return func(t *TableExpr) {
		t.Desc = desc
	}
}
func Edges(e ...*EdgeExpr) TableFn {
	return func(t *TableExpr) {
		t.Edges = append(t.Edges, e...)
	}
}
func Fields(fs ...*FieldExpr) TableFn {
	return func(t *TableExpr) {
		t.Fields = append(t.Fields, fs...)
	}
}
