package moduleTimingWheel

import (
	"time"

	"github.com/Mericusta/go-sgs"
)

// 使用接口的方式防止具体逻辑中直接引用该包

type ITickHandler interface {
	Delay() time.Duration
	Update(time.Duration)
	Trigger(sgs.IModuleEventContext, ...any)

	// 调试接口
	String() string

	GetPlaceCounter() int
	SetPlaceCounter(int)

	GetExpectCounter() int
	SetExpectCounter(int)

	SetUID(int)
	GetUID() int
}
