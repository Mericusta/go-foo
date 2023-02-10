package gcfoo

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

func ForceGCPointerSlice(c int) {
	s := make([]*int, c)
	for i := 0; i < 10; i++ {
		t := time.Now()
		runtime.GC()
		fmt.Printf("the slice has %v number of *int elements, No.%v GC using time %s\n", c, i, time.Since(t))
	}
	runtime.KeepAlive(s)
}

func ForceGCNonPointerSlice(c int) {
	s := make([]int, c)
	for i := 0; i < 10; i++ {
		t := time.Now()
		runtime.GC()
		fmt.Printf("the slice has %v number of int elements, No.%v GC using time %s\n", c, i, time.Since(t))
	}
	runtime.KeepAlive(s)
}

func ForceGCPointerSliceInOSHeap(c int) {
	makeSliceFromOSHeap := func(sliceLength int, elementSize uintptr) reflect.SliceHeader {
		fd := -1
		dataPtr, _, errorNo := syscall.Syscall6(
			syscall.SYS_MMAP,
			0, // address
			uintptr(sliceLength)*elementSize,
			syscall.PROT_READ|syscall.PROT_WRITE,
			syscall.MAP_ANON|syscall.MAP_PRIVATE,
			uintptr(fd), // no file descriptor
			0,           // offset
		)
		if errorNo != 0 {
			panic(errorNo)
		}
		return reflect.SliceHeader{
			Data: dataPtr,
			Len:  sliceLength,
			Cap:  sliceLength,
		}
	}

	var intPtr *int
	sliceHeader := makeSliceFromOSHeap(c, unsafe.Sizeof(intPtr))
	s := *(*[]*int)(unsafe.Pointer(&sliceHeader))

	for i := 0; i < 10; i++ {
		t := time.Now()
		runtime.GC()
		fmt.Printf("the slice has %v number of int elements, No.%v GC using time %s\n", c, i, time.Since(t))
	}
	runtime.KeepAlive(s)
}

func ForceGCNoNeedReleaseString(c int) {
	ss := make([]string, c) // no need release string slice

	// 'string' type is 'reflect.StringHeader'
	// 'reflect.StringHeader' has pointer member 'Data uintptr'
	// GC will scan all pointer in ss

	for i := 0; i < c; i++ {
		s := strconv.Itoa(i)
		ss = append(ss, s)
	}

	for i := 0; i < 10; i++ {
		t := time.Now()
		runtime.GC()
		fmt.Printf("the slice has %v number of string elements, No.%v GC using time %s\n", c, i, time.Since(t))
	}
	runtime.KeepAlive(ss)
}

func AvoidForceGCNoNeedReleaseString(c int) {
	var stringBytes []byte
	var stringOffsets []int

	for i := 1; i <= c; i++ {
		s := strconv.Itoa(i)
		stringBytes = append(stringBytes, s...)
		stringOffsets = append(stringOffsets, int(math.Log(float64(i)))+1)
	}

	for i := 0; i < 10; i++ {
		runtime.GC()
		t := time.Now()
		runtime.GC()
		fmt.Printf("the slice has %v number of string elements, No.%v GC using time %s\n", c, i, time.Since(t))
	}

	// get data
	sStart := 0
	for i := 0; i < 10; i++ {
		sEnd := stringOffsets[i]
		bytes := stringBytes[sStart:sEnd]
		s := *(*string)(unsafe.Pointer(&bytes))
		fmt.Printf("bytes [%v:%v] is %v\n", sStart, sEnd, s)
		sStart = sEnd
	}
}
