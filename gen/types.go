package gen

import "github.com/go-kenka/esql/dsl"

const (
	TypeInvalid = dsl.TypeInvalid
	TypeBool    = dsl.TypeBool
	TypeTime    = dsl.TypeTime
	TypeJSON    = dsl.TypeJSON
	TypeUUID    = dsl.TypeUUID
	TypeBytes   = dsl.TypeBytes
	TypeEnum    = dsl.TypeEnum
	TypeString  = dsl.TypeString
	TypeOther   = dsl.TypeOther
	TypeInt8    = dsl.TypeInt8
	TypeInt16   = dsl.TypeInt16
	TypeInt32   = dsl.TypeInt32
	TypeInt     = dsl.TypeInt
	TypeInt64   = dsl.TypeInt64
	TypeUint8   = dsl.TypeUint8
	TypeUint16  = dsl.TypeUint16
	TypeUint32  = dsl.TypeUint32
	TypeUint    = dsl.TypeUint
	TypeUint64  = dsl.TypeUint64
	TypeFloat32 = dsl.TypeFloat32
	TypeFloat64 = dsl.TypeFloat64
)

var (
	TypeGoNames = [...]string{
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

	TypeNames = [...]string{
		TypeInvalid: "TypeInvalid",
		TypeBool:    "TypeBool",
		TypeTime:    "TypeTime",
		TypeJSON:    "TypeJSON",
		TypeUUID:    "TypeUUID",
		TypeBytes:   "TypeBytes",
		TypeEnum:    "TypeEnum",
		TypeString:  "TypeString",
		TypeOther:   "TypeOther",
		TypeInt:     "TypeInt",
		TypeInt8:    "TypeInt8",
		TypeInt16:   "TypeInt16",
		TypeInt32:   "TypeInt32",
		TypeInt64:   "TypeInt64",
		TypeUint:    "TypeUint",
		TypeUint8:   "TypeUint8",
		TypeUint16:  "TypeUint16",
		TypeUint32:  "TypeUint32",
		TypeUint64:  "TypeUint64",
		TypeFloat32: "TypeFloat32",
		TypeFloat64: "TypeFloat64",
	}

	TypeNameMap = map[string]dsl.Type{
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

const (
	TypeO2O = dsl.TypeO2O
	TypeO2M = dsl.TypeO2M
	TypeM2O = dsl.TypeM2O
	TypeM2M = dsl.TypeM2M
)

var EdgeTypeNameMap = map[string]dsl.EdgeType{
	"TypeO2O": TypeO2O,
	"TypeO2M": TypeO2M,
	"TypeM2O": TypeM2O,
	"TypeM2M": TypeM2M,
}

type Edge struct {
	Name     string
	Type     dsl.EdgeType
	Link     string
	From     string
	Ref      string
	Display  []*Field
	Relation []*Edge
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
	TypeInfo dsl.Type    // 字段类型
	Unique   bool        // 是否唯一
	Nillable bool        // 是否为NULL
	Default  interface{} // 默认值
	Comment  string      // 备注
}
