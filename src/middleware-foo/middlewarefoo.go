package middlewarefoo

import "fmt"

type interfaceA interface {
	A() *implementA
}

type implementA struct{}

func (ia *implementA) A() *implementA { return ia }

type interfaceB interface {
	interfaceA
	B() *implementB
}

type implementB struct {
	*implementA
}

func (ib *implementB) B() *implementB { return ib }

func funcInterfaceA(ia interfaceA) {
	ia.A()
	fmt.Printf("funcInterfaceA\n")
}

func newInterfaceA() interfaceA {
	return &implementA{}
}

func funcInterfaceB(ib interfaceB) {
	ib.B()
	fmt.Printf("funcInterfaceB\n")
}

func newInterfaceB(ia interfaceA) interfaceB {
	return &implementB{implementA: ia.A()}
}

type basement struct {
	middlewareSlice []func(interfaceB)
}

func (b *basement) handle(ia interfaceA) {
	// 中间件包装
	ib := newInterfaceB(ia)
	for _, md := range b.middlewareSlice {
		md(ib)
	}

	// 原始 handler
	funcInterfaceA(ib)
}

// 接口参数包装 middleware
// 将 func(interfaceA) -> 转换成 func(interfaceB)
func HandlerMiddlewareFoo(ia interfaceA) {
	basement := &basement{}
	basement.middlewareSlice = make([]func(interfaceB), 0)
	basement.middlewareSlice = append(basement.middlewareSlice, func(ib interfaceB) {
		fmt.Printf("custom handle\n")
	}, funcInterfaceB)

	basement.handle(ia)
}
