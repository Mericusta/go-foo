package genericfoo

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
