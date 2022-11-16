package slicefoo

import (
	"fmt"
)

// slice 的定义：
// type slice struct {
// 	array unsafe.Pointer
// 	len   int
// 	cap   int
// }

// %p 对于 slice 类型的变量 s，其输出的是 s.array
// %p 对于 slice 类型的变量 s 的取地址操作 &s，其输出的是 s 的地址

// 这里的 s 是外层 s 的一个拷贝
// []int 本质是一个名为 slice 的 struct
// 所以这里是 s 这个结构体的一个拷贝
func appendLowerCapacitySliceFunc(s []int) {
	fmt.Printf("after pass, before append: %p, %v, cap(s) = %v\n", &s, s, cap(s))
	c := cap(s)
	for index := 0; index != c-1; index++ {
		s = append(s, index)
	}
	fmt.Printf("after pass, after append %p, %v, cap(s) = %v\n", &s, s, cap(s))
}

// 这里的 s 是外层 s 的一个拷贝
// 但是 s[0] 这个操作本质上是对 slice.array 这个指针的一次 +0 操作
// s[0] = 1 读取 slice.array+0 这个指针的值并修改值为 1
func updateSliceFunc(s []int) {
	fmt.Printf("after pass, before update: %p, %v, cap(s) = %v\n", &s, s, cap(s))
	s[0] = 1
	fmt.Printf("after pass, after update %p, %v, cap(s) = %v\n", &s, s, cap(s))
}

// 这里的 s 是外层 s 的一个拷贝
// []int 本质是一个名为 slice 的 struct
// 所以这里是 s 这个结构体的一个拷贝
func appendGreaterCapacitySliceFunc(s []int) {
	fmt.Printf("after pass, before append: %p, %v, cap(s) = %v\n", &s, s, cap(s))
	c := cap(s)
	for index := 0; index != c+1; index++ {
		s = append(s, index)
	}
	fmt.Printf("after pass, after append %p, %v, cap(s) = %v\n", &s, s, cap(s))
}

func PassSliceAndChangeIt() {
	s := make([]int, 0, 4)
	fmt.Printf("after init, before append: %p, %v, cap(s) = %v\n", &s, s, cap(s))
	s = append(s, 0)
	fmt.Printf("after init, after append: %p, %v, cap(s) = %v\n", &s, s, cap(s))

	fmt.Printf("s = %p, &s[0] = %p, %%p for slice value 's' is the first element pointer in slice, that is, slice.array\n", s, &s[0])
	fmt.Printf("&s = %p, s = %p, &s[0] = %p, %%p for slice value 's' address is just what you want\n", &s, s, &s[0])

	appendLowerCapacitySliceFunc(s)
	fmt.Printf("after append func: %p, %v, cap(s) = %v\n", &s, s, cap(s))
	updateSliceFunc(s)
	fmt.Printf("after update func: %p, %v, cap(s) = %v\n", &s, s, cap(s))
	appendGreaterCapacitySliceFunc(s)
	fmt.Printf("after append func: %p, %v, cap(s) = %v\n", &s, s, cap(s))
}

func ResetSliceFoo() {
	s := make([]int, 0, 8)
	for index := 0; index != 8; index++ {
		s = append(s, index)
	}
	fmt.Printf("s %v, ptr %p, len %v, cap %v\n", s, s, len(s), cap(s))

	// reset s
	s = s[:0]
	for index := 0; index != 16; index++ {
		s = append(s, index) // grow cap will change ptr
	}
	fmt.Printf("s %v, ptr %p, len %v, cap %v\n", s, s, len(s), cap(s))

	// reset s will not change ptr
	s = s[:0]
	for index := 0; index != 8; index++ {
		s = append(s, index)
	}
	fmt.Printf("s %v, ptr %p, len %v, cap %v\n", s, s, len(s), cap(s))

	// make s will change ptr
	s = s[:0]
	for index := 0; index != 8; index++ {
		s = append(s, index)
	}
	fmt.Printf("s %v, ptr %p, len %v, cap %v\n", s, s, len(s), cap(s))
}
