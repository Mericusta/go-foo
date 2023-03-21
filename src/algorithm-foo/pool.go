package algorithmfoo

import (
	"errors"
	"fmt"
	"go-foo/pkg/utility"
	"math"
	"math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/coocood/freecache"
)

func TConvertObjectToByteArray[T any](o *T) []byte {
	l := unsafe.Sizeof(*o)
	s := &reflect.SliceHeader{Data: uintptr(unsafe.Pointer(o)), Len: int(l), Cap: int(l)}
	b := *(*[]byte)(unsafe.Pointer(s))
	return b
}

// 内存反转函数
func tConvertByteArrayToObject[T any](b []byte) T {
	return *(*T)(unsafe.Pointer(&b))
}

// 内存池全局持有对象，避免内存池中的数据被 GC-SCAN 以及内存空间被 GC
var poolByte []byte

// 获得内存
func getMemory(i, l int) []byte {
	return poolByte[i : i+l]
}

func getObject[T any](i, l int) T {
	return tConvertByteArrayToObject[T](getMemory(i, l))
}

// 分配内存
func allocateMemory(l int) {
	poolByte = make([]byte, l)
}

type _pool1[T any] struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func (p *_pool1[T]) allocateMemory(c int) {
	var e T
	p.eSize = int(unsafe.Sizeof(e))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *_pool1[T]) isFree(b, e int) bool {
	for _, b := range p.b[b:e] {
		if b != 0 {
			return false
		}
	}
	return true
}

func (p *_pool1[T]) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *_pool1[T]) getObject() *T {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			return nil
		}
	}

	p.last = (i + 1) % p.eCount
	return tConvertByteArrayToObject[*T](p.b[i*p.eSize : (i+1)*p.eSize])
}

type anyInit interface {
	Init()
	Use()
}

type _pool2[T any, AT anyInit] struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func (p *_pool2[T, AT]) allocateMemory(c int) {
	var e T
	p.eSize = int(unsafe.Sizeof(e))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *_pool2[T, AT]) isFree(b, e int) bool {
	for _, b := range p.b[b:e] {
		if b != 0 {
			return false
		}
	}
	return true
}

func (p *_pool2[T, AT]) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *_pool2[T, AT]) getObject() AT {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			var at AT
			return at
		}
	}

	p.last = (i + 1) % p.eCount
	return tConvertByteArrayToObject[AT](p.b[i*p.eSize : (i+1)*p.eSize])
}

type TC interface {
	*simpleStruct
	Init()
	Use()
}

type _pool3[AT TC] struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func (p *_pool3[T]) allocateMemory(c int, o T) {
	p.eSize = int(unsafe.Sizeof(*o))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *_pool3[T]) isFree(b, e int) bool {
	for _, b := range p.b[b:e] {
		if b != 0 {
			return false
		}
	}
	return true
}

func (p *_pool3[T]) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *_pool3[T]) getObject() T {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			var t T
			return t
		}
	}

	p.last = (i + 1) % p.eCount
	return tConvertByteArrayToObject[T](p.b[i*p.eSize : (i+1)*p.eSize])
}

type _pool4 struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func (p *_pool4) allocateMemory(c int) {
	p.eSize = int(unsafe.Sizeof(simpleStruct{}))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *_pool4) isFree(b, e int) bool {
	for _, b := range p.b[b:e] {
		if b != 0 {
			return false
		}
	}
	return true
}

func (p *_pool4) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

// 内存反转函数
func convertByteArrayToObject(b []byte) *simpleStruct {
	return *(**simpleStruct)(unsafe.Pointer(&b))
}

func (p *_pool4) getObject() *simpleStruct {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			return nil
		}
	}

	p.last = (i + 1) % p.eCount
	return convertByteArrayToObject(p.b[i*p.eSize : (i+1)*p.eSize])
}

// 简单结构体
type simpleStruct struct {
	b bool
	v int
}

func (s *simpleStruct) Init() {
	s.b, s.v = true, math.MaxInt64
}

func (s *simpleStruct) Use() {
	s.b, s.v = false, rand.Intn(s.v)
}

func (s *simpleStruct) Free() {
	s.b, s.v = false, 0
}

