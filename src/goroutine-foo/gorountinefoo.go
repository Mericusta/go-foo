package goroutinefoo

import (
	"fmt"
	"math"
	"runtime"
	"sync"
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
