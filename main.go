package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Bencher(count int, f func(int), concurrently bool) {
	var wg sync.WaitGroup
	if concurrently {
		wg = sync.WaitGroup{}
		wg.Add(count)
	}
	t1 := time.Now()
	for index := 0; index != count; index++ {
		if concurrently {
			go func() {
				f(index)
				wg.Done()
			}()
		} else {
			f(index)
		}
	}
	if concurrently {
		wg.Wait()
	}
	t2 := time.Now()
	fmt.Printf("using time: %v milli-seconds\n", t2.Sub(t1).Milliseconds())
}

func main() {
	Bencher(1, func(index int) {
		// algorithmfoo.CalculateYearsOldTest()
		// jsonfoo.JsonFoo()
		// functionfoo.ReturnExampleStructTest()
		// mysqlfoo.BatchInsertPrayRecordData()
		// algorithmfoo.OptimusTest()
		// fmt.Printf("uint64MAX := ^uint64(0) = %v\n", ^uint64(0)) // 18446744073709551615
		// regexpfoo.AllRegexpTest()
		// extractorfoo.ExtractGoFileInterfaceDeclarationTest()
		// extractorfoo.ExtractGoFileFunctionScopeTest()
		// slicefoo.PassSliceAndChangeIt()
		// arrayfoo.ClearArrayFoo()
		// arrayfoo.ReturnArrayBeforeIndexFoo()
		// benchmarkfoo.Pray()

		c := make(chan int)
		go func() {
			c <- 1
		}()
		go func() {
			time.Sleep(time.Second)
			<-c
		}()
	}, false)
}