// 复杂结构体
type complexStruct struct {
	ss  *simpleStruct
	sl  []byte
	str string
	m   map[string]int
}

func (s *complexStruct) Init() {
	s.ss, s.sl, s.str, s.m = nil, nil, "", nil
}

func (s *complexStruct) Use(ss *simpleStruct, sl []byte, str string, m map[string]int) {
	s.ss, s.sl, s.str, s.m = ss, sl, str, m
}

func (s *complexStruct) Free() {
	s.ss, s.sl, s.str, s.m = nil, nil, "", nil
}

// 内存池测试函数
func poolFoo(c int) {
	// 第一种写法
	// 初始化：
	// - 不需要初始化
	// 分配内存：
	// - 需要外部计算对象大小
	// 获取对象：
	// - 需要外部提供数组的下标
	// - 需要显式传递类型参数：结构体指针
	// - 不需要一般接口：anyInit
	var t simpleStruct
	tSize := int(unsafe.Sizeof(t))
	allocateMemory(c * tSize)
	fmt.Printf("pool = %v, len %v\n", poolByte, len(poolByte))

	// 第二种写法
	// 初始化：
	// - 需要显式传递类型参数：结构体
	// 分配内存：
	// - 不需要外部计算对象大小
	// 获取对象：
	// - 不需要外部提供数组的下标
	// - 不需要显式传递类型参数
	// - 不需要定义任何接口
	pool1 := &_pool1[simpleStruct]{}
	pool1.allocateMemory(c)
	fmt.Printf("pool1 = %v, len %v\n", pool1.b, len(pool1.b))

	// 第三种写法
	// 初始化：
	// - 需要显式传递类型参数：结构体及其指针类型
	// 分配内存：
	// - 不需要外部计算对象大小
	// 获取对象：
	// - 不需要外部提供数组的下标
	// - 不需要显式传递类型参数
	// - 需要定义一般接口：anyInit
	pool2 := &_pool2[simpleStruct, *simpleStruct]{}
	pool2.allocateMemory(c)
	fmt.Printf("pool2 = %v, len %v\n", pool2.b, len(pool2.b))

	// 第四种写法
	// 初始化：
	// - 需要显式传递类型参数：结构体指针类型
	// 分配内存：
	// - 需要外部计算对象大小
	// 获取对象：
	// - 不需要外部提供数组的下标
	// - 不需要显式传递类型参数
	// - 需要定义泛型接口：TC
	pool3 := &_pool3[*simpleStruct]{}
	pool3.allocateMemory(c, &simpleStruct{})
	fmt.Printf("pool3 = %v, len %v\n", pool3.b, len(pool3.b))

	// 第五种写法：无泛型
	// 初始化：
	// - 不需要显式传递类型参数
	// 分配内存：
	// - 不需要外部计算对象大小
	// 获取对象：
	// - 不需要外部提供数组的下标
	// - 不需要显式传递类型参数
	// - 不需要定义任何接口
	pool4 := &_pool4{}
	pool4.allocateMemory(c)
	fmt.Printf("pool4 = %v, len %v\n", pool4.b, len(pool4.b))

	pool5 := &_pool5[simpleStruct]{}
	pool5.allocateMemory(c)
	fmt.Printf("pool5 = %v, len %v\n", pool5.b, len(pool5.b))

	for i := 0; i < c; i++ {
		o := getObject[*simpleStruct](tSize*i, tSize)
		o.Init()
		fmt.Printf("i %v, o %v, ptr %p\n", i, o, &o)
		o.Use()
		fmt.Printf("i %v, o %v, ptr %p\n", i, o, &o)
		fmt.Printf("i %v, pool = %v\n", i, poolByte)

		o1 := pool1.getObject()
		o1.Init()
		fmt.Printf("i %v, o1 %v, ptr %p\n", i, o1, &o1)
		o1.Use()
		fmt.Printf("i %v, o1 %v, ptr %p\n", i, o1, &o1)
		fmt.Printf("i %v, pool1 = %v\n", i, pool1.b)

		o2 := pool2.getObject()
		o2.Init()
		fmt.Printf("i %v, o2 %v, ptr %p\n", i, o2, &o2)
		o2.Use()
		fmt.Printf("i %v, o2 %v, ptr %p\n", i, o2, &o2)
		fmt.Printf("i %v, pool2 = %v\n", i, pool2.b)

		o3 := pool3.getObject()
		o3.Init()
		fmt.Printf("i %v, o3 %v, ptr %p\n", i, o3, &o3)
		o3.Use()
		fmt.Printf("i %v, o3 %v, ptr %p\n", i, o3, &o3)
		fmt.Printf("i %v, pool3 = %v\n", i, pool3.b)

		o4 := pool4.getObject()
		o4.Init()
		fmt.Printf("i %v, o4 %v, ptr %p\n", i, o4, &o4)
		o4.Use()
		fmt.Printf("i %v, o4 %v, ptr %p\n", i, o4, &o4)
		fmt.Printf("i %v, pool4 = %v\n", i, pool4.b)

		o5 := pool5.getObject()
		o5.V().Init()
		fmt.Printf("i %v, o5 %v, ptr %p\n", i, o5, &o5)
		o5.V().Use()
		fmt.Printf("i %v, o5 %v, ptr %p\n", i, o5, &o5)
		fmt.Printf("i %v, pool5 = %v\n", i, pool5.b)
	}
}

