package mapfoo

import "fmt"

func upperCapacityMapFunc(count int) map[int]int {
	m := make(map[int]int, 8)
	for index := 0; index != count; index++ {
		m[index] = index
	}
	return m
}

func MapCapacityFoo(c int) {
	// fmt.Printf("m = %v, len = %v, &m = %p\n", m, len(m), &m)
	m := upperCapacityMapFunc(c)
	fmt.Sprintf("%v", len(m))
	// fmt.Printf("m = %v, len = %v, &m = %p\n", m, len(m), &m)
}

type mapKey struct {
	k1 int
	k2 int
}

func StructMapKeyFoo() {
	structKeyMap := make(map[mapKey]bool)
	for index := 0; index != 10; index++ {
		structKeyMap[mapKey{k1: index, k2: index}] = index%2 == 0
	}
	v, has := structKeyMap[mapKey{k1: 5, k2: 5}]
	fmt.Printf("v, has = %v, %v\n", v, has)
}
