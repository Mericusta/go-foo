package goroutinefoo

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func OpenSoMuchGoRoutine() {
	c1, c2 := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(math.MaxInt16 * 8)
	for index := 0; index != math.MaxInt16*8; index++ {
		go func() {
			// block here
			<-c1
			c2 <- true
			wg.Done()
		}()
	}
	fmt.Printf("start sleep 1 minute, it will alloc 2GB memory\n")
	fmt.Printf("but if you terminate the process, it will release\n")
	t := time.NewTimer(time.Minute)
	<-t.C
	c1 <- true
	wg.Wait()
}
