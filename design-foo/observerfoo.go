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

var topicStack []topicData

var stackDeep int = 0
var maxStackDeep int = 2

func Report(topic int, value int) {
	fmt.Printf("topic %v, value %v\n", topic, value)

	// push
	topicStack = append(topicStack, topicData{
		topic: topic,
		value: value,
	})

	if stackDeep < maxStackDeep && len(topicStack) > 0 {
		_topic := topicStack[len(topicStack)-1] // pop
		stackDeep++
		switch _topic.topic {
		case VALUE1:
			if Value1Condition(_topic.value) {
				Value1Callback()
			}
		case VALUE2:
			if Value2Condition(_topic.value) {
				Value2Callback()
			}
		case VALUE3:
			if Value3Condition(_topic.value) {
				Value3Callback()
			}
		}
		stackDeep--
		if len(topicStack) == 1 {
			topicStack = topicStack[0:0]
		} else {
			topicStack = topicStack[0 : len(topicStack)-1]
		}
	}
}

// ----------------------------------------------------------------

var value1 int = 9

func IncreaseValue1() {
	value1++

	Report(VALUE1, value1)
}

var value2 int = 9

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
	fmt.Printf("Value 1 Callback\n")
	IncreaseValue2()
}

func Value2Condition(value2 int) bool {
	return value2 == 10
}

func Value2Callback() {
	fmt.Printf("Value 2 Callback\n")
	for index := 0; index != 10; index++ {
		IncreaseValue3()
	}
}

func Value3Condition(value3 int) bool {
	return value3 >= 10
}

func Value3Callback() {
	fmt.Printf("Value 3 Callback\n")
	IncreaseValue1()
}
