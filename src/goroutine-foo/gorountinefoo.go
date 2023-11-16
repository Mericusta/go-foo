package goroutinefoo

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func OpenSoMuchGoRoutine() {
	c1, c2 := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(math.MaxInt16 * 8)
	for index := 0; index != math.MaxInt16*8; index++ {
		go func() {
			// block here
			<-c1
			c2 <- true
			wg.Done()
		}()
	}
	fmt.Printf("start sleep 1 minute, it will alloc 2GB memory\n")
	fmt.Printf("but if you terminate the process, it will release\n")
	t := time.NewTimer(time.Minute)
	<-t.C
	c1 <- true
	wg.Wait()
}

// before go1.14 it will stack
// [Go 语言原本 - 6.8 协作与抢占](https://golang.design/under-the-hood/zh-cn/part2runtime/ch06sched/preemption/)
func AllMIsWorking() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	panic(x)
}

// goroutine 在 panic 和 recover 的表现
func GoroutinePanicAndRecoverFoo() {
	// prepare signal control
	wg := sync.WaitGroup{}

	// condition 1: 同协程，栈顶函数里 panic 并 recover
	// recover at: g1 f1
	// panic at: g1 f1
	// result: g1 end and recover
	wg.Add(1)
	go func(t string) {
		fmt.Printf("condition 1\n")
		fmt.Printf("this is %v\n", t)
		defer func() {
			if panicInfo := recover(); panicInfo != nil {
				fmt.Printf("panic info: %v\n", panicInfo)
			}
			fmt.Printf("%v is end\n", t)
			wg.Done()
		}()
		panic(fmt.Sprintf("panic at %v", t))
	}("root goroutine")
	wg.Wait()
	fmt.Println()

	// condition 2: 同协程，子函数里 panic 并 recover
	// recover at: g1 sub-call f2
	// panic at: g1 sub-call f2
	// result: g1 not end, f2 end and recover
	wg.Add(1)
	go func(t string) {
		fmt.Printf("condition 2\n")
		fmt.Printf("this is %v\n", t)
		func(t string) {
			fmt.Printf("this is %v\n", t)
			defer func() {
				if panicInfo := recover(); panicInfo != nil {
					fmt.Printf("panic info: %v\n", panicInfo)
				}
				fmt.Printf("%v is end\n", t)
			}()

			func() {
				panic(fmt.Sprintf("panic at %v", t))
			}()
		}(fmt.Sprintf("%v sub-call", t))
		time.Sleep(time.Second * 2)
		fmt.Printf("%v is end\n", t)
		wg.Done()
	}("root goroutine")
	wg.Wait()
	fmt.Println()

	// condition 3: 同协程，子函数里面 panic，栈顶函数里面 recover
	// recover at: g1 f1
	// panic at: g1 sub-call f2
	// result: g1 end, and after expreesion f2 can not execute
	wg.Add(1)
	go func(t string) {
		fmt.Printf("condition 3\n")
		fmt.Printf("this is %v\n", t)
		defer func() {
			if panicInfo := recover(); panicInfo != nil {
				fmt.Printf("panic info: %v\n", panicInfo)
			}
			fmt.Printf("%v is end\n", t)
			wg.Done()
		}()
		func(t string) {
			fmt.Printf("this is %v\n", t)
			panic(fmt.Sprintf("panic at %v", t))
		}(fmt.Sprintf("%v sub-call", t))
		time.Sleep(time.Second * 2)
		fmt.Printf("%v is sleep after 2 seconds\n", t)
	}("root goroutine")
	wg.Wait()
	fmt.Println()

	// condition 4: 不同协程，子协程栈顶函数里面 panic，父协程栈顶函数里面 recover
	// recover at: g1 f1
	// panic at: sub-goroutine g2 f1
	// result: sub-goroutine g2 end and because g2 callstack has not recover, it will crash the program
	wg.Add(2)
	go func(t string) {
		fmt.Printf("condition 4\n")
		fmt.Printf("this is %v\n", t)
		defer func() {
			if panicInfo := recover(); panicInfo != nil {
				fmt.Printf("panic info: %v", panicInfo)
			}
			fmt.Printf("%v is end\n", t)
			wg.Done()
		}()
		go func(t string) {
			defer wg.Done()
			fmt.Printf("this is %v\n", t)
			// panic(fmt.Sprintf("panic at %v", t))
			fmt.Printf("%v callstack has not recover, it will crash the program\n", t)
		}("sub goroutine")
		time.Sleep(time.Second * 2)
		fmt.Printf("%v is sleep after 2 seconds\n", t)
	}("root goroutine")
	wg.Wait()
	fmt.Println()

	// time.Sleep(time.Second * 5)
}

