package moduleTimingWheel

import (
	"time"

	sgs "github.com/Mericusta/go-sgs"
)

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

// TODO: base tick handler
