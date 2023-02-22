package genericfoo

import "fmt"

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

type typeConstraints interface{ structA | structB }
type basicnterface interface{ call() }
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

func generalInterfaceCall() {
	useGeneralInterface(&structA{})
	useGeneralInterface(&structB{})

	useGGIOutsideTypeParamScope(&structA{})
	useGGIOutsideTypeParamScope(&structB{})
}
