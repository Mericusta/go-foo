package timingwheel

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

func Test_timingWheel(t *testing.T) {
	// 新建一个时间轮
	timingWheel := NewTimingWheel()

	// 挂载时间轮轮盘的刻度
	err := timingWheel.MountSeriesTimingWheelRound(
		RoundMinute,
		RoundHour,
		RoundDay,
		RoundWeek,
	)
	if err != nil {
		panic(err)
	}

	// 挂载初始任务
	err = timingWheel.AddTickerHandler(
		NewTickerHandler("uid_init_0", 0),
		NewTickerHandler("uid_init_15s", time.Second*15),
		NewTickerHandler("uid_init_1min", time.Minute),
		NewTickerHandler("uid_init_1min15s", time.Second*15+time.Minute),
		NewTickerHandler("uid_init_1h", time.Hour),
		NewTickerHandler("uid_init_1h1min15s", time.Second*15+time.Minute+time.Hour),
		NewTickerHandler("uid_init_1d1h1min15s", time.Second*15+time.Minute+time.Hour+time.Hour*24),
	)
	if err != nil {
		panic(err)
	}

	// 运行过程中随机挂载任务
	go func(tw ITimingWheel) {
		uid := 0
		t := time.NewTicker(time.Second)
		for range t.C {
			uid++
			delay := time.Second * time.Duration(rand.Intn(60))
			handler := NewTickerHandler(fmt.Sprintf("uid_%v", uid), delay)
			err := tw.AddTickerHandler(handler)
			if err != nil {
				fmt.Printf("add ticker handler %s occurs error: %v\n", handler, err)
			}
		}
		// roundPosition := 8
		// delay := time.Second * 55
		// t := time.NewTimer(time.Second * time.Duration(roundPosition))
		// <-t.C
		// tw.AddTickerHandler(NewTickerHandler(fmt.Sprintf("trigger_on_%vs", delay.Seconds()), delay))
	}(timingWheel)

	// 启动时间轮
	err = timingWheel.Start()
	if err != nil {
		panic(err)
	}
}