// 内存池对比测试函数
func poolFooOrigin(c int) {
	pool := &_pool1[simpleStruct]{}
	pool.allocateMemory(c)

	for i := 1; i <= c; i++ {
		os := make([]*simpleStruct, 0, i)
		for j := 0; j < i; j++ {
			o := pool.getObject()
			// fmt.Printf("i %v, j %v, get, o %v, ptr %p\n", i, j, o, &o)
			o.Init()
			// fmt.Printf("i %v, j %v, init o %v, ptr %p\n", i, j, o, &o)
			o.Use()
			// fmt.Printf("i %v, j %v, use o %v, ptr %p\n", i, j, o, &o)
			os = append(os, o)
		}
		// fmt.Printf("before free pool = %v\n", pool.b)
		for _, o := range os {
			// fmt.Printf("before free o %v, ptr %p\n", o, &o)
			o.Free()
			// fmt.Printf("after free o %v, ptr %p\n", o, &o)
		}
		// utility.ForceGC(c, 1)
		// fmt.Printf("after free pool = %v\n", pool.b)
		// fmt.Println()
	}
}

// 内存池对比测试函数
func poolFooCompare(c int) {
	pool := sync.Pool{
		New: func() any {
			// fmt.Println("new simpleStruct")
			return &simpleStruct{}
		},
	}

	for i := 0; i <= c; i++ {
		os := make([]*simpleStruct, 0, i)
		for j := 0; j < i; j++ {
			o := pool.Get().(*simpleStruct)
			// fmt.Printf("i %v, j %v, get, o %v, ptr %p\n", i, j, o, &o)
			o.Init()
			// fmt.Printf("i %v, j %v, init o %v, ptr %p\n", i, j, o, &o)
			o.Use()
			// fmt.Printf("i %v, j %v, use o %v, ptr %p\n", i, j, o, &o)
			os = append(os, o)
		}
		for _, o := range os {
			// fmt.Printf("before put o %v, ptr %p\n", o, &o)
			pool.Put(o)
			// fmt.Printf("after put o %v, ptr %p\n", o, &o)
		}
		// utility.ForceGC(c, 1)
		// fmt.Println()
	}
}

func poolFooOrigin1(c int) {
	pool := &_pool1[complexStruct]{}
	pool.allocateMemory(c)
	m1Pool := &_pool1[simpleStruct]{}
	m1Pool.allocateMemory(c)
	m2Pool := &_pool1[[]byte]{}
	m2Pool.allocateMemory(c)
	m3Pool := &_pool1[string]{}
	m3Pool.allocateMemory(c)
	m4Pool := &_pool1[map[string]int]{}
	m4Pool.allocateMemory(c)

	for i := 1; i <= c; i++ {
		os := make([]*complexStruct, 0, i)
		for j := 0; j < i; j++ {
			o := pool.getObject()
			fmt.Printf("i %v, j %v, get, o %v, ptr %p\n", i, j, o, &o)
			o.Init()
			fmt.Printf("i %v, j %v, init o %v, ptr %p\n", i, j, o, &o)
			o.Use(m1Pool.getObject(), *m2Pool.getObject(), *m3Pool.getObject(), *m4Pool.getObject())
			fmt.Printf("i %v, j %v, use o %v, ptr %p\n", i, j, o, &o)
			os = append(os, o)
		}
		fmt.Printf("before free pool = %v\n", pool.b)
		for _, o := range os {
			fmt.Printf("before free o %v, ptr %p\n", o, &o)
			o.Free()
			fmt.Printf("after free o %v, ptr %p\n", o, &o)
		}
		utility.ForceGC(c, 1)
		fmt.Printf("after free pool = %v\n", pool.b)
		fmt.Println()
	}
}

