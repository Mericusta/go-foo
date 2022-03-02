package regexpfoo

import (
	"fmt"
	"regexp"
)

var (
	GO_IMPORT_SCOPE_CONTENT string = `
import (
	cargo_def "robot-prototype/protocol/tbp_protobuf_gen/cargo_def"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)`
	GO_IMPORT_SCOPE_EXPRESSION string = `import\s*\(`
	GO_EACH_IMPORT_CONTENT     string = `
	cargo_def "robot-prototype/protocol/tbp_protobuf_gen/cargo_def"
	"google.golang.org/protobuf/reflect/protoreflect"`
	GO_EACH_IMPORT_EXPRESSION     string = `((?P<ALIAS>[_\w]+)\s+)?"(?P<PATH>[/_\.\w-]+)"`
	GO_EACH_IMPORT_SUBMATCH_ALIAS string = "ALIAS"
	GO_EACH_IMPORT_SUBMATCH_PATH  string = "PATH"

	GO_PACKAGE_SCOPE_CONTENT                  string = `package msg_def`
	GO_PACKAGE_SCOPE_EXPRESSION               string = `package\s+(?P<NAME>\w+)`
	GO_PACKAGE_SCOPE_EXPRESSION_SUBMATCH_NAME string = "NAME"

	// GO_VARIABLE_DECLARATION in func declaration or struct member declaration
	// in func declaration: func([param variable declaration] [param type declaration])
	// in struct member declaration: [member variable declaration] [member type declaration]
	GO_VARIABLE_DECLARATION_CONTENT                  string = "RankType []*cargo_def.int32 `protobuf:\"varint,1,opt,name=rankType,proto3\" json:\"rankType,omitempty\"`"
	GO_VARIABLE_DECLARATION_EXPRESSION               string = `(?P<NAME>\w+)\s+(?P<TYPE>\S+)\s+`
	GO_VARIABLE_DECLARATION_EXPRESSION_SUBMATCH_NAME string = "NAME"
	GO_VARIABLE_DECLARATION_EXPRESSION_SUBMATCH_TYPE string = "TYPE"

	// GO_VARIABLE_TYPE_DECLARATION_CONTENT in func declaration or struct member declaration
	// in func declaration: func([param variable declaration] [param type declaration])
	// in struct member declaration: [member variable declaration] [member type declaration]

	GO_VARIABLE_TYPE_POINTER_DECLARATION_EXPRESSION               string = `^\*(?P<TYPE>.*)`
	GO_VARIABLE_TYPE_POINTER_DECLARATION_EXPRESSION_SUBMATCH_TYPE string = "TYPE"

	GO_VARIABLE_TYPE_INTEGER_DECLARATION_EXPRESSION  string = `^(u)?int(8|16|32|64)?`
	GO_VARIABLE_TYPE_FLOATING_DECLARATION_EXPRESSION string = `^float(32|64)`
	GO_VARIABLE_TYPE_COMPLEX_DECLARATION_EXPRESSION  string = `^complex(64|128)`
	GO_VARIABLE_TYPE_SPEC_DECLARATION_EXPRESSION     string = `^(byte|rune|uintptr)`

	GO_VARIABLE_TYPE_MAP_DECLARATION_CONTENT                     string = "map[A.Int]map[B.Int][]*C.Int"
	GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION                  string = `^map\[(?P<KEY>[^\[\]\s]+)\](?P<ELEMENT>\S+)`
	GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION_SUBMATCH_KEY     string = "KEY"
	GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION_SUBMATCH_ELEMENT string = "ELEMENT"

	GO_VARIABLE_TYPE_SLICE_DECLARATION_CONTENT                     string = "[][]map[A.Int]map[B.Int][]*C.Int"
	GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION                  string = `^\[\](?P<ELEMENT>\S+)`
	GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION_SUBMATCH_ELEMENT string = "ELEMENT"

	GO_VARIABLE_TYPE_STRUCT_DECLARATION_EXPRESSION               string = `^((?P<FROM>\w+)\.)?(?P<TYPE>\w+)`
	GO_VARIABLE_TYPE_STRUCT_DECLARATION_EXPRESSION_SUBMATCH_FROM string = "FROM"
	GO_VARIABLE_TYPE_STRUCT_DECLARATION_EXPRESSION_SUBMATCH_TYPE string = "TYPE"

	// GO_VARIABLE_SHORT_IDENTIFIER just in function body: [variable] := [type construction]
	GO_VARIABLE_SHORT_IDENTIFIER_CONTENT                  string = "v := make(map[A.Int]map[B.Int][]*C.Int)"
	GO_VARIABLE_SHORT_IDENTIFIER_EXPRESSION               string = `(?P<NAME>\w+)\s*:=\s*(?P<TYPE>\S+)`
	GO_VARIABLE_SHORT_IDENTIFIER_EXPRESSION_SUBMATCH_NAME string = "NAME"
	GO_VARIABLE_SHORT_IDENTIFIER_EXPRESSION_SUBMATCH_TYPE string = "TYPE"

	GO_VARIABLE_TYPE_CONSTRUCTION_CONTENT                  string = "x.make(map[A.Int]map[B.Int][]*C.Int)"
	GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION               string = `(?P<CALL>((?P<FROM>\w+)\.)?(?P<FUNC>\w+))\(.*\)`
	GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION_SUBMATCH_CALL string = "CALL"
	GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION_SUBMATCH_FROM string = "FROM"
	GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION_SUBMATCH_FUNC string = "FUNC"
)

