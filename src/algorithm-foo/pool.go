package algorithmfoo

import (
	"fmt"
	"go-foo/pkg/utility"
	"math"
	"math/rand"
	"sync"
	"unsafe"
)

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

// ----------------------------------------------------------------