func poolFooCompare1(c int) {
	pool := sync.Pool{
		New: func() any {
			// fmt.Println("new complexStruct")
			return &complexStruct{}
		},
	}
	m1Pool := sync.Pool{
		New: func() any {
			// fmt.Println("new simpleStruct")
			return &simpleStruct{}
		},
	}
	m2Pool := sync.Pool{
		New: func() any {
			// fmt.Println("new []byte{}")
			return []byte{}
		},
	}
	m3Pool := sync.Pool{
		New: func() any {
			// fmt.Println("new string")
			return ""
		},
	}
	m4Pool := sync.Pool{
		New: func() any {
			// fmt.Println("new map[string]int{}")
			return map[string]int{}
		},
	}

	for i := 0; i <= c; i++ {
		os := make([]*complexStruct, 0, i)
		for j := 0; j < i; j++ {
			o := pool.Get().(*complexStruct)
			// fmt.Printf("i %v, j %v, get, o %v, ptr %p\n", i, j, o, &o)
			o.Init()
			// fmt.Printf("i %v, j %v, init o %v, ptr %p\n", i, j, o, &o)
			m1 := m1Pool.Get().(*simpleStruct)
			m2 := m2Pool.Get().([]byte)
			m3 := m3Pool.Get().(string)
			m4 := m4Pool.Get().(map[string]int)
			o.Use(m1, m2, m3, m4)
			// fmt.Printf("i %v, j %v, use o %v, ptr %p\n", i, j, o, &o)
			os = append(os, o)
		}
		for _, o := range os {
			// fmt.Printf("before put o %v, ptr %p\n", o, &o)
			pool.Put(o)
			// fmt.Printf("after put o %v, ptr %p\n", o, &o)
		}
		utility.ForceGC(c, 1)
		// fmt.Println()
	}
}

type A struct {
	Name string
}

func (a *A) Reset() {
	a.Name = ""
}

type poolObj[T any] struct {
	b bool
	v T
}

func (o *poolObj[T]) V() *T {
	return &o.v
}

func (o *poolObj[T]) Free() {
	o.b = false
}

type _pool5[T any] struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func (p *_pool5[T]) allocateMemory(c int) {
	var e poolObj[T]
	p.eSize = int(unsafe.Sizeof(e))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *_pool5[T]) isFree(b, e int) bool {
	return p.b[b] == 0
}

func (p *_pool5[T]) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *_pool5[T]) getObject() *poolObj[T] {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			return nil
		}
	}

	p.last = (i + 1) % p.eCount
	p.b[i*(p.eSize)] = 1
	return tConvertByteArrayToObject[*poolObj[T]](p.b[i*p.eSize : (i+1)*p.eSize])
}

// 内存池对比测试函数
func poolFooOrigin2(c int) {
	pool := &_pool1[A]{}
	pool.allocateMemory(c)

	// for i := 0; i < c; i++ {
	// 	o := pool.getObject()
	// 	o.V().Reset()
	// 	o.V().Name = "Hello Pool"
	// 	o.Free()
	// 	// if i == c/2 {
	// 	// 	utility.ForceGC(c, 1)
	// 	// }
	// }

	// fmt.Printf("%v\n", pool.b)

	for i := 1; i <= c; i++ {
		os := make([]*A, 0, i)
		for j := 0; j < i; j++ {
			o := pool.getObject()
			o.Reset()
			o.Name = "Hello Pool"
			os = append(os, o)
		}
		for _, o := range os {
			o.Reset()
		}
		// utility.ForceGC(c, 1)
	}
}

