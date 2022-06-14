package mapfoo

import "fmt"

func upperCapacityMapFunc(m map[int]int, count int) {
	for index := 0; index != count*2; index++ {
		m[index] = index
	}
}

func MapCapacityFoo() {
	m := make(map[int]int)
	fmt.Printf("m = %v, len = %v, &m = %p\n", m, len(m), &m)
	upperCapacityMapFunc(m, 8)
	fmt.Printf("m = %v, len = %v, &m = %p\n", m, len(m), &m)
}
