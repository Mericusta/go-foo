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
)

type topicData struct {
	topic int
	value int
}

var runningTopicStack []topicData
var pendingTopicStack []topicData
var topicStackSlice []topicData

var stackDeep int = 0
var maxStackDeep int = 1

func Report(topic int, value int) {
	td := topicData{topic: topic, value: value}
	fmt.Printf("topic %v, value %v\n", td.topic, td.value)

	if stackDeep < maxStackDeep {
		stackDeep++
		switch td.topic {
		case VALUE1:
			if Value1Condition(td.value) {
				Value1Callback()
			}
		case VALUE2:
			if Value2Condition(td.value) {
				Value2Callback()
			}
		case VALUE3:
			if Value3Condition(td.value) {
				Value3Callback()
			}
		}
		stackDeep--
	} else {
		topicStackSlice = append(topicStackSlice, td)
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
	IncreaseValue3()
	fmt.Printf("Value 2 Callback Done\n")
}

func Value3Condition(value3 int) bool {
	return value3 <= 10
}

func Value3Callback() {
	fmt.Printf("Value 3 Callback Start\n")
	IncreaseValue1()
	fmt.Printf("Value 3 Callback Done\n")
}
