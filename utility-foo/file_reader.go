package utilityfoo

import (
	"fmt"
	"go-foo/utility"
	"regexp"
)

func ReadMarkdownTopic() {
	markdownDepthSpaceWidth := 4
	rootTopicRegexp := regexp.MustCompile(fmt.Sprintf(`^\s*-\s+%v\s*$`, "Artillery wagon"))
	topicRegexp := regexp.MustCompile(`^(?P<DEPTH>\s+)-\s+(?P<TOPIC>.*)$`)
	depthIndex := topicRegexp.SubexpIndex("DEPTH")
	topicIndex := topicRegexp.SubexpIndex("TOPIC")
	if depthIndex == -1 || topicIndex == -1 {
		panic("submatch DEPTH or TOPIC index is -1")
	}
	inTopicScope := false

	utility.ReadFileLineOneByOne("./resources/factorio.md", func(s string) bool {
		switch {
		case rootTopicRegexp.MatchString(s):
			inTopicScope = true
			fmt.Printf("root topic = |%v|", s)
			return true
		case inTopicScope:
			if !topicRegexp.MatchString(s) {
				return false
			}
			stringSubmatchSlice := topicRegexp.FindStringSubmatch(s)
			if depthIndex >= len(stringSubmatchSlice) {
				panic(fmt.Sprintf("not find sub match DEPTH at |%v|", s))
			}
			depth := len(stringSubmatchSlice[depthIndex]) / markdownDepthSpaceWidth
			if topicIndex >= len(stringSubmatchSlice) {
				panic(fmt.Sprintf("not find sub match TOPIC at |%v|", s))
			}
			fmt.Printf("depth = %v, topic = |%v|", depth, stringSubmatchSlice[topicIndex])
			return true
		default:
			return true
		}
	})
}
