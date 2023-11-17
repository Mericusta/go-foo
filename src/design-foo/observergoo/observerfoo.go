package designfoo

import "fmt"

func ObserverPattern() {
	for index := 0; index != 10; index++ {
		IncreaseValue1()
	}
}

// ----------------------------------------------------------------

const (
	VALUE1 int = iota + 1
	VALUE2
	VALUE3
	VALUE4
)

type topicData struct {
	topic int
	value int
}

var runningTopicStack []topicData
var pendingTopicList []topicData
var maxStackDeep int = 3

func Report(topic int, value int) {
	td := topicData{topic: topic, value: value}
	fmt.Printf("report topic %v, value %v\n", td.topic, td.value)

	topicInStack := false
	for _, t := range runningTopicStack {
		if t.topic == td.topic {
			topicInStack = true
			fmt.Printf("topic %v already in stack\n", td.topic)
			break
		}
	}

STACK:
	if len(runningTopicStack) < maxStackDeep && !topicInStack {
		runningTopicStack = append(runningTopicStack, td)
		t := runningTopicStack[len(runningTopicStack)-1]
		switch t.topic {
		case VALUE1:
			if Value1Condition(t.value) {
				Value1Callback()
			}
		case VALUE2:
			if Value2Condition(t.value) {
				Value2Callback()
			}
		case VALUE3:
			if Value3Condition(t.value) {
				Value3TrueCallback()
			} else {
				Value3FalseCallback()
			}
		case VALUE4:
			if Value4Condition(t.value) {
				Value4TrueCallback()
			} else {
				Value4FalseCallback()
			}
		}
		if len(runningTopicStack) > 1 {
			runningTopicStack = runningTopicStack[0 : len(runningTopicStack)-1]
		} else {
			runningTopicStack = nil
		}
	} else {
		pendingTopicList = append(pendingTopicList, td)
		fmt.Printf("truncate topic %v, value %v\n", td.topic, td.value)
		fmt.Printf("pending topic: ")
		for _, t := range pendingTopicList {
			fmt.Printf("%v ", t.topic)
		}
		fmt.Println()
	}

	if len(runningTopicStack) == 0 && len(pendingTopicList) > 0 {
		td = pendingTopicList[0]
		if len(pendingTopicList) > 1 {
			pendingTopicList = pendingTopicList[1:]
		} else {
			pendingTopicList = nil
		}
		goto STACK
	}
}

// ----------------------------------------------------------------

var value1 int = 0

func IncreaseValue1() {
	value1++
	Report(VALUE1, value1)
}

var value2 int = 0

func IncreaseValue2() {
	value2++
	Report(VALUE2, value2)
}

var value3 int

func IncreaseValue3() {
	value3++
	Report(VALUE3, value3)
}

var value4 int

func IncreaseValue4() {
	value4++
	Report(VALUE4, value4)
}

// ----------------------------------------------------------------

func Value1Condition(value1 int) bool {
	return value1 <= 10
}

func Value1Callback() {
	fmt.Printf("Value 1 Callback Start\n")
	IncreaseValue2()
	fmt.Printf("Value 1 Callback Done\n")
}

func Value2Condition(value2 int) bool {
	return value2 <= 10
}

func Value2Callback() {
	fmt.Printf("Value 2 Callback Start\n")
	IncreaseValue2()
	IncreaseValue3()
	fmt.Printf("Value 2 Callback Done\n")
}

func Value3Condition(value3 int) bool {
	return value3 <= 5
}

func Value3TrueCallback() {
	fmt.Printf("Value 3 True Callback Start\n")
	IncreaseValue1()
	fmt.Printf("Value 3 True Callback Done\n")
}

func Value3FalseCallback() {
	fmt.Printf("Value 3 False Callback Start\n")
	IncreaseValue4()
	fmt.Printf("Value 3 False Callback Done\n")
}

func Value4Condition(value4 int) bool {
	return value4 <= 2
}

func Value4TrueCallback() {
	fmt.Printf("Value 4 True Callback Start\n")
	IncreaseValue1()
	fmt.Printf("Value 4 True Callback Done\n")
}

func Value4FalseCallback() {
	fmt.Printf("Value 4 False Callback Start\n")
	IncreaseValue1()
	IncreaseValue2()
	fmt.Printf("Value 4 False Callback Done\n")
}
