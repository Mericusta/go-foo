package builtinfoo

import "fmt"

func Float32ZeroDivide() {
	var zero int
	var one int = 1
	fmt.Printf("float64: one/zero %v less than 10: %v\n", float64(one)/float64(zero), float64(one)/float64(zero) < float64(10))
	fmt.Printf("float64: zero/zero %v less than 10: %v\n", float64(zero)/float64(zero), float64(zero)/float64(zero) < float64(10))
}
