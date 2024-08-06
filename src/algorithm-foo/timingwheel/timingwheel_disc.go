package moduleTimingWheel

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"

	sgs "github.com/Mericusta/go-sgs"
)

// timingWheelDisc 时间轮轮盘
type timingWheelDisc struct {
	// 位置 - 行为列表 聚合指针
	rhsPtr atomic.Uintptr
	// TODO: 挂一个引用防止被 GC
	rhs           *roundHandlers
	roundPosition int64 // 当前轮盘所处周期内的位置
	// TODO: 暂时使用 slice 来处理,后续改成 lock-free queue
	slots [][]ITickHandler  // 该轮盘下所有槽位的待执行行为，所有行为均采用进入时执行
	round *timingWheelRound // 刻度x槽位=当前轮一周期的总时长
	next  *timingWheelDisc  // 下一级轮，必须比当前轮的刻度大
}

type roundHandlers struct {
	roundPosition   int64
	tickHandlersPtr unsafe.Pointer
}

// newTimingWheelDisc 新建一个时间轮轮盘
// @param1            刻度
// @param2            下一级时间轮轮盘
func newTimingWheelDisc(round *timingWheelRound, next *timingWheelDisc) *timingWheelDisc {
	slots := make([][]ITickHandler, round.slot)
	for index := int64(0); index != round.slot; index++ {
		slots[index] = nil
	}

	twd := &timingWheelDisc{round: round, slots: slots, next: next}
	// rhs := &roundHandlers{roundPosition: 0, tickHandlersPtr: unsafe.Pointer(&slots[0])}
	// twd.rhsPtr.Store(uintptr(unsafe.Pointer(rhs)))
	return twd
}

// newSeriesTimingWheelDisc 新建一系列时间轮轮盘
// @param                   系列刻度
func newSeriesTimingWheelDisc(rounds ...*timingWheelRound) *timingWheelDisc {
	// 构造
	timingWheels := make([]*timingWheelDisc, 0, len(rounds))
	for _, round := range rounds {
		timingWheels = append(timingWheels, newTimingWheelDisc(round, nil))
	}
	// 转换成链表
	for i, tw := range timingWheels {
		if i == 0 {
			continue
		}
		timingWheels[i-1].next = tw
	}
	if len(timingWheels) == 0 {
		return nil
	}
	return timingWheels[0]
}

// tick 时间轮盘前进
func (twd *timingWheelDisc) tick(pointer *timingWheelPointer) []ITickHandler {
	// rhs := &roundHandlers{roundPosition: twd.roundPosition + 1, tickHandlersPtr: nil}
	// oldRhsPtr := twd.rhsPtr.Swap(uintptr(unsafe.Pointer(rhs)))
	// oldRhs := (*roundHandlers)(unsafe.Pointer(oldRhsPtr))

	// 获取并清空当前 tick 的行为列表
	handlers := twd.slots[twd.roundPosition]
	twd.slots[twd.roundPosition] = nil
	fmt.Printf("timingWheelDisc.tick, get handlers %v from current %s, position %v, increase\n", len(handlers), twd.round, twd.roundPosition)
	// 所处周期内的位置发生变化
	twd.roundPosition++

	// 达到当前周期最大值时，进位
	if twd.roundPosition == twd.round.slot {
		// 重置当前轮盘的位置
		twd.roundPosition = 0
		fmt.Printf("timingWheelDisc.tick, current %s, reach disc round %vs, reset round position\n", twd.round, twd.round.duration.Seconds())
		// 下一级时间轮盘前进
		if twd.next != nil {
			// 获得下一级时间轮的行为列表
			nextRoundHandlers := twd.next.tick(pointer)
			fmt.Printf("timingWheelDisc.tick, current %s, get and place from next round %s handlers count %v\n", twd.round, twd.next.round, len(nextRoundHandlers))
			// 将行为放置到当前轮盘合适的槽内
			for _, handler := range nextRoundHandlers {
				twd.place(handler)
			}
		} else {
			// TODO: 从无限时间轮中获取若干可以放置到当前时间轮下的行为
		}
	}

	return handlers
}

