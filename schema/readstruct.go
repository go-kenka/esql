package schema

import (
	"errors"
	"fmt"
	"github.com/go-kenka/esql/gen"
	"go/ast"
	"go/token"
	"strconv"
)

func readTable(table *gen.Table, call *ast.CallExpr) {
	fun, ok := call.Fun.(*ast.Ident)
	//等于定义
	if ok && fun.Name == "Table" {
		for _, arg := range call.Args {
			switch a := arg.(type) {
			case *ast.BasicLit:
				table.Name = getStringValue(a)
			case *ast.CallExpr:
				readTableFn(table, a)
			}
		}
	}
}

func readTableFn(table *gen.Table, call *ast.CallExpr) {
	if fun, ok := call.Fun.(*ast.Ident); ok {
		switch fun.Name {
		case "Desc":
			readTableDesc(table, call.Args[0])
		case "Fields":
			readTableFields(table, call.Args)
		case "Edges":
			readTableEdges(table, call.Args)
		}
	}

}

func readTableDesc(table *gen.Table, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		table.Desc = getStringValue(d)
	}
}

func readTableFields(table *gen.Table, args []ast.Expr) {

	for _, arg := range args {

		fs := &gen.Field{}

		if call, ok := arg.(*ast.CallExpr); ok {
			//等于定义
			if fun, ok := call.Fun.(*ast.Ident); ok && fun.Name == "Field" {
				for _, fg := range call.Args {
					switch a := fg.(type) {
					case *ast.BasicLit:
						fs.Name = getStringValue(a)
					case *ast.CallExpr:
						readFieldFn(fs, a)
					}
				}
			}
		}

		table.Fields = append(table.Fields, fs)
	}

}

func readFieldFn(fs *gen.Field, call *ast.CallExpr) {
	if fun, ok := call.Fun.(*ast.Ident); ok {
		switch fun.Name {
		case "Tag":
			readFieldTag(fs, call.Args[0])
		case "TypeInfo":
			readFieldTypeInfo(fs, call.Args[0])
		case "Unique":
			readFieldUnique(fs, call.Args[0])
		case "Nillable":
			readFieldNillable(fs, call.Args[0])
		case "Size":
			readFieldSize(fs, call.Args[0])
		case "Comment":
			readFieldComment(fs, call.Args[0])
		case "Default":
			readFieldDefault(fs, call.Args[0])
		}
	}

}

func readFieldTag(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		fs.Tag = getStringValue(d)
	}
}

func readFieldTypeInfo(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		fs.TypeInfo = gen.TypeNameMap[d.Name]
	}
}

func readFieldUnique(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		fs.Unique = getBoolValue(d.Name)
	}
}

func readFieldNillable(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		fs.Nillable = getBoolValue(d.Name)
	}
}

func readFieldSize(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		if d.Kind == token.INT {
			fs.Size = getIntValue(d)
		}
	}
}

func readFieldComment(fs *gen.Field, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		fs.Comment = getStringValue(d)
	}
}

func readFieldDefault(fs *gen.Field, arg ast.Expr) {
	//TODO: 只解析了基本类型其他类型
	if d, ok := arg.(*ast.BasicLit); ok {
		fs.Default = getBasicValue(d)
	}
}

func readTableEdges(table *gen.Table, args []ast.Expr) {

	for _, arg := range args {

		edge := &gen.Edge{}

		if call, ok := arg.(*ast.CallExpr); ok {
			//等于定义
			if fun, ok := call.Fun.(*ast.Ident); ok && fun.Name == "Edge" {
				for _, fg := range call.Args {
					switch a := fg.(type) {
					case *ast.BasicLit:
						edge.Desc = getStringValue(a)
					case *ast.CallExpr:
						readEdgeFn(edge, a)
					}
				}
			}
		}

		table.Edges = append(table.Edges, edge)
	}

}

func readEdgeFn(edge *gen.Edge, call *ast.CallExpr) {
	if fun, ok := call.Fun.(*ast.Ident); ok {
		switch fun.Name {
		case "Link":
			readEdgeLink(edge, call.Args[0])
		case "EType":
			readEdgeType(edge, call.Args[0])
		case "From":
			readEdgeFrom(edge, call.Args[0])
		case "To":
			readEdgeTo(edge, call.Args[0])
		case "Ref":
			readEdgeRef(edge, call.Args[0])
		case "Display":
			readEdgeDisplay(edge, call.Args)
		}
	}

}

func readEdgeLink(edge *gen.Edge, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		edge.Link = getStringValue(d)
	}
}

func readEdgeType(edge *gen.Edge, arg ast.Expr) {
	if d, ok := arg.(*ast.Ident); ok {
		edge.Type = gen.EdgeTypeNameMap[d.Name]
	}
}

func readEdgeFrom(edge *gen.Edge, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		edge.From = getStringValue(d)
	}
}

func readEdgeTo(edge *gen.Edge, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		edge.To = getStringValue(d)
	}
}

func readEdgeRef(edge *gen.Edge, arg ast.Expr) {
	if d, ok := arg.(*ast.BasicLit); ok {
		edge.Ref = getStringValue(d)
	}
}

func readEdgeDisplay(edgs *gen.Edge, args []ast.Expr) {
	for _, arg := range args {
		fs := &gen.Field{}

		if call, ok := arg.(*ast.CallExpr); ok {
			//等于定义
			if fun, ok := call.Fun.(*ast.Ident); ok && fun.Name == "Field" {
				for _, fg := range call.Args {
					switch a := fg.(type) {
					case *ast.BasicLit:
						fs.Name = getStringValue(a)
					case *ast.CallExpr:
						readFieldFn(fs, a)
					}
				}
			}
		}

		edgs.Display = append(edgs.Display, fs)
	}
}

func getBasicValue(basicLit *ast.BasicLit) interface{} {
	switch basicLit.Kind {
	case token.INT:
		value, err := strconv.Atoi(basicLit.Value)
		if err != nil {
			return err
		}
		return value
	case token.STRING:
		value, err := strconv.Unquote(basicLit.Value)
		if err != nil {
			return err
		}
		return value
	}
	return errors.New(fmt.Sprintf("%s is not support type", basicLit.Kind))
}

func getStringValue(basicLit *ast.BasicLit) string {
	if basicLit.Kind == token.STRING {
		value, _ := strconv.Unquote(basicLit.Value)
		return value
	}
	return ""
}

func getIntValue(basicLit *ast.BasicLit) int {
	if basicLit.Kind == token.INT {
		value, _ := strconv.Atoi(basicLit.Value)
		return value
	}
	return 0
}

func getBoolValue(name string) bool {
	v, _ := strconv.ParseBool(name)
	return v
}
