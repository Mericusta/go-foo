package contextfoo

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用 channel 的方式停止工作中的协程
func stopGoroutineWay1() {
	stopChan := make(chan bool)

	go func() {
		second := 0
		for {
			select {
			case <-stopChan:
				fmt.Println("receive stop channel")
				return
			default:
				second++
				time.Sleep(time.Second * 1)
				fmt.Println("channel monitoring, second =", second)
			}
		}
	}()

	time.Sleep(time.Second * 18)
	fmt.Println("send stop channel")
	stopChan <- true
	time.Sleep(time.Second * 10)
}

// 使用 cancel context 的方式停止工作中的协程
func stopGoroutineWay2() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		second := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("receive context done channel")
				return
			default:
				second++
				time.Sleep(time.Second * 1)
				fmt.Println("channel monitoring, second =", second)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 18)
	fmt.Println("send stop channel")
	cancel()
	time.Sleep(time.Second * 10)
}

// 使用 cancel context 停止多个工作中的协程
func monitorMultiGoroutineWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	watcher := func(ctx context.Context, index int) {
		second := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("receive context done channel on watcher", index)
				return
			default:
				second++
				time.Sleep(time.Second * 1)
				fmt.Println("channel", index, "monitoring, second =", second)
			}
		}
	}
	for index := 0; index != 3; index++ {
		go watcher(ctx, index)
	}

	time.Sleep(time.Second * 18)
	fmt.Println("send stop channel")
	cancel()
	time.Sleep(time.Second * 10)
}

type contextKey int
type contextValue int

// 通过 value context 传递上下文数据
func monitorGoroutineWithContextAndValue() {
	ctx, cancel := context.WithCancel(context.Background())
	for index := 0; index != 3; index++ {
		ctx = context.WithValue(ctx, contextKey(index), contextValue(index))
	}
	go func(ctx context.Context) {
		second := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("receive context done channel")
				return
			default:
				second++
				time.Sleep(time.Second * 1)
				fmt.Println("channel monitoring, second =", second)
				for key := 0; key != 3; key++ {
					fmt.Println("context key =", key, "context value =", ctx.Value(contextKey(key)))
				}
			}
		}
	}(ctx)

	time.Sleep(time.Second * 7)
	fmt.Println("send stop channel")
	cancel()
	time.Sleep(time.Second * 4)
}

// 使用 channel 的方式停止工作中的协程树
// - 主协程
// - 子协程1：日志协程
// - 子协程2：socket 协程
// - 子协程2-1：socket read 协程
// - 子协程2-2: socket write 协程
func contextTreeCloseWay1() {
	stopSingal := make(chan bool)

	go logGoroutine(stopSingal)
	go socketGoroutine(stopSingal)

	time.Sleep(10 * time.Second)
	fmt.Println("main send stop signal to all goroutine")
	close(stopSingal)
	time.Sleep(5 * time.Second)

	fmt.Println("main context exit")
}

// 日志协程
func logGoroutine(stopSingal chan bool) {
	for {
		select {
		case <-stopSingal:
			fmt.Println("log routine receive stop signal")
			return
		default:
			fmt.Println("run logger logic ...")
			fmt.Println("maybe create sub routine for logic io/backup/sync ...")
			time.Sleep(time.Second * 5)
		}
	}
}

// socket 协程
func socketGoroutine(stopSingal chan bool) {
	second := 0
	socketNum := 0
	socketMap := make(map[int]struct {
		readSignal  chan int
		writeSignal chan int
	})
	for {
		select {
		case <-stopSingal:
			fmt.Println("socket routine receive stop signal")
			return
		default:
			fmt.Println("run socket logic ...")
			fmt.Println("create read and write routine for each socket ...")
			if socketNum < 3 {
				readSignal := make(chan int)
				writeSignal := make(chan int)
				socketMap[socketNum] = struct {
					readSignal  chan int
					writeSignal chan int
				}{
					readSignal:  readSignal,
					writeSignal: writeSignal,
				}
				go socketRead(socketNum, readSignal, stopSingal)
				go socketWrite(socketNum, writeSignal, stopSingal)
				socketNum++
			} else {
				for socketIndex, socketStruct := range socketMap {
					fmt.Println("os receive read signal value", second, "to socket", socketIndex, "from Internet")
					socketStruct.readSignal <- second
					fmt.Println("os receive write signal value", second*2, "from socket", socketIndex, "to Internet")
					socketStruct.writeSignal <- second * 2
				}
				time.Sleep(time.Second * 5)
				second += 5
			}
		}
	}
}

// socket read 协程
func socketRead(index int, readSignal chan int, stopSingal chan bool) {
	for {
		select {
		case <-stopSingal:
			fmt.Println("read socket", index, "receive stop signal")
			return
		case v := <-readSignal:
			fmt.Println("socket", index, "receive read signal", v, "from os")
		}
	}
}

