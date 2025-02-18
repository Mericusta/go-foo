package goroutinefoo

import (
	"fmt"
	"math"
	"math/rand"
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

// interrupt routine 协程中断方式模拟
// - mode
// - 1: async/await
// - 2: select/case
func interruptRoutineFoo(mode int) {
	switch mode {
	case 1:
		asyncAwaitInterruptRoutine()
	case 2:
		selectCaseInterruptRoutine()
	}
}

func asyncAwaitInterruptRoutine() {
	// 协程状态
	type routineState int8
	const (
		routine_init       = iota // 协程初始化
		routine_scheduling        // 协程调度中
		routine_suspended         // 协程挂起中
	)
	// 协程 -> 屏蔽细节，每次调度都从函数入口处开始调度而不是从调用栈中间恢复
	type routine struct {
		state    routineState // 协程状态
		callback func()       // 协程入口函数
	}
	// 协程初始化函数
	initRoutine := func(r *routine) { // method -> func
		// 初始化协程操作：分配栈内存等
		// ...

		// 设置协程状态：空闲
		r.state = routine_scheduling
	}
	// 检查协程是否需要挂起
	checkRoutineState := func(r *routine) routineState { // method -> func
		// 检查是否阻塞等待
		if rand.Intn(10) < 5 {
			r.state = routine_suspended
		} else {
			r.state = routine_scheduling
		}
		return r.state
	}
	// 继续协程
	continueRoutine := func(r *routine) { // method -> func
		// 恢复调用栈，恢复栈内存
		// ...
	}
	// 终止协程
	interruptRoutine := func(r *routine) { // method -> func
		// 清空调用栈，释放栈内存
		// ...
	}

	// 协程调度器 -> 类似 GPM 的 P，负责切换协程状态
	type routineManager struct {
		toScheduleRoutine []*routine // 待调度的协程
		toCancelRoutine   []*routine // 待中断的协程
	}
	// 协程调度函数
	scheduleRoutine := func(rm *routineManager, r *routine) { // method -> func
		if r != nil {
			rm.toScheduleRoutine = append(rm.toScheduleRoutine, r)
		}
	}
	// 取消调度协程
	cancelRoutine := func(rm *routineManager, r *routine) { // method -> func
		for _, _r := range rm.toCancelRoutine {
			if _r == r {
				return
			}
		}
		toContinueRoutine := make([]*routine, 0, 16)
		for _, _r := range rm.toScheduleRoutine {
			if _r == r {
				rm.toCancelRoutine = append(rm.toCancelRoutine, _r)
			} else {
				toContinueRoutine = append(toContinueRoutine, _r)
			}
		}
		rm.toScheduleRoutine = toContinueRoutine
	}
	// 协程调度器运行函数
	runRoutine := func(rm *routineManager) { // method -> func
		ticker := time.NewTicker(time.Second) // 放大 CPU 调度间隔
		for range ticker.C {
			// async/await interrupt 机制：调度时检测是否需要取消调度
			for _, _r := range rm.toCancelRoutine {
				// 终止协程
				interruptRoutine(_r)
			}

			// 继续调度剩下的协程：抢占式调度
			for _, _r := range rm.toScheduleRoutine {
				// 检查状态
				rs := checkRoutineState(_r)
				if rs == routine_scheduling {
					continueRoutine(_r)
				}
			}
		}
	}
	// 协程调度器实例
	rm := &routineManager{toScheduleRoutine: make([]*routine, 0, 16), toCancelRoutine: make([]*routine, 0, 16)}

	// 协程启动函数 -> await/go 关键字
	awaitRoutine := func(rm *routineManager, f func()) *routine {
		r := &routine{state: routine_init, callback: f} // 设置协程状态：协程初始化
		initRoutine(r)
		scheduleRoutine(rm, r)
		return r
	}

	f := func() { fmt.Println("hello async/await") } // 协程入口函数
	go runRoutine(rm)                                // 运行协程调度器
	r := awaitRoutine(rm, f)                         // 创建并调度协程
	cancelRoutine(rm, r)                             // 取消调度协程
}

func selectCaseInterruptRoutine() {
	var (
		cancelSignal   bool = false
		topOfStackFunc      = func() {
			// 调用栈内多次判断是否需要中断
			if cancelSignal { // 第一次判断
				return
			}

			fmt.Println("do something ...")

			if cancelSignal { // 第二次判断
				return
			}

			fmt.Println("do something again ...")

			if cancelSignal { // 第三次判断
				return
			}

			fmt.Println("do something finally ...")
		}
		bottomOfStackFunc = func() { // 调用栈底层函数
			// 中断方式1：函数内轮询，类似 select-case-ctx.Done
			for {
				if cancelSignal { // 是否需要中断
					return
				}

				// 中断方式2：步进，类似断点调试
				// 入栈
				topOfStackFunc()
			}
		}
	)

	// 协程入口函数
	bottomOfStackFunc()
}
