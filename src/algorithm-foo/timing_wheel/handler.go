package moduleTimingWheel

import (
	"context"
	"time"
)

// 使用接口的方式防止具体逻辑中直接引用该包

type ITickHandler interface {
	Delay() time.Duration
	Update(time.Duration)
	Trigger(context.Context, ...any)

	// 调试接口
	String() string

	GetPlaceCounter() int
	SetPlaceCounter(int)

	GetExpectCounter() int
	SetExpectCounter(int)

	SetUID(int)
	GetUID() int
}