// socket write 协程
func socketWrite(index int, writeSignal chan int, stopSingal chan bool) {
	for {
		select {
		case <-stopSingal:
			fmt.Println("write socket", index, "receive stop signal")
			return
		case v := <-writeSignal:
			fmt.Println("socket", index, "receive write signal", v, "from program")
		}
	}
}

// 使用 cancel context 的方式停止工作中的协程树
// - 主协程
// - 子协程1：日志协程
// - 子协程2：socket 协程
// - 子协程2-1：socket read 协程
// - 子协程2-2: socket write 协程
func contextTreeCloseWay2() {
	ctx, cancel := context.WithCancel(context.Background())

	go logGoroutineWithContext(ctx)
	go socketGoroutineWithContext(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("main cancel socket goroutine")
	// cancel()
	time.Sleep(5 * time.Second)

	time.Sleep(10 * time.Second)
	fmt.Println("main call cancel")
	cancel()
	time.Sleep(5 * time.Second)

	fmt.Println("main context exit")
}

func logGoroutineWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("log routine receive stop signal")
			return
		default:
			fmt.Println("run logger logic ...")
			fmt.Println("maybe create sub routine for logic io/backup/sync ...")
			time.Sleep(time.Second * 5)
		}
	}
}

func socketGoroutineWithContext(ctx context.Context) {
	second := 0
	socketNum := 0
	socketMap := make(map[int]struct {
		readSignal  chan int
		writeSignal chan int
	})
	for {
		select {
		case <-ctx.Done():
			fmt.Println("socket routine receive stop signal")
			return
		default:
			fmt.Println("run socket logic ...")
			fmt.Println("create read and write routine for each socket ...")
			if socketNum < 3 {
				readSignal := make(chan int)
				writeSignal := make(chan int)
				socketMap[socketNum] = struct {
					readSignal  chan int
					writeSignal chan int
				}{
					readSignal:  readSignal,
					writeSignal: writeSignal,
				}
				go socketReadWithContext(ctx, socketNum, readSignal)
				go socketWriteWithContext(ctx, socketNum, writeSignal)
				socketNum++
			} else {
				for socketIndex, socketStruct := range socketMap {
					fmt.Println("os receive read signal value", second, "to socket", socketIndex, "from Internet")
					socketStruct.readSignal <- second
					fmt.Println("os receive write signal value", second*2, "from socket", socketIndex, "to Internet")
					socketStruct.writeSignal <- second * 2
				}
				time.Sleep(time.Second * 5)
				second += 5
			}
		}
	}
}

func socketReadWithContext(ctx context.Context, index int, readSignal chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("read socket", index, "receive stop signal")
			return
		case v := <-readSignal:
			fmt.Println("socket", index, "receive read signal", v, "from os")
		}
	}
}

func socketWriteWithContext(ctx context.Context, index int, writeSignal chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("write socket", index, "receive stop signal")
			return
		case v := <-writeSignal:
			fmt.Println("socket", index, "receive write signal", v, "from program")
		}
	}
}

// timeout

type fooError struct {
	e error
}

func TimeoutContextFoo(timeoutSeconds, businessSeconds int, businessPanic bool) *fooError {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ctx, canceler := context.WithTimeout(context.Background(), time.Second*time.Duration(timeoutSeconds))

	defer func() {
		fmt.Println("timeoutContextFoo defer")
	}()

	workerFunc := func() {
		fmt.Printf("in worker, long time business begin, %vs\n", businessSeconds)
		if businessPanic {
			fmt.Println("in worker, long time business occurs panic")
			panic("business panic")
		} else {
			time.Sleep(time.Second * time.Duration(businessSeconds))
			fmt.Println("in worker, long time business done")
		}
	}

	workerDone, workerPanic := make(chan struct{}), make(chan any, 1)
	workerGoroutine := func(_ctx context.Context, workerFunc func()) {
		defer func() {
			if p := recover(); p != nil {
				workerPanic <- p
			}
			fmt.Println("workerGoroutine, defer")
			wg.Done()
		}()
		workerFunc()
		close(workerDone)
	}

	guardGoroutine := func(_ctx context.Context, canceler context.CancelFunc) {
		defer func() {
			fmt.Println("guardGoroutine, defer")
			wg.Done()
		}()

		go workerGoroutine(_ctx, workerFunc)
		select {
		case <-workerDone:
			fmt.Println("in guard, business complete on time")
			break
		case p := <-workerPanic:
			fmt.Printf("in guard, business occurs panic, %v\n", p)
			canceler()
			break
		case <-_ctx.Done():
			fmt.Println("in guard, business is overtime, call cancel")
			break
		}
	}

	go guardGoroutine(ctx, canceler)

	wg.Wait()

	return &fooError{e: ctx.Err()}
}

func ControlContextTree() {

}
