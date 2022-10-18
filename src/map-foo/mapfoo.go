package mapfoo

import "fmt"

// 不同的 count 和 capacity 组合在 benchmark 下的性能是不同的
// 只讨论 count <= capacity 的情况，见 ./BenchmarkMapCapacityFoo.sh true 10000
func MapCapacityFoo(count, capacity int) {
	m := make(map[int]int, capacity)
	for index := 0; index != count; index++ {
		m[index] = index
	}
}

type mapKey struct {
	k1  int
	k2  int
	ptr *bool
}

// 非 ptr 类型的 key 会比较其中的所有值，若含有指针，则比较指针的值，而非指针指向的值
func StructMapKeyFoo() {
	structKeyMap := make(map[mapKey]bool)
	ptrSlice := make([]*bool, 10)
	for index := 0; index != 10; index++ {
		v := index%2 == 0
		ptrSlice[10-index-1] = &v
		structKeyMap[mapKey{k1: index, k2: index, ptr: &v}] = v
	}
	fmt.Printf("ptrSlice = %v\n", ptrSlice)
	v, has := structKeyMap[mapKey{k1: 2, k2: 2, ptr: ptrSlice[7]}] // pointer and pointer to value both same
	fmt.Printf("v, has = %v, %v, %v, %v\n", v, has, ptrSlice[7], *ptrSlice[7])
	v, has = structKeyMap[mapKey{k1: 2, k2: 2, ptr: ptrSlice[5]}] // pointer to value is same, pointer is different
	fmt.Printf("v, has = %v, %v, %v, %v\n", v, has, ptrSlice[5], *ptrSlice[5])
	*ptrSlice[7] = !*ptrSlice[7]
	v, has = structKeyMap[mapKey{k1: 2, k2: 2, ptr: ptrSlice[7]}] // pointer is same, pointer to value is different
	fmt.Printf("v, has = %v, %v, %v, %v\n", v, has, ptrSlice[7], *ptrSlice[7])
}

func GetFromMapAsTypeEmptyValueFoo() {
	m1 := make(map[int]int)
	m1[1] = 10
	v1, has := m1[2] // not exist key, return type int empty value 0
	if v1 != 0 && has {
		panic(fmt.Sprintf("%v %v", v1, has))
	}

	m2 := make(map[int][]int)
	m2[1] = []int{10, 20, 30}
	v2, has := m2[2] // not exist key, return type int empty value nil
	if v2 != nil && has {
		panic(fmt.Sprintf("%v %v", v2, has))
	}
}
