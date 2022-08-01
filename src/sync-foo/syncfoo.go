package syncfoo

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupCallFunctionWillCaptureWhenWait() {
	f := func(index int) {
		time.Sleep(time.Second)
		fmt.Printf("%v\n", index)
	}
	w := sync.WaitGroup{}

	// wrong: will capture index
	w.Add(5000)
	for index := 0; index != 5000; index++ {
		go func() {
			f(index)
			w.Done()
		}()
	}
	w.Wait()

	// correct: will capture index
	w.Add(5000)
	for index := 0; index != 5000; index++ {
		go func(i int) {
			f(i) // use closure func argument i but not caputre outer side value index
			w.Done()
		}(index) // pass index to closure func
	}
	w.Wait()
}
