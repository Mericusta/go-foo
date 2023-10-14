package builtinfoo

import (
	"fmt"
	"unsafe"
)

func tmpValueAssignInForRange() {
	s := make([]int, 0, 8)
	for i := 0; i != 8; i++ {
		s = append(s, i)
	}

	for i, v := range s {
		fmt.Printf("i %v, v %v, v ptr %p, s[%v] ptr %p\n", i, v, &v, i, &s[i])
	}

	fmt.Println("in for-range, tmp value v assigned from slice in each iterator")
}

func localValueReassignInFor(catch bool) {
	type tmpS struct{ i int }

	type uintptrHolder struct{ p uintptr }
	sUintptrHolderSlice := make([]*uintptrHolder, 0, 8)
	sUintptrSlice := make([]uintptr, 0, 8)
	sPtrSlice := make([]*tmpS, 0, 8)

	if catch {
		// catch reference
		for i := 0; i != 8; i++ {
			s := &tmpS{i: i}
			fmt.Printf("ptr value %v\n", uintptr(unsafe.Pointer(s)))
			fmt.Printf("ptr ptr %v\n", uintptr(unsafe.Pointer(&s)))
			sUintptrSlice = append(sUintptrSlice, uintptr(unsafe.Pointer(s)))
			// whatever it is, compiler can recognize if 's' is caught, then make up for logic
			if !catch {
				sPtrSlice = append(sPtrSlice, s)
			}
		}
	} else {
		// lost reference
		for i := 0; i != 8; i++ {
			s := &tmpS{i: i}                                         // this will change old pointer pointed object
			fmt.Printf("ptr value %v\n", uintptr(unsafe.Pointer(s))) // new object
			fmt.Printf("ptr ptr %v\n", uintptr(unsafe.Pointer(&s)))  // old pointer
			sUintptrSlice = append(sUintptrSlice, uintptr(unsafe.Pointer(s)))
			sUintptrHolderSlice = append(sUintptrHolderSlice, &uintptrHolder{
				p: uintptr(unsafe.Pointer(s)),
			})
		}
	}

	fmt.Printf("sUintptrHolderSlice = %v\n", sUintptrHolderSlice)
	for i, p := range sUintptrHolderSlice {
		s := (*tmpS)(unsafe.Pointer(p.p))
		fmt.Printf("i %v s %v\n", i, s)
	}
	fmt.Printf("sUintptrSlice = %v\n", sUintptrSlice)
	for i, p := range sUintptrSlice {
		s := (*tmpS)(unsafe.Pointer(p))
		fmt.Printf("i %v s %v\n", i, s)
	}
	fmt.Printf("sPtrSlice = %v\n", sPtrSlice)
	for i, p := range sPtrSlice {
		fmt.Printf("i %v s %v\n", i, p)
	}
}

func forCompareCall(c int) {
	compareFunc := func() int {
		fmt.Println(c)
		return c
	}

	for i := 0; i < compareFunc(); i++ {
		fmt.Printf("i = %v\n", i)
	}
}

type myMap struct {
	m map[int]int
}

func (m *myMap) myForeach(f func(k, v int) bool) {
	for key, value := range m.m {
		if !f(key, value) {
			return
		}
	}
}

func (m *myMap) passValueFoo() map[int]int {
	m.m = make(map[int]int)
	for index := 0; index != 10; index++ {
		m.m[index] = index
	}
	ret := make(map[int]int)
	m.myForeach(func(k, v int) bool {
		ret[k] = v
		return true
	})
	return ret
}
