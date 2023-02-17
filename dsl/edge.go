package dsl

type EdgeExpr struct {
	Desc    string
	Link    string
	From    string
	To      string
	Ref     string
	Display []*FieldExpr
}

type EdgeFn func(e *EdgeExpr)

func Edge(desc string, fns ...EdgeFn) *EdgeExpr {
	t := &EdgeExpr{
		Desc: desc,
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

func To(t string) EdgeFn {
	return func(e *EdgeExpr) {
		e.To = t
	}
}

func Ref(r string) EdgeFn {
	return func(e *EdgeExpr) {
		e.Ref = r
	}
}

func Display(d *FieldExpr) EdgeFn {
	return func(e *EdgeExpr) {
		e.Display = append(e.Display, d)
	}
}
