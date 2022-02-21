package regexpfoo

import (
	"fmt"
	"regexp"
)

var (
	// GO_TYPE_IDENTIFIER_EXPRESSION string = `(?P<NAME>[[:alpha:]][_\w]*)\s+(?P<TYPE>[\S]*)`
	GO_STRUCT_MEMBER_IDENTIFIER_EXPRESSION string = `(?P<NAME>[[:alpha:]][_\w]*)\s+(?P<TYPE>(\*(?P<FROM>[[:alpha:]][_\w]+)\.)?(?P<META>[[:alpha:]][_\w]+))\s+`
	GO_STRUCT_MEMBER_IDENTIFIER_CONTENT    string = "RankType *cargo_def.int32 `protobuf:\"varint,1,opt,name=rankType,proto3\" json:\"rankType,omitempty\"`"
	GO_STRUCT_MEMBER_SUBMATCH_NAME         string = "NAME"
	GO_STRUCT_MEMBER_SUBMATCH_TYPE         string = "TYPE"
	GO_STRUCT_MEMBER_SUBMATCH_TYPE_FROM    string = "FROM"
	GO_STRUCT_MEMBER_SUBMATCH_TYPE_META    string = "META"
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

	for _, submatchSlice := range testRegexp.FindAllStringSubmatch(content, matchCount) {
		for _, submatchName := range submatchNames {
			fmt.Printf("find sub match %v match content |%v|\n", submatchName, submatchSlice[submatchIndexMap[submatchName]])
		}
		fmt.Println()
	}
}
