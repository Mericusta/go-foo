package builtinfoo

import (
	"strings"
)

func CopyStringSliceFromStringsSplit(c int, s string, lenCase string) []string {
	var strSlice []string
	switch lenCase {
	case "==": // equal
		strSlice = make([]string, c)
	case "<":
		strSlice = make([]string, c-1)
	case ">":
		strSlice = make([]string, c+1)
	}
	copy(strSlice, strings.Split(s, ","))
	return strSlice
}
