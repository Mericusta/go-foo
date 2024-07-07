package concurrencyfoo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
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

// ----------------------------------------------------------------

// sync.Pool 切片测试
func SyncFoolWithSliceFoo() {
	defaultCapacity := 16
	p := &sync.Pool{
		New: func() any {
			return make([]byte, 0, defaultCapacity)
		},
	}

	sInfo := func(s []byte) {
		fmt.Printf("s %v, ptr %p, len %v, cap %v\n", s, s, len(s), cap(s))
	}

	anyS := p.Get()
	s := anyS.([]byte)
	sInfo(s)
	s = s[0:0] // reset slice
	sInfo(s)
	for index, c := 0, defaultCapacity; index != c; index++ {
		s = append(s, 1)
	}
	sInfo(s)
	p.Put(s)
	fmt.Println()

	anyS = p.Get()
	s = anyS.([]byte)
	sInfo(s)
	s = s[0:0]
	sInfo(s)
	for index, c := 0, defaultCapacity*2; index != c; index++ {
		s = append(s, 2)
	}
	sInfo(s)
	p.Put(s)
	fmt.Println()

	anyS = p.Get()
	s = anyS.([]byte)
	sInfo(s)
	s = s[0:0]
	sInfo(s)
	for index, c := 0, defaultCapacity; index != c; index++ {
		s = append(s, 3)
	}
	sInfo(s)
	p.Put(s)
	fmt.Println()

	anyS = p.Get() // old one in pool
	s1 := anyS.([]byte)
	anyS = p.Get() // new
	s2 := anyS.([]byte)
	anyS = p.Get() // new
	s3 := anyS.([]byte)

	sInfo(s1)
	sInfo(s2)
	sInfo(s3)

	p.Put(s1)
	p.Put(s2)
	p.Put(s3)

	// go func(p *sync.Pool) {
	// }(p)
}

// 原子操作测试
func atomicFoo(withAtomicWrite bool) {
	value := int64(1)
	// 又读又写协程
	rwFunc := func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			// fmt.Printf("rwFunc, read %v\n", value)
			if withAtomicWrite {
				atomic.AddInt64(&value, 1)
			} else {
				value++
			}
			fmt.Printf("rwFunc, write %v\n", value)
		}
	}

	// 高频读取协程
	rFunc := func() {
		ticker := time.NewTicker(time.Millisecond)
		for range ticker.C {
			fmt.Printf("rFunc, read %v\n", atomic.LoadInt64(&value))
		}
	}

	go rwFunc()
	go rFunc()

	timer := time.NewTimer(time.Second * 5)
	<-timer.C
}

const dequeueBits = 32

func unpack(ptrs uint64) (uint32, uint32) {
	const mask = 1<<dequeueBits - 1
	return uint32((ptrs >> dequeueBits) & mask), uint32(ptrs & mask)
}

func pack(slotCursor, handlerSlicePtr uint32) uint64 {
	const mask = 1<<dequeueBits - 1
	return (uint64(slotCursor) << dequeueBits) | uint64(handlerSlicePtr&mask)
}

func lockFreeDoubleVectorSlice() {
	const (
		// 槽数量
		slotCount = 60
		// 槽游标
		slotCursor = 0
		// 生产者数量
		producerCount = 1 << 0
		// 消费者数量
		consumerCount = 1
	)

	// 槽和列表，60个槽，每个槽都有一个 slice
	slots := [slotCount][]int{}

	for index := 0; index != slotCount; index++ {
		slots[index] = make([]int, 0, 64)
	}

	// 槽游标，列表聚合指针
	var cursorHandlerPtr atomic.Uint64
	// TODO: uintptr 是 uint64，转换成 uint32 会被截断
	cursorHandlerPtr.Store(pack(slotCursor, uint64(uintptr(unsafe.Pointer(&slots[0])))))

	// 生产者
	producer := func() {

	}

	// 消费者
	consumer := func() {

	}
}
