package builtinfoo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go-foo/pkg/utility"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type s struct {
	v int
}

// 传递对象指针并修改的表现
func GoroutinePassObjectPointerFoo(generatePointer bool) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan *s, 10)

	if generatePointer {
		sp := &s{v: 10}
		c <- sp
		sp = &s{v: 11}
		c <- sp
	} else {
		sp := s{v: 10}
		c <- &sp
		sp = s{v: 11}
		c <- &sp
	}

	time.Sleep(time.Second)

	go func(c chan *s) {
		count := 0
		for count != 2 {
			_sp := <-c
			fmt.Printf("_sp = %+v\n", _sp)
			count++
		}
		wg.Done()
	}(c)

	wg.Wait()
}

// in 64 word size platform:
// ┌────────────────────────────────────┬────────┐
// │type                                │size    │
// ├────────────────────────────────────┼────────┤
// │bool                                │1 byte  │
// ├────────────────────────────────────┼────────┤
// │intN, uintN, floatN, complexN       │N/8 byte│
// ├────────────────────────────────────┼────────┤
// │int, uint, uintptr                  │8 byte  │
// ├────────────────────────────────────┼────────┤
// │*T, map, func, chan                 │8 byte  │
// ├────────────────────────────────────┼────────┤
// │string (data, len)                  │16 byte │
// ├────────────────────────────────────┼────────┤
// │interface (tab, data or _type, data)│16 byte │
// ├────────────────────────────────────┼────────┤
// │[]T (array, len, cap)               │24 byte │
// └────────────────────────────────────┴────────┘

// 编译器结构体内存对齐规则
// 规则一
// - 偏移量是结构体成员内存地址相对结构体起始地址的差值
// - 结构体第一个成员变量偏移量为0
// - 后面的成员变量偏移量等于 成员变量的大小(unsafe.Sizeof()) 和 成员类型在编译器的默认对齐系数(unsafe.Alignof()) 两者中较小的那个值的最小整数倍
// - 如果不满足规则，编译器会在前面填充值为0的字节空间
// 规则二
// - 结构体本身也需要内存对齐
// - 其大小等于各成员变量占用内存最大的和编译器默认对齐系数两者中较小的那个值的最小整数倍

type structMemberSizeAlignTypeDesc struct {
	size  int
	align int
	desc  string
}

func StructMemoryAlignCalculateProcess(compilerDefaultAlign int, smDesc []*structMemberSizeAlignTypeDesc) (int, int, []int, string) {
	memberLen := len(smDesc)
	memberAllocation := make([]int, memberLen)
	b := strings.Builder{}
	allocation := 0
	wasting := 0
	maxMemberSize := 0

	// allocate begin
	b.WriteRune('[')
	b.WriteRune(' ')

	// according to rule 1
	for i, sm := range smDesc {
		fmt.Printf("struct No.%v member, size %v, align %v, type %v\n", i+1, sm.size, sm.align, sm.desc)

		min := int(math.Min(float64(sm.size), float64(sm.align)))
		fmt.Printf("No.%v member, min(size, align) = %v\n", i, min)
		if allocation%min == 0 {
			fmt.Printf("allocate offset = %v, offset mod min(size, align) == 0, no need fill\n", allocation)
		} else {
			fillByte := min - allocation
			if allocation > min {
				v := min
				for v <= allocation {
					v += min
				}
				fillByte = v - allocation
			}
			fmt.Printf("allocate offset = %v, offset mod min(size, align) != 0, need fill %v byte\n", allocation, fillByte)
			wasting += fillByte
			// allocate fill byte
			allocation += fillByte
			memberAllocation[i] += fillByte
			for i := 0; i < fillByte; i++ {
				b.WriteRune('_')
				b.WriteRune(' ')
			}
		}

		// allocate member byte
		allocation += sm.size
		memberAllocation[i] += sm.size
		for i := 0; i < sm.size; i++ {
			b.WriteRune('0')
			b.WriteRune(' ')
		}
		if maxMemberSize < sm.size {
			maxMemberSize = sm.size
		}

		// allocate end
		if i != memberLen-1 {
			b.WriteRune('|')
			b.WriteRune(' ')
		}

		// output allocation
		fmt.Printf("struct No.%v member memory allocation: %v\n", i+1, b.String())
	}
	fmt.Printf("according to rule 1, struct member memory allocation size: %v\n", allocation)
	fmt.Printf("allocation: %v\n", b.String())

	// according to rule 2
	fmt.Printf("struct max member size = %v\n", maxMemberSize)
	fmt.Printf("compiler default align = %v\n", compilerDefaultAlign)
	min := int(math.Min(float64(maxMemberSize), float64(compilerDefaultAlign)))
	fmt.Printf("min(maxMemberSize, compilerDefaultAlign) = %v\n", min)
	if allocation%min == 0 {
		fmt.Printf("struct no need fill\n")
	} else {
		fillByte := min - allocation
		if allocation > min {
			v := min
			for v <= allocation {
				v += min
			}
			fillByte = v - allocation
		}
		fmt.Printf("struct need fill %v byte\n", fillByte)
		wasting += fillByte
		// allocate fill byte
		allocation += fillByte
		b.WriteRune('|')
		b.WriteRune(' ')
		for i := 0; i < fillByte; i++ {
			b.WriteRune('_')
			b.WriteRune(' ')
		}
	}

	// allocate end
	b.WriteRune(']')

	fmt.Printf("according to rule 2, struct member memory allocation size: %v\n", allocation)
	fmt.Printf("allocation: %v\n", b.String())
	fmt.Printf("wasting offset: %v, memory utilization %.2f%%\n", wasting, float64(allocation-wasting)/float64(allocation)*100)

	return allocation, wasting, memberAllocation, b.String()
}

