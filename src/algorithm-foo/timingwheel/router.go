package moduleTimingWheel

// TODO:
// ITimingWheelRouter 时间轮路由器接口
type ITimingWheelRouter interface {
	// PointerCount 挂载到路由器的时间轮的指针数量
	PointerCount() int
}

type timingWheelRouter struct {
	pointerCount int
}

// NewTimingWheelConcurrencyRouter 新建一个时间轮多协程路由器
func NewTimingWheelConcurrencyRouter(pointerCount int) *timingWheelRouter {
	return &timingWheelRouter{pointerCount: pointerCount}
}

var DefaultRouter *timingWheelRouter = &timingWheelRouter{pointerCount: 1}

func (twr *timingWheelRouter) PointerCount() int {
	return twr.pointerCount
}
