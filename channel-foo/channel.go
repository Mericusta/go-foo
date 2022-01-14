package channelfoo

import (
	"fmt"
	"math/rand"
	"time"
)

// 发送协程
// - 通过 receiverGoroutineNoticeChan 通知接收协程数据变化
// - 通过 receiverGoroutineExitChan 通知接收协程退出
// - 随机时间点发送退出通知并结束协程
// 接收协程
// - 监听 receiverGoroutineNoticeChan
// - 监听 receiverGoroutineExitChan
// - 通过 mainGoroutineExitChan 通知主协程退出
// - 退出时通知主协程退出并结束协程
// GoRoutineExitThenCloseChannel 发送协程退出并关闭 channel 时接收协程的表现
func GoRoutineExitThenCloseChannel() {
	receiverGoroutineExitChan := make(chan struct{})
	receiverGoroutineNoticeChan := make(chan int)
	mainGoroutineExitChan := make(chan struct{})

	go func() {
		defer func() {
			mainGoroutineExitChan <- struct{}{}
			close(mainGoroutineExitChan)
		}()
		for {
			select {
			case <-receiverGoroutineExitChan:
				fmt.Println("receiver go routine receive sender go routine exit")
				return
			case v, ok := <-receiverGoroutineNoticeChan:
				if !ok {
					fmt.Println("receiver go routine receive sender go routine notice chan is closed")
					return
				}
				fmt.Println("receiver go routine receive sender go routine notice value", v)
			}
		}
	}()

	go func() {
		totalIndex := 100
		exitIndex := rand.Intn(totalIndex)
		fmt.Println("sender go routine rand exit index", exitIndex)
		for index := 0; index != totalIndex; index++ {
			receiverGoroutineNoticeChan <- index
			time.Sleep(time.Millisecond * 100)
			if index == exitIndex {
				break
			}
		}
		receiverGoroutineExitChan <- struct{}{}
		close(receiverGoroutineNoticeChan)
		close(receiverGoroutineExitChan)
	}()

	<-mainGoroutineExitChan
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
