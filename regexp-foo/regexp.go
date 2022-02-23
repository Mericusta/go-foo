package regexpfoo

import (
	"fmt"
	"regexp"
)

var (
	GO_STRUCT_MEMBER_IDENTIFIER_EXPRESSION string = `(?P<NAME>[[:alpha:]][_\w]+)\s+(?P<TYPE>(\*(?P<FROM>[[:alpha:]][_\w]+)\.)?(?P<META>[[:alpha:]][_\w]+))\s+`
	GO_STRUCT_MEMBER_IDENTIFIER_CONTENT    string = "RankType *cargo_def.int32 `protobuf:\"varint,1,opt,name=rankType,proto3\" json:\"rankType,omitempty\"`"
	GO_STRUCT_MEMBER_SUBMATCH_NAME         string = "NAME"
	GO_STRUCT_MEMBER_SUBMATCH_TYPE         string = "TYPE"
	GO_STRUCT_MEMBER_SUBMATCH_TYPE_FROM    string = "FROM"
	GO_STRUCT_MEMBER_SUBMATCH_TYPE_META    string = "META"

	GO_IMPORT_SCOPE_EXPRESSION string = `import\s*\(`
	GO_IMPORT_SCOPE_CONTENT    string = `
import (
	cargo_def "robot-prototype/protocol/tbp_protobuf_gen/cargo_def"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)`

	GO_EACH_IMPORT_EXPRESSION string = `((?P<ALIAS>[[:alpha:]][_\w]+)\s+)?"(?P<PATH>[/_\.\w-]+)"`
	GO_EACH_IMPORT_CONTENT    string = `
	cargo_def "robot-prototype/protocol/tbp_protobuf_gen/cargo_def"
	"google.golang.org/protobuf/reflect/protoreflect"
`
	GO_EACH_IMPORT_SUBMATCH_ALIAS string = "ALIAS"
	GO_EACH_IMPORT_SUBMATCH_PATH  string = "PATH"

	GO_PACKAGE_SCOPE_EXPRESSION    string = `package\s+(?P<NAME>[[:alpha:]][_\w]+)`
	GO_PACKAGE_SCOPE_CONTENT       string = `package msg_def`
	GO_PACKAGE_SCOPE_SUBMATCH_NAME string = "NAME"
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
			// panic(fmt.Sprintf("can not find sub match: %v", submatchName))
			fmt.Printf(fmt.Sprintf("can not find sub match: %v\n", submatchName))
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
