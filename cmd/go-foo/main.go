package main

import (
	"fmt"
	"go-foo/src/algorithm-foo/timingwheel"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	// 新建一个时间轮
	timingWheel := timingwheel.NewTimingWheel()

	// 挂载时间轮轮盘的刻度
	err := timingWheel.MountSeriesTimingWheelRound(
		timingwheel.RoundMinute,
		timingwheel.RoundHour,
		timingwheel.RoundDay,
		timingwheel.RoundWeek,
	)
	if err != nil {
		panic(err)
	}

	// 挂载初始任务
	err = timingWheel.AddTickerHandler(
		timingwheel.NewTickerHandler("uid_init_0", 0),
		timingwheel.NewTickerHandler("uid_init_15s", time.Second*15),
		timingwheel.NewTickerHandler("uid_init_1min", time.Minute),
		timingwheel.NewTickerHandler("uid_init_1min15s", time.Second*15+time.Minute),
		timingwheel.NewTickerHandler("uid_init_1h", time.Hour),
		timingwheel.NewTickerHandler("uid_init_1h1min15s", time.Second*15+time.Minute+time.Hour),
		timingwheel.NewTickerHandler("uid_init_1d1h1min15s", time.Second*15+time.Minute+time.Hour+time.Hour*24),
	)
	if err != nil {
		panic(err)
	}

	// 运行过程中随机挂载任务
	go func(tw timingwheel.ITimingWheel) {
		uid := 0
		t := time.NewTicker(time.Second)
		for range t.C {
			uid++
			delay := time.Second * time.Duration(rand.Intn(60))
			handler := timingwheel.NewTickerHandler(fmt.Sprintf("uid_%v_%vs", uid, delay.Seconds()), delay)
			err := tw.AddTickerHandler(handler)
			if err != nil {
				fmt.Printf("add ticker handler %s occurs error: %v\n", handler, err)
			}
		}
		// roundPosition := 69
		// delay := time.Second * 51
		// t := time.NewTimer(time.Second * time.Duration(roundPosition))
		// <-t.C
		// tw.AddTickerHandler(timingwheel.NewTickerHandler(fmt.Sprintf("debug_handler_%vs", delay.Seconds()), delay))
	}(timingWheel)

	// 启动时间轮
	err = timingWheel.Start()
	if err != nil {
		panic(err)
	}
}
