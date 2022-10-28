package middlewarefoo

import (
	"fmt"
)

// ---------------- 应用层：业务层 ----------------

// handler with other server user context
func funcInterfaceC(ic interfaceC) {
	ic.C()
	fmt.Printf("funcInterfaceC\n")
}

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

func HandlerMiddlewareFoo(ia interfaceA, withUserContext, withOtherServerUserContext bool) {
	userMgr = make(map[interfaceA]struct{})
	if withUserContext {
		userMgr[ia] = struct{}{}
	}
	otherServerUserMgr = make(map[interfaceA]struct{})
	if withOtherServerUserContext {
		otherServerUserMgr[ia] = struct{}{}
	}

	basement := &basement{}
	// TODO: 如何支持多种 middleware
	basement.handlerMiddleware = &userImplementMiddleware{
		f: funcInterfaceB,
	}

	basement.handle(ia)
}

// ---------------- 应用层：服务层 ----------------

var userMgr map[interfaceA]struct{}

// user
type implementB struct {
	*implementA
}

// user implement user context
func (ib *implementB) B() *implementB { return ib }

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

// user middleware
type userImplementMiddleware struct {
	f func(interfaceB)
}

// user middleware implement handler middleware interface
func (m *userImplementMiddleware) Do(ia interfaceA) bool {
	ib := newInterfaceB(ia)
	if ib == nil {
		return true
	}
	m.f(ib)
	return false
}

var otherServerUserMgr map[interfaceA]struct{}

// server user: other server link as user
type implementC struct {
	*implementA
}

// server user implenent server user context
func (ic *implementC) C() *implementC { return ic }

// other server user context
type interfaceC interface {
	interfaceA
	C() *implementC
}

// other server user context constructor
func newInterfaceC(ia interfaceA) interfaceC {
	if _, has := userMgr[ia]; has {
		return &implementC{implementA: ia.A()}
	}
	return nil
}

// other server user middleware
type otherServerUserImplementMiddleware struct {
	f func(interfaceC)
}

// user middleware implement handler middleware interface
func (m *otherServerUserImplementMiddleware) Do(ia interfaceA) bool {
	ib := newInterfaceC(ia)
	if ib == nil {
		return true
	}
	m.f(ib)
	return false
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
	Do(interfaceA) bool
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
	if !b.handlerMiddleware.Do(ia) { // TODO: 原始 handler 不需要中间件，如何跳过？
		return
	}
	// 原始 handler
	funcInterfaceA(ia)
}
