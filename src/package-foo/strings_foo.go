package packagefoo

import "strings"

func stringsTrimFoo(s, cutset string) string {
	return strings.Trim(s, cutset)
}
