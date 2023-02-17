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

	sUintptrSlice := make([]uintptr, 0, 8)
	sPtrSlice := make([]*tmpS, 0, 8)

	if catch {
		// catch reference
		for i := 0; i != 8; i++ {
			s := &tmpS{i: i}
			fmt.Printf("o ptr %v\n", uintptr(unsafe.Pointer(s)))
			fmt.Printf("o ptr %v\n", uintptr(unsafe.Pointer(&s)))
			sUintptrSlice = append(sUintptrSlice, uintptr(unsafe.Pointer(s)))
			// whatever it is, compiler can recognize if 's' is catched, then make up for logic
			if !catch {
				sPtrSlice = append(sPtrSlice, s)
			}
		}
	} else {
		// lost reference
		for i := 0; i != 8; i++ {
			s := &tmpS{i: i}
			fmt.Printf("o ptr %v\n", uintptr(unsafe.Pointer(s)))
			fmt.Printf("o ptr %v\n", uintptr(unsafe.Pointer(&s)))
			sUintptrSlice = append(sUintptrSlice, uintptr(unsafe.Pointer(s)))
		}
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