// 内存池对比测试函数
func poolFooCompare2(c int) {
	pool := sync.Pool{
		New: func() any {
			return new(A)
		},
	}

	// for i := 0; i < c; i++ {
	// 	o := pool.Get().(*A)
	// 	o.Reset()
	// 	o.Name = "Hello Pool"
	// 	pool.Put(o)
	// 	// if i == c/2 {
	// 	// 	utility.ForceGC(c, 1)
	// 	// }
	// }

	for i := 0; i <= c; i++ {
		os := make([]*A, 0, i)
		for j := 0; j < i; j++ {
			o := pool.Get().(*A)
			o.Reset()
			o.Name = "Hello Pool"
			os = append(os, o)
		}
		for _, o := range os {
			pool.Put(o)
		}
		// utility.ForceGC(c, 1)
	}
}

type largeStruct struct {
	X          int
	Y          int
	Z          int
	ATK        int
	DEF        int
	LUCKY      uint
	HP         uint
	MP         uint
	Scroll     int
	AIState    uint
	ConnState  uint
	otherData1 string // TODO: no zero value
	otherData2 []int
	otherData3 map[int]int
	otherData4 chan int
}

func (s *largeStruct) Reset() {
	s.X, s.Y, s.Z = 0, 0, 0
	s.ATK, s.DEF = 0, 0
	s.LUCKY = 0
	s.HP, s.MP = 0, 0
	s.Scroll = 0
	s.AIState = 0
	s.ConnState = 0
	s.otherData1 = ""
	s.otherData2 = nil
	s.otherData3 = nil
	s.otherData4 = nil
}

func (s *largeStruct) Init() {
	s.X, s.Y, s.Z = 1, 2, 3
	s.ATK, s.DEF = 10, 10
	s.LUCKY = 1024
	s.HP, s.MP = 10, 10
	s.Scroll = 10
	s.AIState = 1
	s.ConnState = 2
	s.otherData1 = "asd"
	s.otherData2 = make([]int, 5)
	s.otherData2[2] = 4
	s.otherData3 = make(map[int]int)
	s.otherData3[1] = 2
	s.otherData4 = make(chan int, 10)
}

func (s *largeStruct) Use() {
	s.X, s.Y, s.Z = 1, 2, 3
	if s.X != 1 || s.Y != 2 || s.Z != 3 {
		goto PANIC
	}
	if s.ATK != 10 || s.DEF != 10 {
		goto PANIC
	}
	if s.LUCKY != 1024 {
		goto PANIC
	}
	if s.HP != 10 || s.MP != 10 {
		goto PANIC
	}
	if s.Scroll != 10 {
		goto PANIC
	}
	if s.AIState != 1 {
		goto PANIC
	}
	if s.ConnState != 2 {
		goto PANIC
	}
	if s.otherData1 != "asd" {
		goto PANIC
	}
	if len(s.otherData2) != 5 || s.otherData2[2] != 4 {
		goto PANIC
	}
	if v, has := s.otherData3[1]; v != 2 || !has {
		goto PANIC
	}
	select {
	case s.otherData4 <- 1:
	default:
		goto PANIC
	}
	return
PANIC:
	panic(s)
}

// 内存池对比测试函数
func poolFooOrigin3(c int) {
	pool := &_pool1[largeStruct]{}
	pool.allocateMemory(c)

	for i := 1; i <= c; i++ {
		os := make([]*largeStruct, 0, i)
		for j := 0; j < i; j++ {
			o := pool.getObject()
			o.Init()
			o.Use()
			os = append(os, o)
		}
		for _, o := range os {
			o.Reset()
		}
	}
}

// 内存池对比测试函数
func poolFooCompare3(c int) {
	pool := sync.Pool{
		New: func() any { return new(largeStruct) },
	}

	for i := 0; i <= c; i++ {
		os := make([]*largeStruct, 0, i)
		for j := 0; j < i; j++ {
			o := pool.Get().(*largeStruct)
			o.Init()
			o.Use()
			os = append(os, o)
		}
		for _, o := range os {
			pool.Put(o)
		}
	}
}

// ----------------------------------------------------------------

