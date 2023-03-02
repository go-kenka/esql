package dsl

type EdgeExpr struct {
	Name     string
	Type     EdgeType
	Link     string
	From     string
	Ref      string
	Display  []*FieldExpr
	Relation []*EdgeExpr
}

type EdgeFn func(e *EdgeExpr)

func Edge(name string, fns ...EdgeFn) *EdgeExpr {
	t := &EdgeExpr{
		Name: name,
	}

	for _, fn := range fns {
		fn(t)
	}
	return t
}

func Link(l string) EdgeFn {
	return func(e *EdgeExpr) {
		e.Link = l
	}
}

func From(f string) EdgeFn {
	return func(e *EdgeExpr) {
		e.From = f
	}
}

func Ref(r string) EdgeFn {
	return func(e *EdgeExpr) {
		e.Ref = r
	}
}

func EType(t EdgeType) EdgeFn {
	return func(e *EdgeExpr) {
		e.Type = t
	}
}

func Display(d ...*FieldExpr) EdgeFn {
	return func(e *EdgeExpr) {
		e.Display = append(e.Display, d...)
	}
}

func Relation(r ...*EdgeExpr) EdgeFn {
	return func(e *EdgeExpr) {
		e.Relation = append(e.Relation, r...)
	}
}