type es struct {
	i int8
	s int8
	f int8
}

func (e *es) Encode() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, e)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// 任意对象转换为 []byte 数组
func ConvertAnyObjectToByteArray(o *es) []byte {
	return o.Encode()
}

// 方式1：
// binary 将所有固定长度的类型或只包含固定长度的 struct 转换为 []byte
// gob 将所有非固定长度的类型或 struct 转换为 []byte
// 优点：现有实现组合即可
// 缺点：

// 方式2：
// 模拟 slice 结构
// 按照头指针+长度来扫结构体内存
// 优点：支持可变长度的类型
// 缺点：必须明确类型才能计算出长度（可以用代码处理？）

type convertStruct2 struct {
	i int
	// i8  int8
	// i16 int16
	// i32 int32
	// i64 int64
	// f32 float32
	// f64 float64
	s string
	// sub *convertStruct2
}

var dataSlice [][]byte
var structSlice []*convertStruct2

func setDataSlice(o *convertStruct2) int {
	// fmt.Printf("setDataSlice p ptr = %v\n", uintptr(unsafe.Pointer(o)))
	// o.s = fmt.Sprintf("convert type %v example", len(dataSlice))
	b := convertObjectToByteArray(o)
	dataSlice = append(dataSlice, b)
	// fmt.Printf("setDataSlice data[0] ptr %v\n", uintptr(unsafe.Pointer(&dataSlice[0])))
	// structSlice = append(structSlice, o)
	return len(dataSlice)
}

func getDataSlice(i int) *convertStruct2 {
	if i > len(dataSlice) {
		return nil
	}
	b := dataSlice[i]
	if len(b) == 0 {
		return nil
	}
	return convertByteArrayToObject(b)
}

func convertStruct2Example(c int) {
	dataSlice = make([][]byte, 0, c)

	var o *convertStruct2
	for i := 0; i != c; i++ {
		o = &convertStruct2{i: i, s: fmt.Sprintf("convert type %v example", i)}
		setDataSlice(o)
	}

	// update
	for i := 0; i != c; i++ {
		o := getDataSlice(i)
		co := &convertStruct2{i: i, s: fmt.Sprintf("convert type %v example", i)}
		if !reflect.DeepEqual(o, co) {
			panic(fmt.Sprintf("%+v != %+v", o, co))
		}
	}
}

func convertType2() (bool, []byte) {
	o := &convertStruct2{
		// i:   1,
		// i8:  2,
		// i16: 3,
		// i32: 4,
		// i64: 5,
		// f32: 3.1415,
		// f64: 2.7183,
		s: "convert type 2",
		// sub: &convertStruct2{s: "convert type 2 sub"},
	}
	fmt.Printf("Sizeof = %v\n", unsafe.Sizeof(*o))
	fmt.Printf("Alignof = %v\n", unsafe.Alignof(*o))
	// fmt.Printf("ptr = %p\n", o) // _s.a 的类型是 uintptr 不执行该句会导致返回的 b 结果不一致？
	// b := []byte{47, 195, 90, 0, 0, 0, 0, 0, 14, 0, 0, 0, 0, 0, 0, 0}
	b := convertObjectToByteArray(o)
	// fmt.Printf("b ptr = %p\n", b)
	// fmt.Printf("o ptr = %p\n", o)

	// fmt.Printf("b len = %v\n", len(b))
	// fmt.Printf("b = %v, ptr = %p\n", b, b)

	// b1 := make([]byte, len(b))
	// copy(b1, b)
	// fmt.Printf("b1 = %v, ptr = %p\n", b1, b1)

	// o = nil
	// runtime.GC()

	// b2 := make([]byte, len(b))
	// copy(b2, b)
	// fmt.Printf("b2 = %v, ptr = %p\n", b2, b2)

	// ro := convertByteArrayToObject(b)
	// fmt.Printf("ro = %+v\n", ro)
	// fmt.Printf("ro = %+v\n", convertByteArrayToObject(b))

	return false, b
}

