package middlewarefoo

import "fmt"

// ---------------- 应用层 ----------------

type interfaceB interface {
	interfaceA
	B() *implementB
}

type implementB struct {
	*implementA
}

func (ib *implementB) B() *implementB { return ib }

func funcInterfaceB(ib interfaceB) {
	ib.B()
	fmt.Printf("funcInterfaceB\n")
}

func newInterfaceB(ia interfaceA) interfaceB {
	return &implementB{implementA: ia.A()}
}

// ---------------- 框架层 ----------------
// 框架层不能带有应用层的定义

type implementA struct{}

func (ia *implementA) A() *implementA { return ia }

type interfaceA interface {
	A() *implementA
}

func funcInterfaceA(ia interfaceA) {
	ia.A()
	fmt.Printf("funcInterfaceA\n")
}

func newInterfaceA() interfaceA {
	return &implementA{}
}

type basement struct {
	middlewareSlice []*middleware
}

type middleware struct {
	f func(interfaceB) // TODO: 框架层不能带有应用层的定义
}

func (m *middleware) do(ia interfaceA) {
	ib := newInterfaceB(ia)
	m.f(ib)
}

func (b *basement) handle(ia interfaceA) {
	// 中间件包装
	for _, md := range b.middlewareSlice {
		md.do(ia)
	}

	// 原始 handler
	funcInterfaceA(ia)
}

// 接口参数包装 middleware
// 将 func(interfaceA) -> 转换成 func(interfaceB)
func HandlerMiddlewareFoo(ia interfaceA) {
	basement := &basement{}
	basement.middlewareSlice = make([]*middleware, 0)
	basement.middlewareSlice = append(basement.middlewareSlice, &middleware{
		f: func(ib interfaceB) {
			fmt.Printf("custom handle\n")
		},
	}, &middleware{f: funcInterfaceB})

	basement.handle(ia)
}
