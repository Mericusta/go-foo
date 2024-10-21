package channelfoo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 发送协程
// - 通过 receiverGoroutineNoticeChan 通知接收协程数据变化
// - 通过 receiverGoroutineExitChan 通知接收协程退出
// - 随机时间点发送退出通知并结束协程
// 接收协程
// - 监听 receiverGoroutineOpenChan
// - 监听 receiverGoroutineNoticeChan
// - 监听 receiverGoroutineExitChan
// - 通过 mainGoroutineExitChan 通知主协程退出
// - 退出时通知主协程退出并结束协程
// GoroutineExitThenCloseChannel 接收协程监听已关闭 channel 的表现
// 从一个 nil channel 中接收数据会一直被 block
// 从一个被 close 的 channel 中接收数据不会被阻塞，而是立即返回，接收完已发送的数据后会返回元素类型的零值(zero value)
func GoroutineExitThenCloseChannel() {
	receiverGoroutineExitChan := make(chan struct{})
	receiverGoroutineNoticeChan := make(chan int)
	receiverGoroutineOpenChan := make(chan int)
	mainGoroutineExitChan := make(chan struct{})

	go func() {
		defer func() {
			mainGoroutineExitChan <- struct{}{}
			close(mainGoroutineExitChan)
		}()
		loopCounter := 0
		waitingExit := false
		for {
			select {
			case _, ok := <-receiverGoroutineOpenChan:
				if waitingExit {
					loopCounter++
					fmt.Println("receiver go routine continue at case <-receiverGoroutineOpenChan")
					continue
				}
				if !ok {
					waitingExit = true
				}
			case v, ok := <-receiverGoroutineNoticeChan:
				if waitingExit {
					loopCounter++
					fmt.Println("receiver go routine continue at case <-receiverGoroutineNoticeChan")
					continue
				}
				if !ok {
					fmt.Println("receiver go routine receive sender go routine notice chan is closed")
					waitingExit = true
					continue
				}
				fmt.Println("receiver go routine receive sender go routine notice value", v)
			case <-receiverGoroutineExitChan:
				fmt.Println("receiver go routine receive sender go routine exit at loop", loopCounter)
				return
			}
		}
	}()

	go func() {
		totalIndex := 100
		exitIndex := rand.Intn(totalIndex / 10)
		fmt.Println("sender go routine rand exit index", exitIndex)
		for index := 0; index != totalIndex; index++ {
			receiverGoroutineNoticeChan <- index
			time.Sleep(time.Millisecond * 100)
			if index == exitIndex {
				break
			}
		}
		// receiverGoroutineExitChan <- struct{}{} // 在这发送退出信号不会导致接收协程空转
		close(receiverGoroutineNoticeChan)
		receiverGoroutineNoticeChan = nil       // 如果无法在 close(receiverGoroutineNoticeChan) 前发送退出信号，则需要通过“置空”来禁用监听该 channel 的 select case
		time.Sleep(time.Millisecond)            // 模拟业务等待时间
		receiverGoroutineExitChan <- struct{}{} // 在这发送退出信号会导致接收协程空转
		close(receiverGoroutineExitChan)
	}()

	<-mainGoroutineExitChan
	close(receiverGoroutineOpenChan)
}

// 发送协程
// - 通过 receiverGoroutineExitChan 通知接收协程退出
// - 2秒后发送退出通知并结束协程
// 接收协程
// - 监听 receiverGoroutineExitChan
// - 监听 定时器 t
// - 通过 mainGoroutineExitChan 通知主协程退出
// - 定时器1秒后进入定时器分支并休眠2秒模拟业务执行
// ListenerBlockedChannel 接收协程阻塞时发送协程发送的表现
func ListenerBlockedChannel() {
	receiverGoroutineExitChan := make(chan struct{})
	mainGoroutineExitChan := make(chan struct{})
	waitReceiver := true

	go func() {
		time.Sleep(time.Second * time.Duration(2))
		if waitReceiver {
			receiverGoroutineExitChan <- struct{}{}
			fmt.Println("sender go routine send exit")
		} else {
			select {
			case receiverGoroutineExitChan <- struct{}{}:
				fmt.Println("sender go routine send exit")
			default:
				fmt.Println("sender go routine send exit failed")
			}
		}
		fmt.Println("sender go routine close receiver go routine exit chan")
		close(receiverGoroutineExitChan)
	}()

	go func() {
		t := time.NewTimer(time.Second * time.Duration(1))
		for {
			select {
			case <-receiverGoroutineExitChan:
				fmt.Println("receiver go routine receive sender go routine exit")
				mainGoroutineExitChan <- struct{}{}
				return
			case <-t.C:
				fmt.Println("receiver go routine sleep")
				time.Sleep(time.Second * time.Duration(2))
			}
		}
	}()

	<-mainGoroutineExitChan
	time.Sleep(time.Second * time.Duration(3))
	close(mainGoroutineExitChan)
}

