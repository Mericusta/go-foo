package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Printf("threads = %v\n", threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	panic(x)
}
