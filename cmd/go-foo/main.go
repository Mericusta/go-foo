package main

import (
	gcfoo "go-foo/src/gc-foo"
	"reflect"
	"unsafe"
)

func simpleReturnFunc() ([]int, *reflect.SliceHeader) {
	v := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		v = append(v, i)
	}
	vsh := &reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&v[0])), Len: 10, Cap: 10}
	return v, vsh
}

func EscapeFoo() {
	_, _ = simpleReturnFunc()
}

func main() {
	gcfoo.EscapeFoo()
}
