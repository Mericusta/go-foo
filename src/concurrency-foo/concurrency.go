package concurrencyfoo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type littleStruct struct {
	k int
	v int
}

const CHANNEL_BUFFER_SIZE = 1024

// goroutine 通信性能对比：buffer channel 发送大量小对象
func GoroutineCommunicateByBufferChannelWithLittleStructFoo(senderCount int) {
	c := make(chan *littleStruct, CHANNEL_BUFFER_SIZE)
	ctx, canceler := context.WithCancel(context.Background())

	go func(ctx context.Context, c <-chan *littleStruct) {
		for {
			select {
			case v, ok := <-c:
				if v == nil || !ok {
					return
				}
			case <-ctx.Done():
				for range c {
					fmt.Printf("receiver abandon\n")
					continue
				}
				fmt.Printf("receiver done\n")
				return
			}
		}
	}(ctx, c)

	wg := sync.WaitGroup{}
	wg.Add(senderCount)
	for index := 0; index != senderCount; index++ {
		go func(ctx context.Context, c chan<- *littleStruct, i int) {
			t := time.NewTicker(time.Millisecond * 10)
			for range t.C {
				select {
				case <-ctx.Done():
					t.Stop()
					fmt.Printf("index %v wg done\n", i)
					wg.Done()
					return
				default:
					fmt.Printf("index %v send\n", i)
					c <- &littleStruct{k: rand.Int(), v: rand.Int()}
				}
			}
		}(ctx, c, index)
	}

	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Printf("time's up\n")
	timer.Stop()
	fmt.Printf("canceler\n")
	canceler()
	fmt.Printf("wait\n")
	wg.Wait()
	fmt.Printf("close c\n")
	close(c)
}

// goroutine 通信性能对比：共享内存
func GoroutineCommunicateBySharedMemoryFoo() {
}

// goroutine 通信性能对比：buffer channel 尝试结合共享内存
func GoroutineCommunicateByBufferChannelAndSharedMemoryFoo() {
}
