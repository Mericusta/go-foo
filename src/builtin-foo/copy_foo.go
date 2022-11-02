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

func CopyByteSliceFromStringWithThreeCases(c int, s string, lenCase string) []byte {
	var byteSlice []byte
	switch lenCase {
	case "==": // equal
		byteSlice = make([]byte, c)
	case "<":
		byteSlice = make([]byte, c-1)
	case ">":
		byteSlice = make([]byte, c+1)
	}
	copy(byteSlice, s)
	return byteSlice
}

func CopyByteSliceFromString(s string) []byte {
	byteSlice := make([]byte, len(s))
	copy(byteSlice, s)
	return byteSlice
}
