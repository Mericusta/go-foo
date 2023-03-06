package algorithmfoo

import (
	"fmt"
	"time"
)

var TestDuration time.Duration
var MinScale time.Duration
var Ticker *time.Ticker
var TaskTimer *time.Timer
var ExitTimer *time.Timer
var ExitChan chan int64

type TimeWheelTask struct {
	Delay time.Duration
	GUID  int
	F     func()
}

type TimeWheel struct {
	GUID             int
	Scale            time.Duration
	SlotSize         int
	CurrentSlotIndex int
	ParentTimeWheel  *TimeWheel
	TaskListSlice    [][]*TimeWheelTask
}

func (tw *TimeWheel) Run() {
	fmt.Printf("Timewheel %v run at slot index %v\n", tw.GUID, tw.CurrentSlotIndex)
	for index, task := range tw.TaskListSlice[tw.CurrentSlotIndex] {
		if task != nil {
			task.F()
		}
		tw.TaskListSlice[tw.CurrentSlotIndex][index] = nil
	}
	tw.TaskListSlice[tw.CurrentSlotIndex] = nil
}

func (tw *TimeWheel) Tick() {
	fmt.Printf("Timewheel %v tw.CurrentSlotIndex = %v, tw.SlotSize = %v, (tw.CurrentSlotIndex) MOD tw.SlotSize = %v\n", tw.GUID, tw.CurrentSlotIndex, tw.SlotSize, (tw.CurrentSlotIndex)%tw.SlotSize)
	if tw.CurrentSlotIndex%tw.SlotSize == tw.SlotSize-1 {
		fmt.Printf("Timewheel %v Round End\n", tw.GUID)
		tw.CurrentSlotIndex = 0
		if tw.ParentTimeWheel != nil {
			tw.ParentTimeWheel.Tick()
			for _, taskInParentTimeWheel := range tw.ParentTimeWheel.TaskListSlice[tw.ParentTimeWheel.CurrentSlotIndex] {
				fmt.Printf("Timewheel %v add task %v from parent timewheel %v\n", tw.GUID, taskInParentTimeWheel.GUID, tw.ParentTimeWheel.GUID)
				tw.AddTask(taskInParentTimeWheel.Delay, taskInParentTimeWheel.F, taskInParentTimeWheel.GUID)
			}
		}
	} else {
		tw.CurrentSlotIndex++
	}
}

// 任务延迟时间所对轮数 = ((延迟时间 + 当前轮已经过时长) / 刻度 * 槽数) 取整
// 任务延迟时间所对下标 = (当前轮已经过时长 + (延迟时间 / 刻度) 取整) % 槽数
// 任务延迟时间指的是定时器达到延迟指定的时刻后，立刻开始执行任务，下标需要-1
// 任务延迟时间无法被最小轮刻度整除的，在下一个时刻执行
func (tw *TimeWheel) AddTask(delay time.Duration, op func(), opGUID int) {
	fmt.Printf("timewheel %v current slot index %v add task %v with delay seconds %v\n", tw.GUID, tw.CurrentSlotIndex, opGUID, delay.Seconds())
	round := int((delay + tw.Scale*time.Duration(tw.CurrentSlotIndex)) / (tw.Scale * time.Duration(tw.SlotSize)))
	if round == 0 {
		index := (tw.CurrentSlotIndex+int(delay/tw.Scale))%tw.SlotSize - 1
		if index < 0 {
			index = 0
		}
		fmt.Printf("timewheel %v add task %v at round %v index %v\n", tw.GUID, opGUID, round, index)
		tw.TaskListSlice[index] = append(tw.TaskListSlice[index], &TimeWheelTask{
			Delay: delay,
			GUID:  opGUID,
			F:     op,
		})
	} else {
		if tw.ParentTimeWheel != nil {
			fmt.Printf("timewheel %v add task %v at parent timewheel %v\n", tw.GUID, opGUID, tw.ParentTimeWheel.GUID)
			tw.ParentTimeWheel.AddTask(delay-time.Duration(round)*tw.Scale*time.Duration(tw.SlotSize), op, opGUID)
		} else {
			fmt.Printf("task %v is out of time range\n", opGUID)
		}
	}
	fmt.Printf("\n")
}

var tw *TimeWheel

func TickerRun() {
	for {
		select {
		// 添加任务必须在改变 Tick 之前添加
		case <-TaskTimer.C:
			fmt.Printf("\nTask Timer Add Task\n")
			tw.AddTask(2*time.Second, func() { fmt.Printf("Task 4\n") }, 4)
		case <-Ticker.C:
			fmt.Printf("\nTick On Slot Index %v, Second %v\n", tw.CurrentSlotIndex, tw.CurrentSlotIndex+1)
			tw.Run()
			tw.Tick()
		case <-ExitTimer.C:
			fmt.Printf("\nTime's up\n")
			ExitChan <- time.Now().Unix()
			return
		}
	}
}

func timewheelFoo() {
	MinScale = time.Second
	tw = &TimeWheel{
		GUID:     1,
		Scale:    MinScale,
		SlotSize: 5,
	}
	tw.TaskListSlice = make([][]*TimeWheelTask, tw.SlotSize)

	tw.ParentTimeWheel = &TimeWheel{
		GUID:     2,
		Scale:    tw.Scale * time.Duration(tw.SlotSize),
		SlotSize: 6,
	}
	tw.ParentTimeWheel.TaskListSlice = make([][]*TimeWheelTask, tw.SlotSize)

	tw.AddTask(1*time.Second, func() { fmt.Printf("Execute Task 1\n") }, 1) // 0 - 0
	tw.AddTask(3*time.Second, func() { fmt.Printf("Execute Task 2\n") }, 2) // 0 - 2
	tw.AddTask(9*time.Second, func() { fmt.Printf("Execute Task 3\n") }, 3) // 1 - 0 -> 0 - 3

	Ticker = time.NewTicker(tw.Scale)
	TestDuration = MinScale * 10
	ExitTimer = time.NewTimer(TestDuration)
	TaskTimer = time.NewTimer(time.Second * 4)
	ExitChan = make(chan int64)

	go TickerRun()

	exitUnix := <-ExitChan
	fmt.Printf("exit at unix %v\n", time.Unix(exitUnix, 0))
}
