package builtinfoo

import (
	"fmt"
	substruct "go-foo/src/struct-foo/sub-struct"
	"reflect"
	"unsafe"
)

func TraversalSliceByUsingUnsafePointer(data []byte) {
	bakData := make([]byte, len(data))
	for index := 0; index != len(data); index++ {
		ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + uintptr(index)*unsafe.Sizeof(data[0]))
		bakData[index] = *(*byte)(unsafe.Pointer(ptr))
	}
}

func TraversalSliceByForRange(data []byte) {
	bakData := make([]byte, len(data))
	for i, b := range data {
		bakData[i] = b
	}
}

func TraversalSliceByFor(data []byte) {
	bakData := make([]byte, len(data))
	for i := 0; i != len(data); i++ {
		bakData[i] = data[i]
	}
}

type tmpStructWithSameMemoryLayout struct {
	ph1 int
	ph2 int
}

func ModifyPrivateValue() {
	s := substruct.SubStruct{PubV: 1}
	s.Assign(1)

	sameMemoryLayoutPointer := (*tmpStructWithSameMemoryLayout)(unsafe.Pointer(&s))
	// fmt.Printf("sameMemoryLayoutPointer.ph1 %v\n", sameMemoryLayoutPointer.ph1)
	// fmt.Printf("sameMemoryLayoutPointer.ph2 %v\n", sameMemoryLayoutPointer.ph2)
	sameMemoryLayoutPointer.ph1 = 10
	sameMemoryLayoutPointer.ph2 = 10
	// fmt.Printf("s.Val() %v\n", s.Val())
	// fmt.Printf("s.GetPubV() %v\n", s.GetPubV())
}

func BytesToStringFool(b []byte) string {
	return string(b)
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string must be ASCII code

func StringToBytesFool(s string) []byte {
	return []byte(s)
}

func StringToBytes(s string) []byte {
	if len(s) < (1 << 5) {
		return standardStringToBytes(s)
	}
	return unsafeStringToBytes(s)
}

func standardStringToBytes(s string) []byte {
	return []byte(s)
}

func unsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func assignStructMemberByFieldOffset() {
	type T struct {
		I int     `json:"int"`
		F float64 `json:"float64"`
		S string  `json:"string"`
	}
	tagKey := "json"
	sf := []string{"int", "float64", "string"}
	sa := [][]any{
		{1024, 0.618, "this is gold ratio"},
		{2048, 3.1415, "this is pi"},
		{4096, 9.8, "this is gravity"},
	}

	rs := reflect.TypeOf(T{})
	fieldCount := rs.NumField()

	koMap := make(map[string]uintptr)
	for i := 0; i < fieldCount; i++ {
		jsonKey, has := rs.Field(i).Tag.Lookup(tagKey)
		if !has {
			continue
		}
		koMap[jsonKey] = rs.Field(i).Offset
	}

	ss := make([]*T, 0, len(sa))
	for _, d := range sa {
		s := &T{}
		for i, f := range sf {
			offset, has := koMap[f]
			if !has {
				panic(f)
			}
			l, r := (unsafe.Pointer(uintptr(unsafe.Pointer(s)) + offset)), d[i]
			switchTypeKind(rs.Field(i).Type.Kind())(l, r)
		}
		ss = append(ss, s)
	}

	for _, s := range ss {
		fmt.Printf("%+v\n", s)
	}
}

func switchTypeKind(k reflect.Kind) func(unsafe.Pointer, any) {
	switch k {
	case reflect.Bool:
		return assign[bool]
	case reflect.Int:
		return assign[int]
	case reflect.Int8:
		return assign[int8]
	case reflect.Int16:
		return assign[int16]
	case reflect.Int32:
		return assign[int32]
	case reflect.Int64:
		return assign[int64]
	case reflect.Uint:
		return assign[uint]
	case reflect.Uint8:
		return assign[uint8]
	case reflect.Uint16:
		return assign[uint16]
	case reflect.Uint32:
		return assign[uint32]
	case reflect.Uint64:
		return assign[uint64]
	case reflect.Float32:
		return assign[float32]
	case reflect.Float64:
		return assign[float64]
	case reflect.Complex64:
		return assign[complex64]
	case reflect.Complex128:
		return assign[complex128]
	case reflect.String:
		return assign[string]
	}
	return nil
}

func assign[T any](l unsafe.Pointer, r any) {
	*(*T)(l) = r.(T)
}
