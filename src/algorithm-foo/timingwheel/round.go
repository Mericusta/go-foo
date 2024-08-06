package moduleTimingWheel

import (
	"fmt"
	"time"
)

// timingWheelRound 时间轮的周期
type timingWheelRound struct {
	tick     time.Duration // 当前周期的刻度
	slot     int64         // 当前周期的槽数
	duration time.Duration // 当前周期的时长
}

func (twr *timingWheelRound) String() string {
	return fmt.Sprintf("[round %vs / %v]", twr.tick.Seconds(), twr.slot)
}

// NewTimingWheelRound 新建一个时间轮的周期
// @param1             刻度
// @param2             槽数
func NewTimingWheelRound(tick time.Duration, slot int64) *timingWheelRound {
	return &timingWheelRound{tick: tick, slot: slot, duration: tick * time.Duration(slot)}
}

var (
	RoundMinute *timingWheelRound = NewTimingWheelRound(time.Second, 60)
	RoundHour   *timingWheelRound = NewTimingWheelRound(time.Minute, 60)
	RoundDay    *timingWheelRound = NewTimingWheelRound(time.Hour, 24)
	RoundWeek   *timingWheelRound = NewTimingWheelRound(time.Hour*24, 7)
)
