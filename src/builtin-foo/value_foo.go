package builtinfoo

import "fmt"

func Float32ZeroDivide() {
	var zero int
	var one int = 1
	fmt.Printf("float64: one/zero %v less than 10: %v\n", float64(one)/float64(zero), float64(one)/float64(zero) < float64(10))
	fmt.Printf("float64: zero/zero %v less than 10: %v\n", float64(zero)/float64(zero), float64(zero)/float64(zero) < float64(10))
}

func BitOp1() {
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

func BitOp2() {
	const (
		bits = 8
		mask = 1<<bits - 1
	)
	var (
		v  uint16
		v1 uint8 = 15
		v2 uint8 = 13
		a  []int = make([]int, (1<<bits)/2)
	)
	fmt.Printf("a len = %v\n", len(a))
	fmt.Printf("v1 v2 max = %v, %b\n", 1<<8-1, 1<<8-1)

	// pack
	fmt.Printf("v1 = %b\n", v1)
	fmt.Printf("v2 = %b\n", v2)
	vHigh := uint16(v1) << bits
	vLow := uint16(v2 & mask)
	vLowUnmask := uint16(v2)
	v = vHigh | vLow
	vUnmask := vHigh | vLow
	fmt.Printf("vHigh = %b\n", vHigh)
	fmt.Printf("vLow = %b\n", vLow)
	fmt.Printf("vLowUnmask = %b\n", vLowUnmask)
	fmt.Printf("v = %b\n", v)
	fmt.Printf("vUnmask = %b\n", vUnmask)

	// unpack
	high := (v >> bits) & mask
	highUnmask := (v >> bits)
	low := uint8(v & mask)
	lowUnmask := uint8(v)
	fmt.Printf("high = %b\n", high)
	fmt.Printf("low = %b\n", low)
	fmt.Printf("highUnmask = %b\n", highUnmask)
	fmt.Printf("lowUnmask = %b\n", lowUnmask)

	// push head
	fmt.Printf("full check %b\n", (low+uint8(len(a)))&mask)
}
