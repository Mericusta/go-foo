package builtinfoo

import (
	"fmt"
	"sync"
	"time"
)

type s struct {
	v int
}

// 传递对象指针并修改的表现
func GoroutinePassObjectPointerFoo(generatePointer bool) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan *s, 10)

	if generatePointer {
		sp := &s{v: 10}
		c <- sp
		sp = &s{v: 11}
		c <- sp
	} else {
		sp := s{v: 10}
		c <- &sp
		sp = s{v: 11}
		c <- &sp
	}

	time.Sleep(time.Second)

	go func(c chan *s) {
		count := 0
		for count != 2 {
			_sp := <-c
			fmt.Printf("_sp = %+v\n", _sp)
			count++
		}
		wg.Done()
	}(c)

	wg.Wait()
}
