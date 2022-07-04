package main

import (
	"fmt"
	mapfoo "go-foo/map-foo"
	"sync"
	"time"
)

func Bencher(count int, f func(int), concurrently bool) {
	var wg sync.WaitGroup
	if concurrently {
		wg = sync.WaitGroup{}
		wg.Add(count)
	}
	t1 := time.Now()
	for index := 0; index != count; index++ {
		if concurrently {
			go func(i int) {
				f(i)
				wg.Done()
			}(index)
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
	// outputFile, err := os.OpenFile("./resources/iofoo_WriteFileFoo.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0755)
	// if err != nil {
	// 	fmt.Printf("Open output file %v occurs error: %v\n", "./resources/iofoo_WriteFileFoo.log", err)
	// 	return
	// }
	// defer func() {
	// 	outputFile.Close()
	// }()

	// Bencher(3000, func(i int) {
	// 	t1 := time.Now()
	// 	iofoo.WriteFileFoo(i, outputFile)
	// 	t2 := time.Now()
	// 	fmt.Printf("writer %v/3000 Write File Foo using %v\n", i, t2.Sub(t1).Milliseconds())
	// }, true)

	mapfoo.StructMapKeyFoo()
}
