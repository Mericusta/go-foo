package builtinfoo

import "fmt"

func Float32ZeroDivide() {
	var zero int
	var one int = 1
	fmt.Printf("float64: one/zero %v less than 10: %v\n", float64(one)/float64(zero), float64(one)/float64(zero) < float64(10))
	fmt.Printf("float64: zero/zero %v less than 10: %v\n", float64(zero)/float64(zero), float64(zero)/float64(zero) < float64(10))
}

func BitOp() {
	var (
		is      []int = []int{1, -1}
		total64 uint64
	)

	for _, v := range is {
		left32 := uint64(v) << 32
		total64 += left32
		high32 := total64 >> 32
		low32 := uint32(total64)
		fmt.Printf("v = %v, left32 = %v\n", v, left32)
		fmt.Printf("high32 = %v, low32 = %v\n", high32, low32)
		fmt.Printf("total64 = %v\n", total64)
	}
}
