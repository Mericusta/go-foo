package middlewarefoo

import (
	"fmt"
)

// ---------------- 应用层 ----------------

var userMgr map[interfaceA]struct{}

// user context
type interfaceB interface {
	interfaceA
	B() *implementB
}

// user context constructor
func newInterfaceB(ia interfaceA) interfaceB {
	if _, has := userMgr[ia]; has {
		return &implementB{implementA: ia.A()}
	}
	return nil
}

// user
type implementB struct {
	*implementA
}

// user implement user context
func (ib *implementB) B() *implementB { return ib }

// handler with user context
func funcInterfaceB(ib interfaceB) {
	ib.B()
	fmt.Printf("funcInterfaceB\n")
}

// handler without user context
func funcInterfaceA(ia interfaceA) {
	ia.A()
	fmt.Printf("funcInterfaceA\n")
}

// user middleware
type implementMiddleware struct {
	f func(interfaceB)
}

// user middleware implement handler middleware interface
func (m *implementMiddleware) do(ia interfaceA) bool {
	ib := newInterfaceB(ia)
	if ib == nil {
		return true
	}
	m.f(ib)
	return false
}

func HandlerMiddlewareFoo(ia interfaceA, withUserContext bool) {
	userMgr = make(map[interfaceA]struct{})
	basement := &basement{}
	basement.handlerMiddleware = &implementMiddleware{
		f: funcInterfaceB,
	}
	if withUserContext {
		userMgr[ia] = struct{}{}
	}

	basement.handle(ia)
}

// ---------------- 框架层 ----------------

// link
type implementA struct{}

// link implement framework context
func (ia *implementA) A() *implementA { return ia }

// framework context
type interfaceA interface {
	A() *implementA
}

// handler middleware interface
type middleware interface {
	do(interfaceA) bool
}

// dispatcher
type basement struct {
	handlerMiddleware middleware
}

// middleware
// - 接口参数包装，将 func(interfaceA) -> 转换成 func(interfaceB)
// - 流程控制，不需要多个 middleware
// logic goroutine
func (b *basement) handle(ia interfaceA) {
	// 中间件包装
	if !b.handlerMiddleware.do(ia) { // TODO: 原始 handler 不需要中间件，如何跳过？
		return
	}
	// 原始 handler
	funcInterfaceA(ia)
}
