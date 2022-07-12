package arrayfoo

func ClearArrayFoo() {
	a := make([]int, 10)
	// fmt.Printf("after init, len(a) = %v, cap(a) = %v, a = %v\n", len(a), cap(a), a)
	for index := 0; index != len(a); index++ {
		a[index] = index
	}
	// fmt.Printf("after assign, len(a) = %v, cap(a) = %v, a = %v\n", len(a), cap(a), a)
	a = a[:]
	// fmt.Printf("after clear, len(a) = %v, cap(a) = %v, a = %v\n", len(a), cap(a), a)
}

func returnArrayBeforeIndex(a []int, i, c int) []int {
	if i >= c {
		return nil
	}
	for index := 0; index != c; index++ {
		a[index] = index
	}
	return a[:i]
}

func ReturnArrayBeforeIndexFoo() {
	a := make([]int, 10)
	// fmt.Printf("after init, len(a) = %v, cap(a) = %v, a = %v, &a = %p, a = %p\n", len(a), cap(a), a, &a, a)
	_ = returnArrayBeforeIndex(a, 5, 10)
	// fmt.Printf("after return, len(a) = %v, cap(a) = %v, a = %v, &a = %p, a = %p\n", len(a), cap(a), a, &a, a)
	// fmt.Printf("after return, len(ra) = %v, cap(ra) = %v, ra = %v, &ra = %p, ra = %p\n", len(ra), cap(ra), ra, &ra, ra)
}
