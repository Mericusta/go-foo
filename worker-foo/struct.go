package workerfoo

import (
	extractorfoo "go-foo/extractor-foo"
	"strings"
)

var (
	CONSTRUCT_STRUCT_TEMPLATE          = `[RP_POINTER][RP_STRUCT_NAME]{}`
	REPLACE_KEYWORD_STRUCT_KEY_POINTER = "[RP_POINTER]"
	REPLACE_KEYWORD_STRUCT_KEY_TYPE    = "[RP_STRUCT_NAME]"
)

func ConstructStruct(name string, isPointer bool, memberDeclarationMap map[string]*extractorfoo.GoStructMemberDefinition) string {
	pointerSignature := ""
	if isPointer {
		pointerSignature = "&"
	}
	s := strings.Replace(CONSTRUCT_STRUCT_TEMPLATE, REPLACE_KEYWORD_STRUCT_KEY_POINTER, pointerSignature, -1)
	s = strings.Replace(s, REPLACE_KEYWORD_STRUCT_KEY_TYPE, name, -1)
	return s
}

func ConstructStructTest() {
	// d := extractorfoo.
}
