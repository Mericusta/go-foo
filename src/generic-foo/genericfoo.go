package genericfoo

import (
	"fmt"
)

// 真·接口
type Adder[T any] interface {
	Add(T, T)
}

type Adder2 interface {
	Add2(any, any)
}

type Adder3[T MyString | MyInt] interface {
	*MyString | *MyInt
	Add3(*T)
}

func Add3[T Adder3[TT], TT MyString | MyInt](l T, r *TT) {
	l.Add3(r)
}

type MyString struct{}

func NewMyString() *MyString {
	return &MyString{}
}

func (a *MyString) Add(S1, S2 MyString) {

}

func (a *MyString) Add2(S1, S2 any) {

}

func (a *MyString) Add3(s *MyString) {

}

type MyInt struct{}

func (a *MyInt) Add(S1, S2 MyInt) {

}

func (a *MyInt) Add2(S1, S2 any) {

}

func (a *MyInt) Add3(s *MyInt) {

}

func Insert[T any](e ...Adder[T]) {
	var c []Adder[T]
	c = append(c, e...)
}

func Insert2(e ...Adder2) {
	var c []Adder2
	c = append(c, e...)
}

func f() {
	var i1 Adder[MyString] = &MyString{}
	var i2 Adder[MyInt] = &MyInt{}
	Insert(i1)
	Insert(i2)

	Insert2(&MyString{}, &MyInt{})

	Add3(&MyString{}, &MyString{})
	// Add3(&MyString{}, &MyInt{})
}

// --------------------------------

type IBase interface {
	Init()
}

type IA interface {
	IBase
	A()
}

type IB interface {
	IBase
	B()
}

type A struct{}

func (a *A) Init() {}

func (a *A) A() {}

var Ap *A = &A{}

type B struct{}

func (b *B) Init() {}

func (b *B) B() {}

var Bp *B = &B{}

var iMap map[IBase]interface{} = map[IBase]interface{}{
	Ap: Ap,
	Bp: Bp,
}

func InitAllWithMap() {
	for i := range iMap {
		i.Init()
	}
}

func InitAll() {
	Ap.Init()
	Bp.Init()
}

func GetI[T IBase]() T {
	var iT T
	return iMap[iT].(T)
}

func GetWithMap() {
	_ = GetI[IA]()
	_ = GetI[IB]()
}

func GetIA() IA {
	return Ap
}

func GetIB() IB {
	return Bp
}

func Get() {
	_ = GetIA()
	_ = GetIB()
}

// --------------------------------

type structA struct{}

func (s *structA) call() { fmt.Println("structA call") }

type structB struct{}

func (s *structB) call() { fmt.Println("structB call") }

type basicnterface interface{ call() }

func useBasicInterface(s basicnterface) {
	s.call()
}

type typeConstraints interface{ structA | structB }
type generalInterface interface {
	*structA | *structB
	call()
}

func useGeneralInterface[T generalInterface](s T) {
	s.call()
}

// // to make it work, we need 'generic general interface'
// func canNotWork[T typeConstraints](s *T) {
// 	// type constraints determines what operations are available on T
// 	// it doesn't imply anything about *T
// 	// https://stackoverflow.com/questions/71444847/go-with-generics-type-t-is-pointer-to-type-parameter-not-type-parameter
// 	s.call() // type *T is pointer to type parameter, not type parameter
// }

// style 1: declare generic general interface inside type param scope
func useGGIInsideTypeParamScope[
	T typeConstraints, // type 1, use typeconstraints to constrain types
	PT interface { // type 2, use a general interface to constrain types
		*T     // type constraints
		call() // method
	}](s *T) {
	PT(s).call()
}

// style 2: declare generic general interface outside type param scope
type genericGeneralInterface[T typeConstraints] interface { // generic
	*T            // type constraints
	basicnterface // method
}

func useGGIOutsideTypeParamScope[T typeConstraints, PT genericGeneralInterface[T]](s *T) {
	PT(s).call()
}

func useGGIOutsideTypeParamScope1[T typeConstraints, PT genericGeneralInterface[T]](s PT) {
	s.call()
}

func generalInterfaceCall() {
	useBasicInterface(&structA{})
	useBasicInterface(&structB{})

	useGeneralInterface(&structA{})
	useGeneralInterface(&structB{})

	useGGIOutsideTypeParamScope(&structA{})
	useGGIOutsideTypeParamScope(&structB{})

	useGGIOutsideTypeParamScope1(&structA{})
	useGGIOutsideTypeParamScope1(&structB{})
}

// --------

// 指针的结构体类型萃取 *p -> p
// 利用 类型参数表 中 指针类型的类型参数 萃取指针的结构体类型
// 为了在 类型参数表 中 构造指针类型的参数，需要使用 类型的类型（模板的模板）约束
func structPointerTypeTraitFoo[T any](tv *T) T {
	return *tv
}

// 指针的结构体类型萃取，传递到函数中
type traitStruct struct{}

func (s *traitStruct) GetProtocolID() uint32 { return 1024 }

func structPointerTypeTraitFooWithFunc() {
	RegisterServerHandler(func(IRobotServerContext, *traitStruct) {})
}

type ProtocolMsg any
type ProtocolID uint32

var protocolMakerMap map[ProtocolID]func() ProtocolMsg = make(map[ProtocolID]func() ProtocolMsg)

func RegisterProtocolMaker(id ProtocolID, f func() ProtocolMsg) {
	protocolMakerMap[id] = f
}

type IRobotServerContext any
type MessageProtocol interface{ GetProtocolID() uint32 }

var serverHandlerMap = make(map[ProtocolID]func(IRobotServerContext, MessageProtocol))

func RegisterServerHandler[T func(IRobotServerContext, *TT), TT any](handler T) {
	msgID := register(serverHandlerMap, handler)
	RegisterProtocolMaker(msgID, func() ProtocolMsg { return new(TT) })
}

func register[T any, C any](m map[ProtocolID]func(C, MessageProtocol), handler func(C, T)) ProtocolID {
	msgID := ProtocolID(func(t any) MessageProtocol { return t.(MessageProtocol) }(*new(T)).GetProtocolID())
	f := func(ctx C, iMsg MessageProtocol) {
		if msg, ok := iMsg.(T); ok {
			handler(ctx, msg)
		}
	}
	m[msgID] = f
	return msgID
}
