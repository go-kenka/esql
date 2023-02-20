package gen

import (
	"github.com/go-kenka/esql/dsl"
	"github.com/gobeam/stringy"
	"strings"
)

func CamelCase(str string) string {
	return stringy.New(str).CamelCase()
}
func Lower(str string) string {
	return strings.ToLower(str)
}

func GoType(t dsl.Type) string {
	return TypeNames[t]
}

func Add(a, b int) int {
	return a + b
}
