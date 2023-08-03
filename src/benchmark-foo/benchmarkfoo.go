package benchmarkfoo

import (
	"sync"
	"time"
)

func producer(wg *sync.WaitGroup, p PoolDequeue, v any, produceCount int) {
	for i := 0; i < produceCount; i++ {
		p.PushHead(v)
	}
	wg.Done()
	// fmt.Printf("producer %v done\n", v)
}

func consumer(wg *sync.WaitGroup, p PoolDequeue, totalCount int) {
	i := 0
	for i < totalCount-1 {
		v, ok := p.PopTail()
		if !ok {
			time.Sleep(time.Millisecond)
			// fmt.Printf("consumer continue at %v\n", i)
			continue
		}
		// fmt.Printf("consumer receive value %v, i = %v\n", v, i)
		if v == nil {
			// fmt.Printf("consumer receive nil value at %v\n", i)
			continue
		}
		i++
		if i > totalCount {
			panic("overload")
		}
	}
	// fmt.Printf("consumer done\n")
	wg.Done()
}

func queueFoo(produceCount, producerCount int, p PoolDequeue) {
	if producerCount == 0 || produceCount == 0 {
		return
	}

	var (
		wg            = &sync.WaitGroup{}
		consumerCount = 1
		totalCount    = producerCount * produceCount
	)

	// consumer
	wg.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go consumer(wg, p, totalCount)
	}

	// producer
	wg.Add(producerCount)
	for i := 0; i < producerCount; i++ {
		go producer(wg, p, i+1, produceCount)
	}

	wg.Wait()
}
