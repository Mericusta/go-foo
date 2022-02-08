package main

import (
	"fmt"
	channelfoo "go-foo/channel-foo"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Bencher(count int, f func(int), concurrently bool) {
	var wg sync.WaitGroup
	if concurrently {
		wg = sync.WaitGroup{}
		wg.Add(count)
	}
	t1 := time.Now()
	for index := 0; index != count; index++ {
		if concurrently {
			go func() {
				f(index)
				wg.Done()
			}()
		} else {
			f(index)
		}
	}
	if concurrently {
		wg.Wait()
	}
	t2 := time.Now()
	fmt.Printf("using time: %v milli-seconds\n", t2.Sub(t1).Milliseconds())
}

func main() {
	counterMapLocker := &sync.Mutex{}
	xCounterMap, yCounterMap := make(map[int]int), make(map[int]int)
	Bencher(100, func(index int) {
		x, y := channelfoo.GoroutineOutputOrder()
		counterMapLocker.Lock()
		xCounterMap[x]++
		yCounterMap[y]++
		counterMapLocker.Unlock()
	}, false)
	fmt.Printf("x counter map = %v\n", xCounterMap)
	fmt.Printf("y counter map = %v\n", yCounterMap)
}
