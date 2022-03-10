package workerfoo

import (
	"fmt"
	extractorfoo "go-foo/extractor-foo"
	"strings"
)

var (
	MAKE_SLICE_TEMPLATE                = `make([][RP_SLICE_ELEMENT_TYPE], 0[RP_SLICE_CAPACITY])`
	REPLACE_KEYWORD_SLICE_ELEMENT_TYPE = "[RP_SLICE_ELEMENT_TYPE]"
	REPLACE_KEYWORD_SLICE_CAPACITY     = "[RP_SLICE_CAPACITY]"
)

func MakeSlice(elementD *extractorfoo.GoTypeDeclaration, capacity int) string {
	s := strings.Replace(MAKE_SLICE_TEMPLATE, REPLACE_KEYWORD_SLICE_ELEMENT_TYPE, elementD.MakeUp(), -1)
	s = strings.Replace(s, REPLACE_KEYWORD_SLICE_CAPACITY, fmt.Sprintf(", %v", capacity), -1)
	return s
}

func MakeSliceTest() {
	d := extractorfoo.ExtractGoVariableTypeDeclaration("[][]map[Float]map[A.Int][]*B.Int")
	d.Traversal(0)
	d.TraversalFunc(func(v *extractorfoo.GoTypeDeclaration) bool {
		if v.MetaType == extractorfoo.GO_META_TYPE_SLICE {
			fmt.Printf("%v\n", MakeSlice(v.ElementType, 0))
		}
		return true
	})
}
