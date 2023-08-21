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

func BitOp3() {
	var head, tail, mask uint8 = 5, 5, 1<<8 - 1
	fmt.Printf("head = %v, %b\n", head, head)
	fmt.Printf("tail = %v, %b\n", tail, tail)
	for len := 0; len < 8; len++ {
		var _v uint8 = uint8(len - 1)
		slotIndex := head & _v
		fullChecker := tail + uint8(len)&mask
		fmt.Printf("len = %v\n", len)
		fmt.Printf("- head = %v %b\n", head, head)
		fmt.Printf("- tail = %v %b\n", tail, tail)
		fmt.Printf("- fullChecker = %v, %b\n", fullChecker, fullChecker)
		fmt.Printf("- _v = %v, %b\n", _v, _v)
		fmt.Printf("- slotIndex = %v, %b\n", slotIndex, slotIndex)
		fmt.Println()
	}
}

func Mod() {
	r, i1, i2 := 5, 1, 3
	fmt.Printf("%v mod %v = %v\n", i1, r, i1%r)
	fmt.Printf("%v mod %v = %v\n", i2, r, i2%r)
	r, i1, i2 = 5, 3, 1
	fmt.Printf("%v mod %v = %v\n", i1, r, i1%r)
	fmt.Printf("%v mod %v = %v\n", i2, r, i2%r)
}