// place 放置行为
func (twd *timingWheelDisc) place(handler ITickHandler) {
	// 使用指针所在的位置和进位的偏移量计算 handler 的相对时刻
	relativeDuration := time.Duration(twd.roundPosition)*twd.round.tick + handler.Delay()
	// 超过当前周期的总时长
	if relativeDuration >= twd.round.duration {
		// 进位时加上当前的偏移量
		handler.Update(relativeDuration)
		twd.next.place(handler)
		return
	}
	// handler 的合适的位置
	placePosition := int64(relativeDuration / twd.round.tick)
	if placePosition >= twd.round.slot {
		fmt.Printf("timingWheelDisc.place, handler %s, in %s invalid place position %v, relativeDuration %vs, round position %v\n", handler, twd.round, placePosition, relativeDuration.Seconds(), twd.roundPosition)
		return
	} else if placePosition < 0 {
		placePosition = 0
	}
	// 放入到指定位置的队列中
	fmt.Printf("timingWheelDisc.place, handler %v, place to %s position %v, relativeDuration %vs, round position %v\n", handler, twd.round, placePosition, relativeDuration.Seconds(), twd.roundPosition)
	twd.slots[placePosition] = append(twd.slots[placePosition], handler)
}

// timingWheelPointer 时间轮的指针，指针的顶层轮盘决定每次 tick 经过多久
type timingWheelPointer struct {
	ticker *time.Ticker     // 直接依赖 golang 的 ticker，它是 runtime 实现的
	disc   *timingWheelDisc // 当前时间轮的指针所使用的轮盘
	// TODO: 暂时使用 chan 来接收行为,后续改成 lock-free queue
	receiver chan ITickHandler // 接收时刻行为的 chan
	// debug
	counter atomic.Int64 // 计数器
}

// newTimingWheelPointer 新建时间轮指针
func newTimingWheelPointer(disc *timingWheelDisc) *timingWheelPointer {
	return &timingWheelPointer{
		ticker:   time.NewTicker(disc.round.tick),
		disc:     disc,
		receiver: make(chan ITickHandler, 256), // 每个 tick 超过 256 个行为可能会导致精度下降（TODO: 这里应该是所有 disc 的可接收数量加起来，暂时使用最低精度的）
	}
}

// start 启动时间轮指针
func (tmp *timingWheelPointer) start(ctx sgs.IModuleEventContext) {
	// 执行启动时，即0秒的行为
	for {
		select {
		case <-tmp.ticker.C: // 首次 tick 是经过了 disc.round.tick 时间之后的
			// TODO: 这里有可能因为 tmp.receiver 中阻塞
			// TODO: 这里有可能因为执行时间过长而导致精度下降
			fmt.Printf("trigger tick at counter %v\n", tmp.counter.Load())
			tmp.tick(ctx)
			// 进入第 counter 秒
			tmp.counter.Add(1)
		case handler, ok := <-tmp.receiver: // TODO: 有了 lock-free queue 之后不再使用这种形式
			// TODO: 这里有可能因为 ticker.C 中阻塞住导致 receiver 达到上限而被阻塞
			if !ok {
				// TODO: log
				fmt.Printf("receive handler failed then, continue at counter %vs\n", tmp.counter.Load())
				continue
			}
			// 小于最低级时间轮盘的最小刻度直接执行
			if handler.Delay() <= tmp.disc.round.tick {
				fmt.Printf("trigger tick handler %s, trigger at counter %vs\n", handler, tmp.counter.Load())
				handler.Trigger(ctx, 0)
			} else {
				if tmp.counter.Load() != int64(handler.GetPlaceCounter()) {
					fmt.Printf("timingWheelPointer.start, receive and place handler %s at counter %vs not equal\n", handler, tmp.counter.Load())
				}
				tmp.disc.place(handler)
			}
		}
	}
}

// tick 时间轮指针前进
func (tmp *timingWheelPointer) tick(ctx sgs.IModuleEventContext) {
	// 指针所用的轮盘前进
	tickHandlers := tmp.disc.tick(tmp)
	for _, handler := range tickHandlers {
		handler.Trigger(ctx, tmp.counter.Load())
	}
}

// AddTickerHandler 在时间轮中添加时刻行为
func (tmp *timingWheelPointer) addTickerHandler(ths ...ITickHandler) error {
	if tmp.disc == nil {
		return errorTimingWheelDiscNotExists
	}
	for _, handler := range ths {
		handler.SetPlaceCounter(int(tmp.counter.Load()))
		handler.SetExpectCounter(int(tmp.counter.Load()) + int(handler.Delay()/time.Second))
		fmt.Printf("timingWheelPointer.AddTickerHandler %s delay %vs at counter %vs\n", handler, handler.Delay().Seconds(), tmp.counter.Load())
		select {
		case tmp.receiver <- handler:
		default:
			fmt.Printf("timingWheelPointer.AddTickerHandler blocked\n")
		}
	}
	return nil
}
