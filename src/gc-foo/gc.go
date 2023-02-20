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

func forceGC(n, c int) {
	totalDuration := time.Duration(0)
	for i := 0; i != c; i++ {
		t := time.Now()
		runtime.GC()
		d := time.Since(t)
		totalDuration += d
		fmt.Printf("number of int elements %v, No.%v GC using time %s\n", n, i, d)
	}
	fmt.Printf("number of int elements %v, average GC using time %s\n", n, totalDuration/10)
}

func ForceGCPointerSlice(c int) {
	s := make([]*int, c)
	forceGC(c, 10)
	runtime.KeepAlive(s)
}

func ForceGCNonPointerSlice(c int) {
	s := make([]int, c)
	forceGC(c, 10)
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

	forceGC(c, 10)
	runtime.KeepAlive(s)
}

func ForceGCNoNeedReleaseStringSlice(c int) {
	ss := make([]string, c) // no need release string slice

	// 'string' type is 'reflect.StringHeader'
	// 'reflect.StringHeader' has pointer member 'Data uintptr'
	// GC will scan all pointer in ss

	for i := 0; i < c; i++ {
		s := strconv.Itoa(i)
		ss = append(ss, s)
	}

	forceGC(c, 10)
	runtime.KeepAlive(ss)
}

func AvoidForceGCNoNeedReleaseStringSlice(c int) {
	var stringBytes []byte
	var stringOffsets []int

	for i := 1; i <= c; i++ {
		s := strconv.Itoa(i)
		stringBytes = append(stringBytes, s...)
		stringOffsets = append(stringOffsets, int(math.Log(float64(i)))+1)
	}

	forceGC(c, 10)

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

func ForceGCNoNeedReleaseStringMap(c int) {
	sm := make(map[string]int)

	for i := 0; i < c; i++ {
		s := strconv.Itoa(i)
		sm[s] = i
	}

	forceGC(c, 10)
	runtime.KeepAlive(sm)
}

func AvoidForceGCNoNeedReleaseStringMap(c int) {
	sm := make(map[int]int)

	for i := 0; i < c; i++ {
		sm[i] = i
	}

	forceGC(c, 10)
	runtime.KeepAlive(sm)
}

type gcStruct struct {
	i int
	v int
}

func ForceGCStructPointerMap(c int) {
	m := make(map[int]*gcStruct)
	for i := 0; i < c; i++ {
		m[i] = &gcStruct{
			i: i,
			v: i,
		}
	}

	forceGC(c, 10)
	runtime.KeepAlive(m)
}

func ForceGCStructPointerSlice(c int) {
	s := make([]*gcStruct, 0, c)

	for i := 0; i < c; i++ {
		s = append(s, &gcStruct{
			i: i,
			v: i,
		})
	}

	forceGC(c, 10)
	runtime.KeepAlive(s)
}

func ForceGCByteSlice(c int) {
	s := make([][]byte, 0, c)

	for i := 0; i < c; i++ {
		s = append(s, make([]byte, 0, 24))
	}

	forceGC(c, 10)
	runtime.KeepAlive(s)
}

func ForceGCByteSliceMap(c int) {
	m := make(map[int][]byte)

	for i := 0; i < c; i++ {
		m[i] = make([]byte, 0, 24)
	}

	forceGC(c, 10)
	runtime.KeepAlive(m)
}

type gcStruct2 struct {
	i1 int
	i2 int
	i3 int
	i4 int
	i5 int
	i6 int
}

func GCScanCompare(c int) {
	size := unsafe.Sizeof(gcStruct2{})
	fmt.Printf("gcStruct2 size = %v\n", size)

	s := make([]*gcStruct2, 0, c)
	for i := 0; i < c; i++ {
		sp := &gcStruct2{
			i1: i + 1, i2: i + 2, i3: i + 3,
			i4: i + 4, i5: i + 5, i6: i + 6,
		}

		s = append(s, sp)
	}

	// uintptr lost object memory reference
	// but []byte catch object memory reference
	// it will not release by GC
	forceGC(c, 10)
	runtime.KeepAlive(s)
	time.Sleep(time.Second)

	lost := 0
	for i := 0; i < c; i++ {
		sp := s[i]
		if sp.i1 != i+1 || sp.i2 != i+2 || sp.i3 != i+3 {
			lost++
			fmt.Printf("%v not equal, sp = %+v, s[i] = %v\n", i, sp, s[i])
		}
	}
	fmt.Printf("1, lost count = %v\n", lost)
}

func AvoidGCScanByUintptr(c int) {
	fmt.Printf("gcStruct2 size = %v\n", unsafe.Sizeof(gcStruct2{}))

	s := make([]uintptr, 0, c)

	var sp *gcStruct2
	for i := 0; i < c; i++ {
		sp = &gcStruct2{
			i1: i + 1, i2: i + 2, i3: i + 3,
			i4: i + 4, i5: i + 5, i6: i + 6,
		}
		s = append(s, uintptr(unsafe.Pointer(sp)))
	}

	// uintptr lost object memory reference
	// the object will be released by GC
	forceGC(c, 10)
	runtime.KeepAlive(s)
	time.Sleep(time.Second)

	lost := 0
	for i := 0; i < c; i++ {
		sp := (*gcStruct2)(unsafe.Pointer(s[i]))
		if sp.i1 != i+1 || sp.i2 != i+2 || sp.i3 != i+3 {
			lost++
			fmt.Printf("%v not equal, sp = %+v, s[i] = %v\n", i, sp, s[i])
		}
	}
	fmt.Printf("1, lost count = %v\n", lost)
}

func AvoidGCScanByByteSlice(c int) {
	size := unsafe.Sizeof(gcStruct2{})
	fmt.Printf("gcStruct2 size = %v\n", size)

	s := make([][]byte, 0, c)

	var sp *gcStruct2
	type _s struct {
		p uintptr
		l int
		c int
	}
	for i := 0; i < c; i++ {
		sp = &gcStruct2{
			i1: i + 1, i2: i + 2, i3: i + 3,
			i4: i + 4, i5: i + 5, i6: i + 6,
		}

		_s := &_s{p: uintptr(unsafe.Pointer(sp)), l: int(size), c: int(size)}
		s = append(s, *(*[]byte)(unsafe.Pointer(_s)))
	}

	// uintptr lost object memory reference
	// but []byte catch object memory reference
	// it will not release by GC
	forceGC(c, 10)
	runtime.KeepAlive(s)
	time.Sleep(time.Second)

	lost := 0
	for i := 0; i < c; i++ {
		sp := *(**gcStruct2)(unsafe.Pointer(&s[i]))
		if sp.i1 != i+1 || sp.i2 != i+2 || sp.i3 != i+3 {
			lost++
			fmt.Printf("%v not equal, sp = %+v, s[i] = %v\n", i, sp, s[i])
		}
	}
	fmt.Printf("1, lost count = %v\n", lost)
}
