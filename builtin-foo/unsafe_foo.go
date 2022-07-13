package builtinfoo

import (
	substruct "go-foo/struct-foo/sub-struct"
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
