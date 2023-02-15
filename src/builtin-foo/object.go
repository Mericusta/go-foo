package builtinfoo

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	i   int   // 可变长度 -> 64 位机器上是8位，对齐长度为8，第1次对齐，占8位
	i8  int8  // 长度1位，第2次对齐，占2位，第2次对齐内再次对齐
	i32 int32 // 长度4位，第2次对齐，占4位，第2次对齐内再次对齐
	i16 int16 // 长度2位，第2次对齐，占2位，第2次对齐内再次对齐
	i64 int64 // 长度8位，第3次对齐，占8位
	// f32 float32
	// f64 float64
	// s   string          // 可变长度
	// sub *convertStruct2 // 可变长度
}

func convertType2() (bool, []byte) {
	o := &convertStruct2{
		i:   1,
		i8:  2,
		i16: 3,
		i32: 4,
		// i64: 5,
		// f32: 3.1415,
		// f64: 2.7183,
		// s:   "convert type 2",
		// sub: &convertStruct2{s: "convert type 2 sub"},
	}
	fmt.Printf("Sizeof = %v\n", unsafe.Sizeof(*o))
	fmt.Printf("Alignof = %v\n", unsafe.Alignof(*o))
	b := convertObjectToByteArray(o)
	fmt.Printf("b len = %v\n", len(b))
	fmt.Printf("b = %v\n", b)
	// fmt.Printf("i = %v\n", b[0:8])
	// fmt.Printf("i8 = %v\n", b[8:9])
	// fmt.Printf("i16 = %v\n", b[9:11])
	// fmt.Printf("i32 = %v\n", b[11:15])
	// fmt.Printf("i64 = %v\n", b[15:23])
	// fmt.Printf("f32 = %v\n", b[23:31])
	// fmt.Printf("f64 = %v\n", b[31:39])
	// fmt.Printf("s = %v\n", b[39:])
	ro := convertByteArrayToObject(b)
	return reflect.DeepEqual(o, ro), b
}

type _s struct {
	p uintptr
	l int
	c int
}

func convertObjectToByteArray(o *convertStruct2) []byte {
	l := unsafe.Sizeof(*o) // 注意这里必须是具体类型
	s := &_s{p: uintptr(unsafe.Pointer(o)), l: int(l), c: int(l)}
	b := *(*[]byte)(unsafe.Pointer(s))
	return b
}

func convertByteArrayToObject(b []byte) *convertStruct2 {
	return *(**convertStruct2)(unsafe.Pointer(&b))
}

// 方式3：
// switch o.(type) 递归自行实现
// 优点：
// 缺点：
