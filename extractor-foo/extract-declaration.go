package extractorfoo

import (
	"fmt"
	"strings"
)

const (
	GO_META_TYPE_POINTER = iota + 1
	GO_META_TYPE_INTEGER
	GO_META_TYPE_FLOATING
	GO_META_TYPE_COMPLEX
	GO_META_TYPE_SPEC
	GO_META_TYPE_STRUCT
	GO_META_TYPE_SLICE
	GO_META_TYPE_MAP
)

// package ex
// [][]map[Float]map[A.Int][]*B.Int
// [] + []map[Float]map[A.Int][]*B.Int
// [] + [] + map[Float]map[A.Int][]*B.Int
// [] + [] + map + ex.Float + map[A.Int][]*B.Int
// [] + [] + map + ex.Float + map + A.Int + []*B.Int
// [] + [] + map + ex.Float + map + A.Int + [] + * + B.Int

type GoTypeDeclaration struct {
	Content     string
	MetaType    int
	FromPkg     string
	KeyType     *GoTypeDeclaration
	ElementType *GoTypeDeclaration
}

func (d *GoTypeDeclaration) Traversal(deep int) {
	fmt.Printf("%v- Content: %v\n", strings.Repeat("\t", deep), d.Content)
	fmt.Printf("%v- MetaType: %v\n", strings.Repeat("\t", deep), d.MetaType)
	fmt.Printf("%v- FromPkg: %v\n", strings.Repeat("\t", deep), d.FromPkg)
	if d.KeyType != nil {
		fmt.Printf("%v- KeyType:\n", strings.Repeat("\t", deep))
		d.KeyType.Traversal(deep + 1)
	}
	if d.ElementType != nil {
		fmt.Printf("%v- ElementType:\n", strings.Repeat("\t", deep))
		d.ElementType.Traversal(deep + 1)
	}
	fmt.Printf("%v- MakeUp: %v\n", strings.Repeat("\t", deep), d.MakeUp())
}

func (d *GoTypeDeclaration) MakeUp() string {
	switch d.MetaType {
	case GO_META_TYPE_POINTER:
		return fmt.Sprintf("*%v", d.ElementType.MakeUp())
	case GO_META_TYPE_INTEGER, GO_META_TYPE_FLOATING, GO_META_TYPE_COMPLEX, GO_META_TYPE_SPEC:
		return d.Content
	case GO_META_TYPE_STRUCT:
		if len(d.FromPkg) == 0 {
			return d.Content
		} else {
			return fmt.Sprintf("%v", d.Content)
		}
	case GO_META_TYPE_SLICE:
		return fmt.Sprintf("[]%v", d.ElementType.MakeUp())
	case GO_META_TYPE_MAP:
		return fmt.Sprintf("map[%v]%v", d.KeyType.MakeUp(), d.ElementType.MakeUp())
	default:
		panic("unknown meta type")
	}
}