// RecoverAtHandler 在函数调用栈上 recover
func RecoverAtHandler(gCount, hCount, mod int, concurrently bool) (int64, int32) {
	handleCounter := int64(0)
	recoverCounter := int32(0)
	var wg sync.WaitGroup
	if concurrently {
		wg = sync.WaitGroup{}
		wg.Add(gCount)
	}

	h := func(i int) {
		defer func() {
			if e := recover(); e != nil { // 空 panic 消耗
				atomic.AddInt32(&recoverCounter, 1)
			}
		}()

		if i != 0 && i%mod == 0 {
			panic(i)
		} else {
			atomic.AddInt64(&handleCounter, 1)
		}
	}

	g := func() {
		for i := 0; i != hCount; i++ {
			h(i)
		}

		if concurrently {
			wg.Done()
		}
	}

	for i := 0; i != gCount; i++ {
		if concurrently {
			go g()
		} else {
			g()
		}
	}

	if concurrently {
		wg.Wait()
	}

	return handleCounter, recoverCounter
}

// RecoverAtGoroutine 在协程（函数调用栈顶）上 recover
func RecoverAtGoroutine(gCount, hCount, mod int, concurrently bool) (int64, int32) {
	handleCounter := int64(0)
	recoverCounter := int32(0)
	var wg sync.WaitGroup
	if concurrently {
		wg = sync.WaitGroup{}
		wg.Add(gCount)
	}

	h := func(i int) {
		if i != 0 && i%mod == 0 {
			panic(i)
		} else {
			atomic.AddInt64(&handleCounter, 1)
		}
	}

	var g func(ri int)

	g = func(ri int) {
		defer func() {
			if e := recover(); e != nil {
				atomic.AddInt32(&recoverCounter, 1)
				go g(e.(int) + 1)
			}
		}()

		for i := ri; i != hCount; i++ {
			h(i)
		}

		if concurrently {
			wg.Done() // if h panic, it will skip here
		}
	}

	for i := 0; i != gCount; i++ {
		if concurrently {
			go g(0)
		} else {
			g(0)
		}
	}

	if concurrently {
		wg.Wait()
	}

	return handleCounter, recoverCounter
}

// it will cost all of CPU
func ConcurrentlyReadWriteSlice() {
	s := make([]int, 0, 8)
	for v := 0; v != 8; v++ {
		s = append(s, v)
	}
	g1 := func(_i int, _s []int) {
		for {
			for i, v := range _s {
				fmt.Printf("g %v, i %v, v %v\n", _i, i, v)
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
	g2 := func(_i int, _s []int) {
		for {
			for i := 0; i < len(_s); i++ {
				_s[i] += _i
			}
			time.Sleep(time.Millisecond * 10)
		}
	}

	for index := 0; index != 100; index++ {
		go g1(index, s)
		go g2(index, s)
	}

	select {}
}

// fatal error, will kill process
func ConcurrentlyReadWriteMap() {
	m := make(map[int]int)
	for v := 0; v != 8; v++ {
		m[v] = v
	}
	g1 := func(_i int, _m map[int]int) {
		for {
			for i, v := range _m {
				fmt.Printf("g %v, i %v, v %v\n", _i, i, v)
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
	g2 := func(_i int, _m map[int]int) {
		for {
			for i := 0; i != 8; i++ {
				_m[i] += _i
			}
			time.Sleep(time.Millisecond * 10)
		}
	}

	for index := 0; index != 100; index++ {
		go g1(index, m)
		go g2(index, m)
	}

	select {}
}
