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
