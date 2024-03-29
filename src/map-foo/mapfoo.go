package mapfoo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

// 并发读取 map 的数据（在完全禁止写入 map 的情况下）
// 只要发生修改操作，都会 panic（增删改）
func ReadConcurrently(c, s int) {
	m := make(map[int]int)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(c)
	for index := 0; index != c; index++ {
		m[index] = index * 10
	}
	timer := time.NewTimer(time.Second * time.Duration(s))
	for index := 0; index != c; index++ {
		go func(ctx context.Context, i int) {
			counter := 0
			t := time.NewTicker(time.Microsecond * time.Duration(rand.Intn(c)+1))
			for {
				select {
				case <-t.C:
					v, has := m[i]
					if !has {
						panic(fmt.Sprintf("index %v access map but not find value", i))
					}
					if v != i*10 {
						panic(fmt.Sprintf("index %v access map find value wrong %v", i, v))
					}
					counter++
				case <-ctx.Done():
					t.Stop()
					fmt.Printf("time %v index %v counter %v\n", time.Now().UnixNano(), i, counter)
					wg.Done()
					return
				}
			}
		}(ctx, index)
	}
	<-timer.C
	cancel()
	wg.Wait()
}

type mStruct struct {
	i int
	s string
}

// 并发读取 map 的复杂结构体数据，并且修结构体的值（在完全禁止写入 map 的情况下）
// 只要对 map 发生修改操作，都会 panic（增删改）
func ReadComplexDataStructConcurrently(c, s int) {
	m := make(map[int]*mStruct)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(c)
	for index := 0; index != c; index++ {
		m[index] = &mStruct{
			i: index * 10,
			s: fmt.Sprintf("%v*10+0", index),
		}
	}
	timer := time.NewTimer(time.Second * time.Duration(s))
	for index := 0; index != c; index++ {
		go func(ctx context.Context, i int) {
			counter := 0
			t := time.NewTicker(time.Microsecond * time.Duration(rand.Intn(c)+1))
			for {
				select {
				case <-t.C:
					v, has := m[i]
					if !has {
						panic(fmt.Sprintf("index %v access map but not find value", i))
					}
					if v.i != i*10+counter || v.s != fmt.Sprintf("%v*10+%v", i, counter) {
						panic(fmt.Sprintf("index %v access map find value wrong %+v", i, v))
					}
					counter++
					v.i = i*10 + counter
					v.s = fmt.Sprintf("%v*10+%v", i, counter)
				case <-ctx.Done():
					t.Stop()
					fmt.Printf("time %v index %v counter %v\n", time.Now().UnixNano(), i, counter)
					wg.Done()
					return
				}
			}
		}(ctx, index)
	}
	<-timer.C
	cancel()
	wg.Wait()

	for i := 0; i != c; i++ {
		fmt.Printf("i %v, v %+v\n", i, m[i])
	}
}
