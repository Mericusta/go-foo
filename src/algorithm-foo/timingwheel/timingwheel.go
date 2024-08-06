package moduleTimingWheel

import (
	"fmt"
	"sort"
	"time"

	sgs "github.com/Mericusta/go-sgs"
	"go.uber.org/zap"
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

// ModuleTimingWheel 时间轮模块
type ModuleTimingWheel struct {
	// 组合基础模块
	sgs.ModuleBase

	uidGenerator     int                 // ID 生成器
	rounds           []*timingWheelRound // 时间刻度
	pointer          *timingWheelPointer // 时间轮的指针
	initTickHandlers []ITickHandler      // 时间轮的初始时刻挂载的行为
}

// 时间轮有精度问题，可以通过添加更高精度的时间轮盘来解决
// 精度问题举例：假如当前高精度是1s，那么嵌套的轮盘的 slot 分布如下所示
// |   轮盘  | slot                                                            ...|
// |小时级轮盘| 0                                                               ...|
// |分钟级轮盘| 0                                1                              ...|
// |秒钟级轮盘| 0 1 2 3 4 5 ... 55 56 57 58 59 | 0 1 2 3 4 5 ... 55 56 57 58 59 ...|
// - 初始启动时，立刻执行秒钟级轮盘 slot 0 的所有行为，指针 round 为 0
// - tick 时的执行顺序如下：
//   - 触发 tick，获取需要执行的行为
//   - 轮盘的指针发生变化 round++
//   - 客观计数器发生变化 counter++
//   - 执行 handler
// - 假设存在某个行为：每 n 秒执行一次，需要应用层在执行时手动添加下一次执行
// - 假设在秒钟级轮盘 slot 2 添加行为，那么预期的执行 round/counter 是 2+n
// - 假设当前正在进行 2+n 的 tick 的逻辑，由于 tick 和 add 在同一协程下阻塞执行，所以执行流程如下：
//   - round 2+n -> 2+n+1
//   - counter 2+n -> 2+n+1
//   - 执行 handler，然后添加行为，此时 round 和 counter 已经是 2+n+1
// - 由此，会导致 执行行为 和 添加行为 分布在两个不同的 round/counter 中，这两个不同的 round/counter 的差值就是最高精度的值
// - 要解决以上问题，可以提供一个 recycle 行为模式，将 执行行为 和 添加行为 包装在一起，从而使得 round/counter 没有差值

func (mtw *ModuleTimingWheel) Mounted() {
	mtw.Logger().Debug("OBSERVE: Mounted, mount series timing wheel rounds", zap.Any("rounds", mtw.rounds))

	if len(mtw.rounds) == 0 {
		// 没有指定轮盘刻度
		return
	}
	err := mtw.mountSeriesTimingWheelRound(mtw.rounds...)
	if err != nil {
		panic(err)
	}

	mtw.uidGenerator = 0
}

// mountSeriesTimingWheelRound 挂载一系列时间刻度
func (mtw *ModuleTimingWheel) mountSeriesTimingWheelRound(rounds ...*timingWheelRound) error {
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
	mtw.pointer = newTimingWheelPointer(newSeriesTimingWheelDisc(rounds...))
	return nil
}

// Start 启动时间轮
func (mtw *ModuleTimingWheel) Run() {
	mtw.Logger().Debug("OBSERVE: Run")

	mtw.pointer.start(mtw)
}

// addTickerHandler 添加时刻回调
func (mtw *ModuleTimingWheel) addTickerHandler(ths ...ITickHandler) error {
	if mtw.pointer == nil {
		return errorTimingWheelDiscNotExists
	}
	for _, th := range ths {
		mtw.uidGenerator++
		th.SetUID(mtw.uidGenerator)
	}
	return mtw.pointer.addTickerHandler(ths...)
}

func (mtw *ModuleTimingWheel) HandleEvent(event *sgs.ModuleEvent) {
	switch data := event.Data().(type) {
	case ITickHandler:
		mtw.handleTimingBehavior(data)
	default:
		// 处理 原生 ModuleEvent
		panic("default")
	}
}

func (mtw *ModuleTimingWheel) handleTimingBehavior(timingBehavior ITickHandler) {
	err := mtw.addTickerHandler(timingBehavior)
	if err != nil {
		panic(err)
	}
}
