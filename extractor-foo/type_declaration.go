package extractorfoo

import (
	"fmt"
	"regexp"
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
	Content      string
	MetaType     int
	FromPkgAlias string
	FromPkgPath  string
	KeyType      *GoTypeDeclaration
	ElementType  *GoTypeDeclaration
}

func (d *GoTypeDeclaration) Traversal(deep int) {
	fmt.Printf("%v- Content: %v\n", strings.Repeat("\t", deep), d.Content)
	fmt.Printf("%v- MetaType: %v\n", strings.Repeat("\t", deep), d.MetaType)
	fmt.Printf("%v- FromPkgAlias: %v\n", strings.Repeat("\t", deep), d.FromPkgAlias)
	fmt.Printf("%v- FromPkgPath: %v\n", strings.Repeat("\t", deep), d.FromPkgPath)
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

func (d *GoTypeDeclaration) TraversalFunc(f func(v *GoTypeDeclaration) bool) {
	if !f(d) {
		return
	}
	if d.KeyType != nil {
		d.KeyType.TraversalFunc(f)
	}
	if d.ElementType != nil {
		d.ElementType.TraversalFunc(f)
	}
}

func (d *GoTypeDeclaration) MakeUp() string {
	switch d.MetaType {
	case GO_META_TYPE_POINTER:
		return fmt.Sprintf("*%v", d.ElementType.MakeUp())
	case GO_META_TYPE_INTEGER, GO_META_TYPE_FLOATING, GO_META_TYPE_COMPLEX, GO_META_TYPE_SPEC:
		return d.Content
	case GO_META_TYPE_STRUCT:
		if len(d.FromPkgAlias) == 0 {
			return d.Content
		} else {
			return fmt.Sprintf("%v.%v", d.FromPkgAlias, d.Content)
		}
	case GO_META_TYPE_SLICE:
		return fmt.Sprintf("[]%v", d.ElementType.MakeUp())
	case GO_META_TYPE_MAP:
		return fmt.Sprintf("map[%v]%v", d.KeyType.MakeUp(), d.ElementType.MakeUp())
	default:
		panic("unknown meta type")
	}
}

var (
	// GO_VARIABLE_DECLARATION in func declaration or struct member declaration
	// in func declaration: func([param variable declaration] [param type declaration])
	// in struct member declaration: [member variable declaration] [member type declaration]
	GO_VARIABLE_TYPE_POINTER_DECLARATION_EXPRESSION          string = `^\*(?P<TYPE>.*)`
	GoVariableTypePointerDeclarationRegexp                          = regexp.MustCompile(GO_VARIABLE_TYPE_POINTER_DECLARATION_EXPRESSION)
	GoVariableTypePointerDeclarationRegexpSubmatchTypeIndex         = GoVariableTypePointerDeclarationRegexp.SubexpIndex("TYPE")
	GO_VARIABLE_DECLARATION_EXPRESSION                       string = `(?P<NAME>\w+)\s+(?P<TYPE>\S+)\s+`
	GoVariableDeclarationRegexp                                     = regexp.MustCompile(GO_VARIABLE_DECLARATION_EXPRESSION)
	GoVariableDeclarationRegexpSubmatchNameIndex                    = GoVariableDeclarationRegexp.SubexpIndex("NAME")
	GoVariableDeclarationRegexpSubmatchTypeIndex                    = GoVariableDeclarationRegexp.SubexpIndex("TYPE")
	GO_VARIABLE_TYPE_INTEGER_DECLARATION_EXPRESSION          string = `^(u)?int(8|16|32|64)?`
	GoVariableTypeIntegerDeclarationRegexp                          = regexp.MustCompile(GO_VARIABLE_TYPE_INTEGER_DECLARATION_EXPRESSION)
	GO_VARIABLE_TYPE_FLOATING_DECLARATION_EXPRESSION         string = `^float(32|64)`
	GoVariableTypeFloatingDeclarationRegexp                         = regexp.MustCompile(GO_VARIABLE_TYPE_FLOATING_DECLARATION_EXPRESSION)
	GO_VARIABLE_TYPE_COMPLEX_DECLARATION_EXPRESSION          string = `^complex(64|128)`
	GoVariableTypeComplexDeclarationRegexp                          = regexp.MustCompile(GO_VARIABLE_TYPE_COMPLEX_DECLARATION_EXPRESSION)
	GO_VARIABLE_TYPE_SPEC_DECLARATION_EXPRESSION             string = `^(string|byte|rune|uintptr|bool)`
	GoVariableTypeSpecDeclarationRegexp                             = regexp.MustCompile(GO_VARIABLE_TYPE_SPEC_DECLARATION_EXPRESSION)
	GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION            string = `^\[\](?P<ELEMENT>\S+)`
	GoVariableTypeSliceDeclarationRegexp                            = regexp.MustCompile(GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION)
	GoVariableTypeSliceDeclarationRegexpSubmatchElementIndex        = GoVariableTypeSliceDeclarationRegexp.SubexpIndex("ELEMENT")
	GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION              string = `^map\[(?P<KEY>[^\[\]\s]+)\](?P<ELEMENT>\S+)`
	GoVariableTypeMapDeclarationRegexp                              = regexp.MustCompile(GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION)
	GoVariableTypeMapDeclarationRegexpSubmatchKeyIndex              = GoVariableTypeMapDeclarationRegexp.SubexpIndex("KEY")
	GoVariableTypeMapDeclarationRegexpSubmatchElementIndex          = GoVariableTypeMapDeclarationRegexp.SubexpIndex("ELEMENT")
	GO_VARIABLE_TYPE_STRUCT_DECLARATION_EXPRESSION           string = `^((?P<FROM>\w+)\.)?(?P<TYPE>\w+)`
	GoVariableTypeStructDeclarationRegexp                           = regexp.MustCompile(GO_VARIABLE_TYPE_STRUCT_DECLARATION_EXPRESSION)
	GoVariableTypeStructDeclarationRegexpSubmatchFromIndex          = GoVariableTypeStructDeclarationRegexp.SubexpIndex("FROM")
	GoVariableTypeStructDeclarationRegexpSubmatchTypeIndex          = GoVariableTypeStructDeclarationRegexp.SubexpIndex("TYPE")
)

func ExtractGoVariableTypeDeclaration(content string) *GoTypeDeclaration {
	if len(content) == 0 {
		return nil
	}

	d := &GoTypeDeclaration{
		Content: content,
	}

	// 为了避免在 expression 中定义识别关键字，select 必须有先后顺序：先做带有关键字的判断，最后再做非关键字判断
	switch {
	case GoVariableTypePointerDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_POINTER
		submatchSlice := GoVariableTypePointerDeclarationRegexp.FindStringSubmatch(content)
		d.ElementType = ExtractGoVariableTypeDeclaration(submatchSlice[GoVariableTypePointerDeclarationRegexpSubmatchTypeIndex])
	case GoVariableTypeSliceDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_SLICE
		submatchSlice := GoVariableTypeSliceDeclarationRegexp.FindStringSubmatch(content)
		d.ElementType = ExtractGoVariableTypeDeclaration(submatchSlice[GoVariableTypeSliceDeclarationRegexpSubmatchElementIndex])
	case GoVariableTypeMapDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_MAP
		submatchSlice := GoVariableTypeMapDeclarationRegexp.FindStringSubmatch(content)
		d.KeyType = ExtractGoVariableTypeDeclaration(submatchSlice[GoVariableTypeMapDeclarationRegexpSubmatchKeyIndex])
		d.ElementType = ExtractGoVariableTypeDeclaration(submatchSlice[GoVariableTypeMapDeclarationRegexpSubmatchElementIndex])
	case GoVariableTypeIntegerDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_INTEGER
	case GoVariableTypeFloatingDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_FLOATING
	case GoVariableTypeComplexDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_COMPLEX
	case GoVariableTypeSpecDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_SPEC
	case GoVariableTypeStructDeclarationRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_STRUCT
		submatchSlice := GoVariableTypeStructDeclarationRegexp.FindStringSubmatch(content)
		d.FromPkgAlias = submatchSlice[GoVariableTypeStructDeclarationRegexpSubmatchFromIndex]
		d.Content = submatchSlice[GoVariableTypeStructDeclarationRegexpSubmatchTypeIndex]
	}
	return d
}

func ExtractGoTypeDeclarationImportPkg(d *GoTypeDeclaration) map[string]map[string]struct{} {
	importMap := make(map[string]map[string]struct{})
	if len(d.FromPkgAlias) != 0 {
		if _, has := importMap[d.FromPkgAlias]; !has {
			importMap[d.FromPkgAlias] = make(map[string]struct{})
		}
		importMap[d.FromPkgAlias][d.Content] = struct{}{}
	}
	if d.KeyType != nil {
		if keyTypeImportPkgMap := ExtractGoTypeDeclarationImportPkg(d.KeyType); len(keyTypeImportPkgMap) > 0 {
			for keyTypeImportPkgAlias, keyTypeImportStructMap := range keyTypeImportPkgMap {
				if _, has := importMap[keyTypeImportPkgAlias]; !has {
					importMap[keyTypeImportPkgAlias] = make(map[string]struct{})
				}
				for keyTypeImportStruct := range keyTypeImportStructMap {
					importMap[keyTypeImportPkgAlias][keyTypeImportStruct] = struct{}{}
				}
			}
		}
	}
	if d.ElementType != nil {
		if elementTypeImportPkgMap := ExtractGoTypeDeclarationImportPkg(d.ElementType); len(elementTypeImportPkgMap) > 0 {
			for elementTypeImportPkgAlias, elementTypeImportStructMap := range elementTypeImportPkgMap {
				if _, has := importMap[elementTypeImportPkgAlias]; !has {
					importMap[elementTypeImportPkgAlias] = make(map[string]struct{})
				}
				for elementTypeImportStruct := range elementTypeImportStructMap {
					importMap[elementTypeImportPkgAlias][elementTypeImportStruct] = struct{}{}
				}
			}
		}
	}
	return importMap
}
