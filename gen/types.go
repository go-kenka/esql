package gen

// A Type represents a field type.
type Type uint8

// List of field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeTime
	TypeJSON
	TypeUUID
	TypeBytes
	TypeEnum
	TypeString
	TypeOther
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint
	TypeUint64
	TypeFloat32
	TypeFloat64
	endTypes
)

var (
	TypeNames = [...]string{
		TypeInvalid: "invalid",
		TypeBool:    "bool",
		TypeTime:    "time.Time",
		TypeJSON:    "json.RawMessage",
		TypeUUID:    "[16]byte",
		TypeBytes:   "[]byte",
		TypeEnum:    "string",
		TypeString:  "string",
		TypeOther:   "other",
		TypeInt:     "int",
		TypeInt8:    "int8",
		TypeInt16:   "int16",
		TypeInt32:   "int32",
		TypeInt64:   "int64",
		TypeUint:    "uint",
		TypeUint8:   "uint8",
		TypeUint16:  "uint16",
		TypeUint32:  "uint32",
		TypeUint64:  "uint64",
		TypeFloat32: "float32",
		TypeFloat64: "float64",
	}

	TypeNameMap = map[string]Type{
		"TypeInvalid": TypeInvalid,
		"TypeBool":    TypeBool,
		"TypeTime":    TypeTime,
		"TypeJSON":    TypeJSON,
		"TypeUUID":    TypeUUID,
		"TypeBytes":   TypeBytes,
		"TypeEnum":    TypeEnum,
		"TypeString":  TypeString,
		"TypeOther":   TypeOther,
		"TypeInt":     TypeInt,
		"TypeInt8":    TypeInt8,
		"TypeInt16":   TypeInt16,
		"TypeInt32":   TypeInt32,
		"TypeInt64":   TypeInt64,
		"TypeUint":    TypeUint,
		"TypeUint8":   TypeUint8,
		"TypeUint16":  TypeUint16,
		"TypeUint32":  TypeUint32,
		"TypeUint64":  TypeUint64,
		"TypeFloat32": TypeFloat32,
		"TypeFloat64": TypeFloat64,
	}
)

type Edge struct {
	Desc    string
	Link    string
	From    string
	To      string
	Ref     string
	Display []*Field
}

type Table struct {
	Name   string   // 表名称
	Fields []*Field // 表字段集合
	Desc   string   // 备注
	Edges  []*Edge  // 关系
}

type Field struct {
	Tag      string      // 生成go结构体，附加的tag内容
	Name     string      // 字段名称
	Size     int         // 大小
	TypeInfo Type        // 字段类型
	Unique   bool        // 是否唯一
	Nillable bool        // 是否为NULL
	Default  interface{} // 默认值
	Comment  string      // 备注
}
