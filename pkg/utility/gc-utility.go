package utility

import (
	"runtime"
	"time"
)

func ForceGC(n, c int) {
	totalDuration := time.Duration(0)
	for i := 0; i != c; i++ {
		t := time.Now()
		runtime.GC()
		d := time.Since(t)
		totalDuration += d
		// fmt.Printf("number of elements %v, No.%v GC using time %s\n", n, i, d)
	}
	// fmt.Printf("number of elements %v, average GC using time %s\n", n, totalDuration/10)
}