func RegexpTest(matchCount int, content, expression string, submatchNames ...string) {
	testRegexp := regexp.MustCompile(expression)
	if testRegexp == nil {
		panic("can not compile expression")
	}

	submatchIndexMap := make(map[string]int, len(submatchNames))
	for _, submatchName := range submatchNames {
		submatchIndex := testRegexp.SubexpIndex(submatchName)
		if submatchIndex == -1 {
			fmt.Printf("can not find sub match: %v\n", submatchName)
			continue
		}
		submatchIndexMap[submatchName] = submatchIndex
	}

	if testRegexp.MatchString(content) {
		fmt.Println("expression can match content")
	} else {
		fmt.Println("expression can not match content")
	}

	if findStringSlice := testRegexp.FindAllString(content, -1); len(findStringSlice) > 0 {
		fmt.Println("expression can find string in content")
	} else {
		fmt.Println("expression can not find string in content")
	}

	for _, submatchSlice := range testRegexp.FindAllStringSubmatch(content, matchCount) {
		for _, submatchName := range submatchNames {
			fmt.Printf("find sub match %v match content |%v|\n", submatchName, submatchSlice[submatchIndexMap[submatchName]])
		}
	}
}

const (
	GO_META_TYPE_INTEGER = iota + 1
	GO_META_TYPE_FLOATING
	GO_META_TYPE_COMPLEX
	GO_META_TYPE_SPEC
	GO_META_TYPE_STRUCT
	GO_META_TYPE_SLICE
	GO_META_TYPE_MAP
)

type goTypeDeclaration struct {
	MetaType    int
	IsPointer   bool
	FromPkg     string
	Content     string // [][]map[A.Int]map[B.Int][]*C.Int
	KeyType     *goTypeDeclaration
	ElementType *goTypeDeclaration // [][]map[A.Int]map[B.Int][]*C.Int -> []map[A.Int]map[B.Int][]*C.Int
}

func TraitGoVariableTypeDeclaration(content string) *goTypeDeclaration {
	if len(content) == 0 {
		return nil
	}

	d := &goTypeDeclaration{
		Content: content,
	}

	fmt.Printf("content = |%v|\n", d.Content)

	pointerRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_POINTER_DECLARATION_EXPRESSION)
	sliceRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION)
	mapRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION)
	integerRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_INTEGER_DECLARATION_EXPRESSION)
	floatingRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_FLOATING_DECLARATION_EXPRESSION)
	complexRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_COMPLEX_DECLARATION_EXPRESSION)
	specRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_SPEC_DECLARATION_EXPRESSION)
	structRegexp := regexp.MustCompile(GO_VARIABLE_TYPE_STRUCT_DECLARATION_EXPRESSION)

	if pointerRegexp.MatchString(content) {
		d.IsPointer = true
		submatchSlice := pointerRegexp.FindStringSubmatch(content)
		content = submatchSlice[pointerRegexp.SubexpIndex("TYPE")]
	}

	// 为了避免在 expression 中定义识别关键字，select 必须有先后顺序：先做带有关键字的判断，最后再做非关键字判断
	switch {
	case sliceRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_SLICE
		submatchSlice := sliceRegexp.FindStringSubmatch(content)
		d.ElementType = TraitGoVariableTypeDeclaration(submatchSlice[sliceRegexp.SubexpIndex("ELEMENT")])
	case mapRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_MAP
		submatchSlice := mapRegexp.FindStringSubmatch(content)
		d.KeyType = TraitGoVariableTypeDeclaration(submatchSlice[mapRegexp.SubexpIndex("KEY")])
		d.ElementType = TraitGoVariableTypeDeclaration(submatchSlice[mapRegexp.SubexpIndex("ELEMENT")])
	case integerRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_INTEGER
	case floatingRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_FLOATING
	case complexRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_COMPLEX
	case specRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_SPEC
	case structRegexp.MatchString(content):
		d.MetaType = GO_META_TYPE_STRUCT
		submatchSlice := structRegexp.FindStringSubmatch(content)
		d.FromPkg = submatchSlice[structRegexp.SubexpIndex("FROM")]
		d.ElementType = &goTypeDeclaration{
			Content:  submatchSlice[structRegexp.SubexpIndex("TYPE")],
			MetaType: GO_META_TYPE_STRUCT,
		}
	}

	fmt.Printf("meta type = |%v|\n", d.MetaType)
	fmt.Printf("is pointer = %v\n", d.IsPointer)
	fmt.Printf("from pkg = |%v|\n", d.FromPkg)
	fmt.Printf("key type = |%+v|\n", d.KeyType)
	fmt.Printf("element type = |%+v|\n", d.ElementType)
	return d
}
