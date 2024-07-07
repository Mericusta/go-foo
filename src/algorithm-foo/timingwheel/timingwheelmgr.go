package timingwheel

import (
	"fmt"
	"sort"
	"time"
)

var (
	errorTimingWheelAlreadyRunning          error = fmt.Errorf("timing wheel already running")
	errorTimingWheelRoundAlreadyExists      error = fmt.Errorf("timing wheel round already exists")
	errorInvalidTimingWheelConcurrencyCount error = fmt.Errorf("timing wheel concurrency count is invalid")
	errorTimingWheelDiscNotExists           error = fmt.Errorf("timing wheel disc not exists")
)

var (
	DefaultTick             time.Duration = time.Second // 默认最小刻度
	DefaultConcurrencyCount int           = 1           // 默认并发数
)

// timingWheel 时间轮
type timingWheel struct {
	id               string              // 该时间轮的标识
	running          bool                // 是否运行中
	pointer          *timingWheelPointer // 时间轮的指针
	initTickHandlers []*tickHandler      // 时间轮的初始时刻挂载的行为
}

// NewTimingWheel 新建一个时间轮
func NewTimingWheel() *timingWheel {
	return &timingWheel{}
}

// MountSeriesTimingWheelRound 挂载一系列时间刻度
func (tmm *timingWheel) MountSeriesTimingWheelRound(rounds ...*timingWheelRound) error {
	// 是否运行中
	if tmm.running {
		return errorTimingWheelAlreadyRunning
	}
	// 去重，去错
	roundsMap := make(map[time.Duration]int)
	for index, round := range rounds {
		if round == nil || time.Duration(round.slot)*round.tick == 0 {
			continue
		}
		roundsMap[round.tick] = index
	}
	// 排序
	_rounds := make([]*timingWheelRound, 0, len(roundsMap))
	for _, index := range roundsMap {
		_rounds = append(_rounds, rounds[index])
	}
	sort.Slice(_rounds, func(i, j int) bool { return _rounds[i].tick < _rounds[j].tick })
	// 设置
	tmm.pointer = newTimingWheelPointer(newSeriesTimingWheelDisc(rounds...))
	return nil
}

// Start 启动时间轮
func (tmm *timingWheel) Start() error {
	// 是否运行中
	if tmm.running {
		return errorTimingWheelAlreadyRunning
	}
	tmm.running = true
	fmt.Printf("timingWheelMgr.Start\n")
	tmm.pointer.start()
	return nil
}

// AddTickerHandler 添加时刻回调
func (tmm *timingWheel) AddTickerHandler(ths ...*tickHandler) error {
	if tmm.pointer == nil {
		return errorTimingWheelDiscNotExists
	}
	return tmm.pointer.addTickerHandler(ths...)
}
