package algorithmfoo

import (
	"fmt"
	"time"
)

type Timer struct {
	expire  int64
	round   int
	slot    int
	payload interface{}
}

type TimeWheel struct {
	tick            int64
	wheelSize       int
	interval        int64
	currentTime     int64
	slots           [][]*Timer
	next            *TimeWheel
	addTimerChan    chan *Timer
	removeTimerChan chan *Timer
	stopChan        chan struct{}
}

func NewTimeWheel(tick int64, wheelSize int, currentTime int64) *TimeWheel {
	slots := make([][]*Timer, wheelSize)
	for i := 0; i < wheelSize; i++ {
		slots[i] = make([]*Timer, 0)
	}
	return &TimeWheel{
		tick:            tick,
		wheelSize:       wheelSize,
		interval:        tick * int64(wheelSize),
		currentTime:     currentTime,
		slots:           slots,
		addTimerChan:    make(chan *Timer),
		removeTimerChan: make(chan *Timer),
		stopChan:        make(chan struct{}),
	}
}

func (tw *TimeWheel) AddTimer(t *Timer) {
	tw.addTimerChan <- t
}

func (tw *TimeWheel) RemoveTimer(t *Timer) {
	tw.removeTimerChan <- t
}

func (tw *TimeWheel) Start() {
	go tw.run()
}

func (tw *TimeWheel) Stop() {
	close(tw.stopChan)
}

func (tw *TimeWheel) run() {
	ticker := time.NewTicker(time.Duration(tw.tick) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			tw.advance()
		case t := <-tw.addTimerChan:
			tw.add(t)
		case t := <-tw.removeTimerChan:
			tw.remove(t)
		case <-tw.stopChan:
			ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) advance() {
	tw.currentTime += tw.tick
	idx := (tw.currentTime / tw.tick) % int64(tw.wheelSize)
	timers := tw.slots[idx]
	tw.slots[idx] = []*Timer{}
	for _, t := range timers {
		if t.round > 0 {
			t.round--
			tw.add(t)
		} else {
			go tw.trigger(t)
		}
	}
}

func (tw *TimeWheel) add(t *Timer) {
	if t.expire < tw.currentTime {
		go tw.trigger(t)
		return
	}
	if t.expire < tw.currentTime+tw.interval {
		idx := (t.expire / tw.tick) % int64(tw.wheelSize)
		t.slot = int(idx)
		tw.slots[idx] = append(tw.slots[idx], t)
	} else if tw.next != nil {
		tw.next.AddTimer(t)
	} else {
		go tw.trigger(t)
	}
}

func (tw *TimeWheel) remove(t *Timer) {
	idx := (t.expire / tw.tick) % int64(tw.wheelSize)
	slot := tw.slots[idx]
	for i, timer := range slot {
		if timer == t {
			tw.slots[idx] = append(slot[:i], slot[i+1:]...)
			return
		}
	}
}

func (tw *TimeWheel) trigger(t *Timer) {
	fmt.Println("Timer triggered:", t.payload)
}

func timewheelFoo() {
	tw1 := NewTimeWheel(100, 10, 0)   // 每个槽 100 毫秒，共 10 个槽，总时间 1 秒
	tw2 := NewTimeWheel(1000, 10, 0)  // 每个槽 1 秒，共 10 个槽，总时间 10 秒
	tw3 := NewTimeWheel(10000, 10, 0) // 每个槽 10 秒，共 10 个槽，总时间 100 秒

	// 设置多级时间轮
	tw1.next = tw2
	tw2.next = tw3

	tw1.Start()
	tw2.Start()
	tw3.Start()

	defer tw1.Stop()
	defer tw2.Stop()
	defer tw3.Stop()

	// 添加定时器
	now := time.Now().UnixNano() / int64(time.Millisecond)
	tw1.AddTimer(&Timer{
		expire:  now + 1500, // 1.5 秒后触发
		payload: "Timer 1",
	})
	tw1.AddTimer(&Timer{
		expire:  now + 6500, // 6.5 秒后触发
		payload: "Timer 2",
	})
	tw1.AddTimer(&Timer{
		expire:  now + 25000, // 25 秒后触发
		payload: "Timer 3",
	})

	// 运行一段时间以查看输出
	time.Sleep(30 * time.Second)
}
