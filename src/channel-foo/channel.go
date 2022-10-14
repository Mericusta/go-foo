package channelfoo

import (
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
