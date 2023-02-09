package builtinfoo

import (
	"fmt"
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

func CopyMyself(b []int, from, to int) []int {
	copy(b[0:], b[from:to])
	return b
}

// if for-range from a func, the func will execute only once
func ForRangeFoo(rangeFunc func() []int) {
	if rangeFunc == nil {
		rangeFunc = func() []int {
			s := make([]int, 0, 10)
			for index := 0; index != 10; index++ {
				s = append(s, index)
			}
			fmt.Printf("return s = %v, ptr %p\n", s, s)
			return s
		}
	}
	for i, v := range rangeFunc() {
		fmt.Printf("index %v, value %v\n", i, v)
	}
}