type _slice struct {
	p uintptr // lost object memory reference
	l int
	c int
}

func convertObjectToByteArray(o *convertStruct2) []byte {
	l := unsafe.Sizeof(*o) // 注意这里必须是具体类型
	// 用 _s.a 指向了这个对象的内存空间，以避免这个对象被 GC
	s := &_slice{p: uintptr(unsafe.Pointer(o)), l: int(l), c: int(l)}
	// 然后将 _s 这个结构转化成 []byte，可以转换的原因如下：
	// 1 struct 中所有成员变量的内存是连续分布的
	// 2 _s 结构跟 []type 的底层结构是一样的
	b := *(*[]byte)(unsafe.Pointer(s))
	// fmt.Println("after convertObjectToByteArray")
	// fmt.Printf("s ptr = %p\n", s)                                          // 指向 s 本身
	// fmt.Printf("s.a = %v\n", s.a)                                          // 指向 转换的对象
	// fmt.Printf("b ptr = %p\n", b)                                          // 指向 b 的第一个元素
	// fmt.Printf("s.a == &b[0] %v\n", s.a == uintptr(unsafe.Pointer(&b[0]))) // 等于 转换的对象
	return b
}

func TConvertObjectToByteArray[T any](o *T) []byte {
	l := unsafe.Sizeof(*o) // 注意这里必须是具体类型
	s := &_slice{p: uintptr(unsafe.Pointer(o)), l: int(l), c: int(l)}
	// 然后将 _s 这个结构转化成 []byte，可以转换的原因如下：
	// 1 struct 中所有成员变量的内存是连续分布的
	// 2 _s 结构跟 []type 的底层结构是相似的
	// 不会被 GC 的原因如下：
	// []byte 的底层结构 runtime.slice 中会将指针转换为 unsafe.Pointer
	b := *(*[]byte)(unsafe.Pointer(s))
	// fmt.Println("after convertObjectToByteArray")
	// fmt.Printf("s ptr = %p\n", s)                                          // 指向 s 本身
	// fmt.Printf("s.a = %v\n", s.a)                                          // 指向 转换的对象
	// fmt.Printf("b ptr = %p\n", b)                                          // 指向 b 的第一个元素
	// fmt.Printf("s.a == &b[0] %v\n", s.a == uintptr(unsafe.Pointer(&b[0]))) // 等于 转换的对象
	return b
}

func convertByteArrayToObject(b []byte) *convertStruct2 {
	return *(**convertStruct2)(unsafe.Pointer(&b))
}

func TConvertByteArrayToObject[T any](b []byte) *T {
	return *(**T)(unsafe.Pointer(&b))
}

// 方式3：
// switch o.(type) 递归自行实现
// 优点：
// 缺点：

// ----

type _string struct {
	p uintptr
	l int
}

type stringStruct struct {
	s1 string
	s2 string
	s3 string
}

func ConvertStringToStringStruct0(s string) *stringStruct {
	sStruct := new(stringStruct)
	sSlice := strings.Split(s, ",")
	l := len(sSlice)
	i := 0
	if l > i {
		sStruct.s1 = sSlice[i]
		i++
	}
	if l > i {
		sStruct.s2 = sSlice[i]
		i++
	}
	if l > i {
		sStruct.s3 = sSlice[i]
		i++
	}

	utility.ForceGC(l, 10)

	return sStruct
}

func ConvertStringToStringStruct1(s string) *stringStruct {
	sSlice := strings.Split(s, ",")
	offset := unsafe.Sizeof(s) // 16
	len := len(sSlice)

	sStruct := new(stringStruct)
	ptr := uintptr(unsafe.Pointer(sStruct))
	for index := uintptr(0); uintptr(len) > index; index++ {
		*(*string)(unsafe.Pointer(ptr + offset*index)) = sSlice[index]
	}

	utility.ForceGC(len, 10)

	return sStruct
}

func ConvertStringToStringStruct2(s string) *stringStruct {
	sSlice := strings.Split(s, ",")
	offset := unsafe.Sizeof(s) // 16
	count := len(sSlice)

	b := make([]byte, int(offset)*count)

	sPtr := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	sLen := len(s)
	subIndex, subLen := 0, 0
	for i := 0; i < sLen; i++ {
		r := s[i]
		if r == ',' || i == sLen-1 {
			if i == sLen-1 {
				sPtr++
				subLen++
			}
			sb := TConvertObjectToByteArray(&reflect.StringHeader{Data: sPtr - uintptr(subLen), Len: subLen})
			copy(b[subIndex*int(offset):subIndex*int(offset)+int(offset)], sb)
			subIndex++
			subLen = 0
		} else {
			subLen++
		}
		sPtr++
	}

	utility.ForceGC(sLen, 10)

	return *(**stringStruct)(unsafe.Pointer(&b))
}