func GoroutineExitThenCloseChannelSimpleCase() {
	c := make(chan int)
	close(c)

	loopCounter := 0
	for {
		loopCounter++
		if loopCounter > 100 {
			fmt.Println("return at loop counter greater than 100")
			return
		}
		select {
		case _, ok := <-c:
			if !ok {
				continue
			}
		}
	}
}

func GoroutineOutputOrder() (int, int) {
	s := []int{1, 2, 3, -1, -2, -3}
	c := make(chan int)
	sumFunc := func(s []int, sc chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		sc <- sum
	}
	go sumFunc(s[:len(s)/2], c)
	go sumFunc(s[len(s)/2:], c)
	x, y := <-c, <-c
	return x, y
}

// 两个协程交叉按顺序打印 1~10
func GoroutineOutputOrder2() {
	// s := strings.Builder{}

	// c1 := make(chan int)
	// wg := sync.WaitGroup{}
	// wg.Add(2)
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	for i := 0; i < 10; i += 2 {
	// 		fmt.Printf("g1 c1 <- 1 start\n")
	// 		c1 <- 1
	// 		fmt.Printf("g1 c1 <- 1 done\n")
	// 		fmt.Printf("g1 output %v\n", i)
	// 	}
	// 	wg.Done()
	// }()
	// go func() {
	// 	for i := 1; i < 10; i += 2 {
	// 		fmt.Printf("g2 <-c1 start\n")
	// 		<-c1
	// 		fmt.Printf("g2 <-c1 done\n")
	// 		fmt.Printf("g2 output %v\n", i)
	// 	}
	// 	wg.Done()
	// }()
	// wg.Wait()

	s := make([]int, 0, 10)

	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i += 2 {
			c1 <- 1
			// fmt.Printf("g1 %v\n", i)
			if i%2 == 0 {
				fmt.Printf("g1 even %v\n", i)
				s = append(s, i)
			} else {
				fmt.Printf("g1 odd %v\n", i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i < 10; i += 2 {
			<-c1
			// fmt.Printf("g2 %v\n", i)
			if i%2 == 1 {
				fmt.Printf("g2 odd %v\n", i)
				s = append(s, i)
			} else {
				fmt.Printf("g2 even %v\n", i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println()
	fmt.Printf("%v\n", s)
}

// channel 引发 goroutine 阻塞，但不会导致 M 阻塞，goroutine 只是被挂起
func GoChannelBlock() {
	s := make(chan bool)
	sendGo := func() {
		s <- true
	}

	go sendGo()

	select {
	case <-s:
		fmt.Printf("receive init \n")
	}
}

// https://github.com/kubernetes/kubernetes/blob/7509c4eb478a3ab94ff26be2b4068da53212d538/pkg/controller/nodelifecycle/scheduler/taint_manager.go#L244
func PriorityChannel() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	ctx, canceler := context.WithCancel(context.Background())
	eventChannel := make(chan int, 1024)
	recvChannel := make(chan int, 1024)
	// priority sort: context > recvChannel > eventChannel
	// but recvChannel == eventChannel
	go func(ctx context.Context, eventChannel, recvChannel chan int) {
		c := 0
	LOOP:
		for {
			select {
			case v, ok := <-eventChannel: // 主动发送
			priority:
				for {
					select {
					case <-ctx.Done(): // 主动结束，本端主动通过 context cancel 结束，必须保证本端先关闭 event channel
						fmt.Printf("Note: dispatcher receive context done\n")
						break LOOP
					default:
						fmt.Printf("Note: break priority\n")
						break priority
					}
				}
				if !ok {
					c++
					fmt.Printf("Note: dispatcher send channel closed\n")
					// break LOOP // 由于对端断开 tcp 套接字而结束
					if c >= 10 {
						break LOOP
					}
					goto LOOP // 由于本端主动 cancel 而结束
				}
				// 发送逻辑
				fmt.Printf("Note: dispatcher handle eventChannel, receive value %v\n", v)
			case v, ok := <-recvChannel: // 被动接收
				if !ok { // 被动结束，对端断开 tcp 套接字
					fmt.Printf("Note: dispatcher receive channel closed\n")
					close(eventChannel)
					goto LOOP
				}
				// 接收逻辑
				fmt.Printf("Note: dispatcher handle recvChannel, receive value %v\n", v)
			}
		}
		wg.Done()
	}(ctx, eventChannel, recvChannel)

	go func(ctx context.Context, eventChannel chan int) {
		counter := 0
		ticker := time.NewTicker(time.Second)
	LOOP:
		for {
			select {
			case <-ticker.C:
				eventChannel <- 1
				counter++
				if counter == 5 {
					fmt.Printf("event channel counter done\n")
					break LOOP
				}
			case <-ctx.Done():
				fmt.Printf("event channel ctx done\n")
				break LOOP
			}
		}
		fmt.Printf("event channel closed\n")
		close(eventChannel)
		wg.Done()
	}(ctx, eventChannel)

	go func(recvChannel chan int) {
		counter := 0
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			recvChannel <- 2
			counter++
			if counter == 15 {
				break
			}
		}
		wg.Done()
	}(recvChannel)

	// situation 1: 主动关闭
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	canceler() // 主动关闭必须先关闭 event channel
	wg.Wait()
}

// select case 中同时存在已关闭和未关闭的 channel
func SelectClosedAndUnclosedChannelFoo() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	const BUFFER_CHANNEL_LEN = 16

	// 主动 cancel 但是 recv channel 还有值
	ctx, canceler := context.WithCancel(context.Background())
	recvChannel := make(chan int, BUFFER_CHANNEL_LEN)

	// goroutine 1: 发送 BUFFER_CHANNEL_LEN 数量的值到 channel 里面然后 cancel（应该关闭 channel 但为了营造则不关闭）
	go func(sendChannel chan<- int, canceler context.CancelFunc) {
		counter := 0
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			counter++
			if counter == BUFFER_CHANNEL_LEN {
				break
			}
			sendChannel <- counter
		}
		fmt.Printf("\t- send ticker 1 done and cancel 1\n")
		canceler()
		fmt.Printf("\t- send goroutine close send channel, len = %v\n", len(sendChannel))
		// close(sendChannel)
		wg.Done()
	}(recvChannel, canceler)

	// goroutine 2:
	go func(ctx context.Context, recvChannel chan int) {
		// 等待 BUFFER_CHANNEL_LEN/2 秒之后开始处理 recv channel
		time.Sleep(time.Second * BUFFER_CHANNEL_LEN / 2)
		// 保证首次 tick 有 50% 的概率执行 recv channel，否则 tick 内等待1s，下次又有 50% 概率
		ticker := time.NewTicker(time.Second) // 1s 1 tick
		time.Sleep(time.Second)               // 等待 1s 保证 ticker 就绪，select case 有 50% 概率
	LOOP:
		for {
			fmt.Printf("LOOP begin, recv channel release len %v\n", len(recvChannel))
			select {
			case <-ticker.C: // 主动发送
			priority:
				for {
					select {
					case <-ctx.Done(): // 主动结束
						// 营造此时 channel 中还有值的现象
						fmt.Printf("ctx Done, recv channel len %v\n", len(recvChannel))
						release := BUFFER_CHANNEL_LEN - len(recvChannel)
						for index := 0; index != release; index++ {
							fmt.Printf("ctx Done, mock recv channel value %v, len %v\n", (index+1)*10, len(recvChannel))
							recvChannel <- (index + 1) * 10
						}
						fmt.Printf("Note: stop ticker, recv channel len %v\n", len(recvChannel))
						ticker.Stop()
						wg.Done()
						fmt.Printf("close recv channel\n")
						close(recvChannel) // 应该在 goroutine 1 中关闭，即谁发谁关闭，但这里为了营造环境所以在这里关闭
						fmt.Printf("Note: goto LOOP\n")
						goto LOOP
					default:
						fmt.Printf("break priority\n")
						break priority
					}
				}
				// 发送逻辑
				fmt.Printf("Note: ticker handle send logic\n")
				time.Sleep(time.Second) // 等待 1s 保证 ticker 就绪，下次又有 50% 概率
			case i, ok := <-recvChannel: // 被动接收
				if !ok { // 被动结束
					fmt.Printf("Note: receive channel closed with len %v\n", len(recvChannel))
					ticker.Stop()
					wg.Done()
					break LOOP
				}
				// 接收逻辑
				fmt.Printf("Note: receive value %v from channel\n", i)
			}
		}
	}(ctx, recvChannel)

	wg.Wait()
}

// select case 中同时存在已关闭和未关闭的 channel
func SelectClosedAndUnclosedChannel1() {
	c1 := make(chan int, 16)
	for index := 0; index != 16; index++ {
		c1 <- index
	}
	c2 := make(chan int, 16)
	for index := 0; index != 16; index++ {
		c2 <- index
	}
	// ctx, canceler := context.WithCancel(context.Background())

	// close(c1) // control close
	// close(c2) // can't control close

LOOP:
	for {
		select {
		case v, c1ok := <-c1: // can control
			fmt.Printf("receive c1: %v, %v\n", v, c1ok)
			if !c1ok {
				// 关闭了 c1 之后 select case 仍然可以进来
				// 因为 case 无法检查是否已关闭
				// 所以这里不能 goto LOOP 处理 c2，只能立刻处理
				fmt.Printf("c1 closed\n")
				fmt.Printf("close c2\n")
				close(c2)           // 首先关闭 c2
				for v := range c2 { // 处理 c2 剩余的内容
					fmt.Printf("after c1 closed, c2: %v\n", v)
				}
				fmt.Printf("break LOOP\n")
				break LOOP // 结束
			}
			fmt.Printf("handle c1: %v, %v\n", v, c1ok)
		case v, c2ok := <-c2: // can't control
			fmt.Printf("receive c2: %v, %v\n", v, c2ok)
			if !c2ok {
				// 关闭了 c2 之后 select case 仍然可以进来
				// 因为 case 无法检查是否已关闭
				// 所以这里不能 goto LOOP 处理 c1，只能立刻处理
				fmt.Printf("c2 closed\n")
				fmt.Printf("close c1\n")
				close(c1)           // 首先关闭 c1
				for v := range c1 { // 处理 c1 剩余的内容
					fmt.Printf("after c2 closed, c1: %v\n", v)
				}
				fmt.Printf("break LOOP\n")
				break LOOP // 结束
			}
			fmt.Printf("handle c2: %v, %v\n", v, c2ok)
		}
	}
}

// 多个 goroutine 同时监听相同的 channel
// 写入 channel 的 goroutine 是高速操作
// 读取 channel 的 goroutine 是低速操作
func MultiGoroutineSelectCaseOneChannel(size, count int, handleDuration time.Duration) {
	m := sync.Map{}
	c := make(chan int, size)
	ctx, canceler := context.WithCancel(context.Background())
	for index := 0; index != count; index++ {
		go func(ctx context.Context, i int, c <-chan int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("index %v done\n", i)
					return
				case k, ok := <-c:
					v, has := m.LoadAndDelete(k)
					if !has {
						fmt.Printf("index %v, ok %v, not has k %v\n", i, ok, k)
					} else {
						fmt.Printf("index %v, ok %v k: %v, v %v, has %v\n", i, ok, k, v, has)
					}
					time.Sleep(handleDuration)
				}
			}
		}(ctx, index, c)
	}

	go func(ctx context.Context, c chan<- int) {
		ticker := time.NewTicker(handleDuration * time.Duration(size) / time.Duration(count))
		counter := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("sender done\n")
				return
			case <-ticker.C:
				v, has := m.LoadOrStore(counter, counter)
				if has {
					panic(fmt.Sprintf("m already has key %v value %v", counter, v))
				}
				c <- counter
				counter++
			}
		}
	}(ctx, c)

	time.Sleep(time.Second * 2)
	canceler()
}

