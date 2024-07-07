package timingwheel

import (
	"fmt"
	"time"
)

// tickHandler 时刻行为
type tickHandler struct {
	id            string        // 唯一标识
	delay         time.Duration // 多久后触发
	placeCounter  int           // 放置时的计数器
	expectCounter int           // 预期的计数器
}

// NewTickerHandler 新建一个时刻行为
func NewTickerHandler(id string, delay time.Duration) *tickHandler {
	return &tickHandler{id: id, delay: delay}
}

func (th *tickHandler) String() string {
	return fmt.Sprintf("[id %v, place at %vs, expect at %vs]", th.id, th.placeCounter, th.expectCounter)
}
