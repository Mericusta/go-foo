package extractorfoo

import "regexp"

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
	GO_VARIABLE_TYPE_SPEC_DECLARATION_EXPRESSION             string = `^(byte|rune|uintptr)`
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

	// fmt.Printf("content = |%v|\n", d.Content)

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
		d.FromPkg = submatchSlice[GoVariableTypeStructDeclarationRegexpSubmatchFromIndex]
		// d.ElementType = &GoTypeDeclaration{
		// 	Content:  submatchSlice[GoVariableTypeStructDeclarationRegexpSubmatchTypeIndex],
		// 	MetaType: GO_META_TYPE_STRUCT,
		// }
	}

	// fmt.Printf("meta type = |%v|\n", d.MetaType)
	// fmt.Printf("is pointer = %v\n", d.IsPointer)
	// fmt.Printf("from pkg = |%v|\n", d.FromPkg)
	// fmt.Printf("key type = |%+v|\n", d.KeyType)
	// fmt.Printf("element type = |%+v|\n", d.ElementType)
	return d
}

func ExtractGoTypeDeclarationImportPkg(d *GoTypeDeclaration) map[string]map[string]struct{} {
	importMap := make(map[string]map[string]struct{})
	if len(d.FromPkg) != 0 {
		if _, has := importMap[d.FromPkg]; !has {
			importMap[d.FromPkg] = make(map[string]struct{})
		}
		importMap[d.FromPkg][d.Content] = struct{}{}
	}
	if d.KeyType != nil {
		if keyTypeImportPkgMap := ExtractGoTypeDeclarationImportPkg(d.KeyType); len(keyTypeImportPkgMap) > 0 {
			for keyTypeImportPkg, keyTypeImportStructMap := range keyTypeImportPkgMap {
				if _, has := importMap[keyTypeImportPkg]; !has {
					importMap[keyTypeImportPkg] = make(map[string]struct{})
				}
				for keyTypeImportStruct := range keyTypeImportStructMap {
					importMap[keyTypeImportPkg][keyTypeImportStruct] = struct{}{}
				}
			}
		}
	}
	if d.ElementType != nil {
		if elementTypeImportPkgMap := ExtractGoTypeDeclarationImportPkg(d.ElementType); len(elementTypeImportPkgMap) > 0 {
			for elementTypeImportPkg, elementTypeImportStructMap := range elementTypeImportPkgMap {
				if _, has := importMap[elementTypeImportPkg]; !has {
					importMap[elementTypeImportPkg] = make(map[string]struct{})
				}
				for elementTypeImportStruct := range elementTypeImportStructMap {
					importMap[elementTypeImportPkg][elementTypeImportStruct] = struct{}{}
				}
			}
		}
	}
	return importMap
}