type complexStruct struct {
	v int
	s string
}

// 发送复杂数据结构
func SendComplexStructFoo() {
	s := complexStruct{
		v: 10,
		s: "10",
	}
	sp := &complexStruct{
		v: 11,
		s: "11",
	}

	ctx, canceler := context.WithCancel(context.Background())
	recvChan := make(chan interface{})
	changeChan := make(chan bool)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-recvChan:
				if !ok {
					continue
				}
				switch v.(type) {
				case complexStruct:
					cs := v.(complexStruct)
					cs.v++
					cs.s = "ok"
					fmt.Printf("recv struct change %v\n", cs)
					changeChan <- true
				case *complexStruct:
					cs := v.(*complexStruct)
					cs.v++
					cs.s = "ok"
					fmt.Printf("recv struct pointer change %+v\n", cs)
					changeChan <- true
				}
			}
		}
	}(ctx)

	recvChan <- s
	<-changeChan
	fmt.Printf("after change %v\n", s)

	recvChan <- sp
	<-changeChan
	fmt.Printf("after change %+v\n", sp)

	canceler()
}

// 修改被监听的 channel
func ChangeChannelWhichIsSelectedFoo(stack bool) {
	c := make(chan int, 16)
	ctx, canceler := context.WithCancel(context.Background())

	go func(ctx context.Context, c <-chan int) {
		for {
			select {
			case v, ok := <-c:
				fmt.Printf("from channel %v, %v\n", v, ok)
			case <-ctx.Done():
				fmt.Printf("cancel\n")
				return
			}
		}
	}(ctx, c)

	for index := 0; index != 16; index++ {
		c <- index
	}

	time.Sleep(time.Second)

	if stack {
		c = make(chan int)
		for index := 0; index != 16; index++ {
			c <- index // deadlock here
		}
		time.Sleep(time.Second)
	}

	canceler()
}

