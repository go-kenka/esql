package gen

import (
	"github.com/gobeam/stringy"
	"strings"
)

func CamelCase(str string) string {
	return stringy.New(str).CamelCase()
}
func Lower(str string) string {
	return strings.ToLower(str)
}

func GoType(t Type) string {
	return TypeNames[t]
}

func Add(a, b int) int {
	return a + b
}
