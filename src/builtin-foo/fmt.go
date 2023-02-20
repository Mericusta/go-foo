package builtinfoo

import (
	"fmt"
	"reflect"
	"unsafe"
)

type fmtStruct struct {
	i int
	s string
}

func FmtPtr() {
	var i int = 1
	var iPtr *int = &i
	var str string = "abc"
	var strPtr *string = &str
	var m map[int]int = map[int]int{1: 1} // as same as make(map[int]int)
	var mPtr *map[int]int = &m
	var slice []int = []int{1} // as same as make([]int, 0, 2)
	var slicePtr *[]int = &slice
	var s fmtStruct = fmtStruct{i: 1, s: "abc"}
	var sPtr *fmtStruct = &s

	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Printf("v %v\n", i)
	fmt.Printf("p -\n")
	fmt.Println(iPtr)
	fmt.Println(reflect.TypeOf(iPtr))
	fmt.Printf("v %v\n", iPtr)
	fmt.Printf("p %p\n", iPtr)
	fmt.Printf("&i == iPtr %v\n", &i == iPtr)
	fmt.Println()

	fmt.Printf("for integer type, including intN, floatN\n")
	fmt.Printf("- %%v is its origin value\n")
	fmt.Printf("- can not output %%p because it is not a pointer\n")
	fmt.Printf("for integer type's pointer\n")
	fmt.Printf("- %%v is equal to %%p\n")
	fmt.Printf("- %%p is pointer value\n")
	fmt.Println()

	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
	fmt.Printf("v %v\n", s)
	fmt.Printf("p -\n")
	fmt.Println(sPtr)
	fmt.Println(reflect.TypeOf(sPtr))
	fmt.Printf("v %v\n", sPtr)
	fmt.Printf("p %p\n", sPtr)
	fmt.Printf("&s == sPtr %v\n", &s == sPtr)
	fmt.Printf("sPtr == &s.i %v\n", uintptr(unsafe.Pointer(sPtr)) == uintptr(unsafe.Pointer(&s.i)))
	fmt.Println()

	fmt.Printf("for struct type\n")
	fmt.Printf("- %%v is struct content value\n")
	fmt.Printf("- can not output %%p because it is not a pointer\n")
	fmt.Printf("for struct type's pointer\n")
	fmt.Printf("- %%v is pointed value, same like output %%v struct\n")
	fmt.Printf("- %%p is pointer value\n")
	fmt.Printf("- %%p is equal to struct first element's address\n")
	fmt.Println()

	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))
	fmt.Printf("v %v\n", m)
	fmt.Printf("p %p\n", m)
	fmt.Println(mPtr)
	fmt.Println(reflect.TypeOf(mPtr))
	fmt.Printf("v %v\n", mPtr)
	fmt.Printf("p %p\n", mPtr)
	fmt.Printf("*mPtr == &m %v\n", reflect.DeepEqual(*mPtr, m))
	fmt.Println()

	fmt.Printf("for map type\n")
	fmt.Printf("- the underlying type is *hmap, a pointer to a struct named hmap\n")
	fmt.Printf("- make func defined in runtime/map.go\n")
	fmt.Printf("- %%v is map content value\n")
	fmt.Printf("- %%p is pointer value\n")
	fmt.Printf("for map type's pointer\n")
	fmt.Printf("- the underlying type is **hmap\n")
	fmt.Printf("- %%v is *hmap\n")
	fmt.Printf("- %%p is pointer value\n")
	fmt.Println()

	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Printf("v %v\n", slice)
	fmt.Printf("p %p\n", slice)
	fmt.Printf("first element p %p\n", &slice[0])
	fmt.Println(slicePtr)
	fmt.Println(reflect.TypeOf(slicePtr))
	fmt.Printf("v %v\n", slicePtr)
	fmt.Printf("p %p\n", slicePtr)
	fmt.Printf("reflect.SliceHeader.Data == &slice[0] %v\n", (*reflect.SliceHeader)(unsafe.Pointer(slicePtr)).Data == uintptr(unsafe.Pointer(&slice[0])))
	fmt.Println()

	fmt.Printf("for slice type\n")
	fmt.Printf("- the underlying type is slice, a struct named slice defined in runtime/slice.go\n")
	fmt.Printf("- struct looks like reflect.SliceHeader but using unsafe.Pointer to hold content memory address\n")
	fmt.Printf("- %%v is struct content value\n")
	fmt.Printf("- %%p is different from struct, befause fmt.Printf output its first element address\n")
	fmt.Printf("for slice type's pointer\n")
	fmt.Printf("- %%v is pointed value, same like output %%v struct\n")
	fmt.Printf("- %%p is pointer value\n")
	fmt.Println()

	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))
	fmt.Printf("v %v\n", str)
	fmt.Printf("p -\n")
	fmt.Println(strPtr)
	fmt.Println(reflect.TypeOf(strPtr))
	fmt.Printf("v %v\n", strPtr)
	fmt.Printf("p %p\n", strPtr)
	fmt.Printf("string -> []byte %v\n", []byte(str))
	fmt.Printf("reflect.Header %v\n", (*reflect.StringHeader)(unsafe.Pointer(strPtr)))
	stringContentPtr := (*reflect.StringHeader)(unsafe.Pointer(strPtr)).Data
	for i := uintptr(0); int(i) < len([]byte(str)); i++ {
		fmt.Printf("No.%v byte %v\n", i, *(*byte)(unsafe.Pointer(stringContentPtr + i)))
	}
	fmt.Println()

	fmt.Printf("for string type\n")
	fmt.Printf("- the underlying type is stringStruct, a struct named stringStruct defined in runtime/string.go\n")
	fmt.Printf("- struct looks like reflect.StringHeader but using unsafe.Pointer to hold content memory address\n")
	fmt.Printf("- struct has a pointer, pointing to []byte, an array, element type is byte\n")
	fmt.Printf("- %%v is array's content value\n")
	fmt.Printf("- can not output %%p because it is not a pointer\n")
	fmt.Printf("for string type's pointer\n")
	fmt.Printf("- %%v is equal to %%p\n")
	fmt.Printf("- %%p is pointer value\n")
	fmt.Printf("for reflect.StringHeader\n")
	fmt.Printf("- reflect.StringHeader.Data is string content memory address, but using unsafe.Pointer\n")
	fmt.Printf("- string content memory allocation is as same as array, so it can accessed by iteration content pointer\n")
	fmt.Println()
}