func MultiSelectReceiverChannel() {
	c := make(chan int)

	// multi receiver
	for index := 0; index != 3; index++ {
		go func(i int) {
			fmt.Printf("index %v wait\n", i)
			<-c
			fmt.Printf("index %v receive\n", i)
		}(index)
	}

	time.Sleep(time.Second)

	// single sender
	c <- 1

	select {}
}

// select case default 实现非阻塞发送
func NoBlockSend() {
	noBufferSendChan := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)

	for index := 0; index != 3; index++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case noBufferSendChan <- i:
					fmt.Println(i, "send done")
					return
				default:
					fmt.Println(i, "send failed")
					time.Sleep(time.Second)
				}
			}
		}(index)
	}

	recvCount := 0
	for {
		recv := <-noBufferSendChan
		fmt.Println("recv", recv)
		recvCount++
		if recvCount >= 3 {
			break
		}
		time.Sleep(time.Second)
		fmt.Println()
	}
	wg.Wait()
}

// 构造带有 buffer 的 chan 后预发送，后开启监听
// 先发送的不会阻塞，后监听的可以处理之前发送的
// 切忌监听时使用 for range 会有引用的问题
func PreSendToBufferedChannel() {
	c := make(chan int, 64)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		index := 0
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			if index < 10 {
				index++
				c <- index
				fmt.Printf("send c %v at %v\n", index, time.Now())
			} else {
				close(c)
				wg.Done()
				return
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		for {
			select {
			case v, ok := <-c:
				if !ok {
					wg.Done()
					return
				}
				fmt.Printf("receive c %v at %v\n", v, time.Now())
			}
		}
	}()

	wg.Wait()
}

// 关闭一个 buffer 中还有数据的 channel 时
// channel 的接收协程的表现
// 还未监听时关闭，则开启监听后依次处理完其中的内容
// 监听时关闭，则依次处理完其中的内容
func closeBufferChannelFoo() {
	count := 16
	c := make(chan int, count)
	for index := 0; index != count; index++ {
		c <- index
	}

	select {
	case c <- 999:
		fmt.Printf("send to full chan ok")
	default:
	}
	fmt.Printf("send to full chan failed\n")

	close(c)

	select {
	case c <- 999:
		fmt.Printf("send to closed chan ok")
	default:
		fmt.Printf("send to closed chan default")
	}
	fmt.Printf("send to closed chan failed")

	for {
		select {
		case v, ok := <-c:
			fmt.Printf("receive v %v, ok %v\n", v, ok)
			if !ok {
				fmt.Printf("receive close\n")
				goto EXIT
			}
		}
	}
EXIT:
}
