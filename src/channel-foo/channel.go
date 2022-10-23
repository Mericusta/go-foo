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

// 使用 select case 语法处理发送的 channel
func GoSelectSendChannel() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := make(chan int)
	go func() {
		for index := 0; index != 3; index++ {
			fmt.Printf("send %v and block\n", index)
			select {
			case c <- index:
				fmt.Printf("send %v done\n", index)
			}
		}
		wg.Done()
	}()

	time.Sleep(time.Second)

	go func() {
		for index := 0; index != 3; index++ {
			time.Sleep(time.Second)
			select {
			case v, ok := <-c:
				fmt.Printf("recv %v done\n", v)
				if !ok {
					break
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

// https://github.com/kubernetes/kubernetes/blob/7509c4eb478a3ab94ff26be2b4068da53212d538/pkg/controller/nodelifecycle/scheduler/taint_manager.go#L244
func PriorityChannel() {

}

// select case 中同时存在已关闭和未关闭的 channel
func SelectClosedAndUnclosedChannel() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	const BUFFER_CHANNEL_LEN = 16

	ctx1, canceler1 := context.WithCancel(context.Background())
	recvChannel1 := make(chan int, BUFFER_CHANNEL_LEN)

	// situation 1: cancel but recv channel has release value
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
	}(recvChannel1, canceler1)

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
	}(ctx1, recvChannel1)

	wg.Wait()

	// ctx2, canceler2 := context.WithCancel(context.Background())
	// recvChannel2 := make(chan int, 16)
	// go func(sendChannel chan<- int) {
	// 	ticker := time.NewTicker(time.Second)
	// 	for range ticker.C {
	// 		sendChannel <- 1
	// 	}
	// 	fmt.Printf("send ticker 2 done\n")
	// }(recvChannel2)

	// go func(ctx context.Context, recvChannel <-chan int) {
	// 	ticker := time.NewTicker(time.Second)
	// LOOP:
	// 	for {
	// 		select {
	// 		case <-ticker.C: // 主动发送
	// 		priority:
	// 			for {
	// 				select {
	// 				case <-ctx.Done(): // 主动结束
	// 					ticker.Stop()
	// 					fmt.Printf("Note: stop ticker\n")
	// 					goto LOOP
	// 				default:
	// 					break priority
	// 				}
	// 			}
	// 			// 发送逻辑
	// 			fmt.Printf("Note: ticker send logic\n")
	// 		case i, ok := <-recvChannel: // 被动接收
	// 			if !ok { // 被动结束
	// 				fmt.Printf("Note: receive channel closed\n")
	// 				ticker.Stop()
	// 				break LOOP
	// 			}
	// 			// 接收逻辑
	// 			fmt.Printf("Note: receive value from channel %v\n", i)
	// 		}
	// 	}
	// }(ctx2, recvChannel2)
}