type poolCache[T any] struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func (p *poolCache[T]) allocateMemory(c int) {
	var e T
	p.eSize = int(unsafe.Sizeof(e))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *poolCache[T]) isFree(b, e int) bool {
	for _, b := range p.b[b:e] {
		if b != 0 {
			return false
		}
	}
	return true
}

func (p *poolCache[T]) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *poolCache[T]) getObject() *T {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			return nil
		}
	}

	p.last = (i + 1) % p.eCount
	return tConvertByteArrayToObject[*T](p.b[i*p.eSize : (i+1)*p.eSize])
}

func (p *poolCache[T]) isKey(b, e int, k []byte) bool {
	for i, b := range p.b[b:e] {
		if b != k[i] {
			return false
		}
	}
	return true
}

func (p *poolCache[T]) scanKey(b, e int, k []byte) int {
	for i := b; i < e; i++ {
		if !p.isKey(i*p.eSize, (i+1)*p.eSize, k) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *poolCache[T]) getCache(k []byte) *T {
	i := p.scanKey(0, p.last, k)
	if i == -1 {
		i = p.scanKey(p.last, p.eCount, k)
		if i == -1 {
			return nil
		}
	}

	return tConvertByteArrayToObject[*T](p.b[i*p.eSize : (i+1)*p.eSize])
}

func poolCacheFooOrigin(c int) {
	pool := &poolCache[A]{}
	pool.allocateMemory(c)

	for i := 1; i <= c; i++ {
		for j := 0; j < i; j++ {
			var o *A
			key := []byte(fmt.Sprintf("Hello Pool %v", j))
			ob := pool.getCache(key) // TODO: not save key
			if ob == nil {
				o = pool.getObject()
			}
			o.Reset()
			o.Name = "Hello Pool"
		}
	}
}

func poolCacheFooCompare(c int) {
	var a A
	pool := freecache.NewCache(int(unsafe.Sizeof(a)) * c)

	for i := 0; i <= c; i++ {
		for j := 0; j < i; j++ {
			var o *A
			key := []byte(fmt.Sprintf("Hello Pool %v", j))
			ob, err := pool.Get(key)
			if err != nil && !errors.Is(err, freecache.ErrNotFound) {
				panic(err)
			}
			if len(ob) == 0 {
				o = &A{}
			} else {
				o = tConvertByteArrayToObject[*A](ob)
			}
			o.Reset()
			o.Name = "Hello Pool"
			pool.Set(key, TConvertObjectToByteArray(&a), 0)
		}
	}
}

// ----------------------------------------------------------------

// poolDequeue is a lock-free fixed-size single-producer,
// multi-consumer queue. The single producer can both push and pop
// from the head, and consumers can pop from the tail.
//
// It has the added feature that it nils out unused slots to avoid
// unnecessary retention of objects. This is important for sync.Pool,
// but not typically a property considered in the literature.
type poolDequeue struct {
	// headTail packs together a 32-bit head index and a 32-bit
	// tail index. Both are indexes into vals modulo len(vals)-1.
	//
	// tail = index of oldest data in queue
	// head = index of next slot to fill
	//
	// Slots in the range [tail, head) are owned by consumers.
	// A consumer continues to own a slot outside this range until
	// it nils the slot, at which point ownership passes to the
	// producer.
	//
	// The head index is stored in the most-significant bits so
	// that we can atomically add to it and the overflow is
	// harmless.
	headTail uint64

	// vals is a ring buffer of interface{} values stored in this
	// dequeue. The size of this must be a power of 2.
	//
	// vals[i].typ is nil if the slot is empty and non-nil
	// otherwise. A slot is still in use until *both* the tail
	// index has moved beyond it and typ has been set to nil. This
	// is set to nil atomically by the consumer and read
	// atomically by the producer.
	vals []eface
}

type eface struct {
	typ, val unsafe.Pointer
}

const dequeueBits = 32

// dequeueLimit is the maximum size of a poolDequeue.
//
// This must be at most (1<<dequeueBits)/2 because detecting fullness
// depends on wrapping around the ring buffer without wrapping around
// the index. We divide by 4 so this fits in an int on 32-bit.
const dequeueLimit = (1 << dequeueBits) / 4

// dequeueNil is used in poolDequeue to represent interface{}(nil).
// Since we use nil to represent empty slots, we need a sentinel value
// to represent nil.
type dequeueNil *struct{}

func (d *poolDequeue) unpack(ptrs uint64) (head, tail uint32) {
	const mask = 1<<dequeueBits - 1
	head = uint32((ptrs >> dequeueBits) & mask) // 为什么要和 mask 做 & 运算？
	tail = uint32(ptrs & mask)                  // 为什么要和 mask 做 & 运算？
	return
}

func (d *poolDequeue) pack(head, tail uint32) uint64 {
	const mask = 1<<dequeueBits - 1
	return (uint64(head) << dequeueBits) | uint64(tail&mask) // 为什么要和 mask 做 & 运算？
}

// pushHead adds val at the head of the queue. It returns false if the
// queue is full. It must only be called by a single producer.
func (d *poolDequeue) pushHead(val any) bool {
	ptrs := atomic.LoadUint64(&d.headTail)
	head, tail := d.unpack(ptrs)
	if (tail+uint32(len(d.vals)))&(1<<dequeueBits-1) == head {
		// Queue is full.
		return false
	}
	slot := &d.vals[head&uint32(len(d.vals)-1)]

	// Check if the head slot has been released by popTail.
	typ := atomic.LoadPointer(&slot.typ)
	if typ != nil {
		// Another goroutine is still cleaning up the tail, so
		// the queue is actually still full.
		return false
	}

	// The head slot is free, so we own it.
	if val == nil {
		val = dequeueNil(nil)
	}
	*(*any)(unsafe.Pointer(slot)) = val

	// Increment head. This passes ownership of slot to popTail
	// and acts as a store barrier for writing the slot.
	atomic.AddUint64(&d.headTail, 1<<dequeueBits)
	return true
}

// popHead removes and returns the element at the head of the queue.
// It returns false if the queue is empty. It must only be called by a
// single producer.
func (d *poolDequeue) popHead() (any, bool) {
	var slot *eface
	for {
		ptrs := atomic.LoadUint64(&d.headTail)
		head, tail := d.unpack(ptrs)
		if tail == head {
			// Queue is empty.
			return nil, false
		}

		// Confirm tail and decrement head. We do this before
		// reading the value to take back ownership of this
		// slot.
		head--
		ptrs2 := d.pack(head, tail)
		if atomic.CompareAndSwapUint64(&d.headTail, ptrs, ptrs2) {
			// We successfully took back slot.
			slot = &d.vals[head&uint32(len(d.vals)-1)]
			break
		}
	}

	val := *(*any)(unsafe.Pointer(slot))
	if val == dequeueNil(nil) {
		val = nil
	}
	// Zero the slot. Unlike popTail, this isn't racing with
	// pushHead, so we don't need to be careful here.
	*slot = eface{}
	return val, true
}

// popTail removes and returns the element at the tail of the queue.
// It returns false if the queue is empty. It may be called by any
// number of consumers.
func (d *poolDequeue) popTail() (any, bool) {
	var slot *eface
	for {
		ptrs := atomic.LoadUint64(&d.headTail)
		head, tail := d.unpack(ptrs)
		if tail == head {
			// Queue is empty.
			return nil, false
		}

		// Confirm head and tail (for our speculative check
		// above) and increment tail. If this succeeds, then
		// we own the slot at tail.
		ptrs2 := d.pack(head, tail+1)
		if atomic.CompareAndSwapUint64(&d.headTail, ptrs, ptrs2) {
			// Success.
			slot = &d.vals[tail&uint32(len(d.vals)-1)]
			break
		}
	}

	// We now own slot.
	val := *(*any)(unsafe.Pointer(slot))
	if val == dequeueNil(nil) {
		val = nil
	}

	// Tell pushHead that we're done with this slot. Zeroing the
	// slot is also important so we don't leave behind references
	// that could keep this object live longer than necessary.
	//
	// We write to val first and then publish that we're done with
	// this slot by atomically writing to typ.
	slot.val = nil
	atomic.StorePointer(&slot.typ, nil)
	// At this point pushHead owns the slot.

	return val, true
}
