package dsl

type Fn func(f *FieldExpr)

type FieldExpr struct {
	Tag      string      // 生成go结构体，附加的tag内容
	Name     string      // 字段名称
	Size     int         // 大小
	TypeInfo Type        // 字段类型
	Unique   bool        // 是否唯一
	Nillable bool        // 是否为NULL
	Default  interface{} // 默认值
	Comment  string      // 备注
}

func Field(name string, fns ...Fn) *FieldExpr {
	f := &FieldExpr{
		Name: name,
	}
	for _, o := range fns {
		o(f)
	}
	return f
}

func Tag(tag string) Fn {
	return func(f *FieldExpr) {
		f.Tag = tag
	}
}

func Size(size int) Fn {
	return func(f *FieldExpr) {
		f.Size = size
	}
}

func TypeInfo(typeInfo Type) Fn {
	return func(f *FieldExpr) {
		f.TypeInfo = typeInfo
	}
}

func Unique(unique bool) Fn {
	return func(f *FieldExpr) {
		f.Unique = unique
	}
}

func Nillable(nillable bool) Fn {
	return func(f *FieldExpr) {
		f.Nillable = nillable
	}
}

func Default(def interface{}) Fn {
	return func(f *FieldExpr) {
		f.Default = def
	}
}

func Comment(com string) Fn {
	return func(f *FieldExpr) {
		f.Comment = com
	}
}
