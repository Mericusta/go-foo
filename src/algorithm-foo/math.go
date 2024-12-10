package algorithmfoo

import (
	"fmt"
	"math"
)

func CalculateNumberLenByCycling(n int) int {
	l := 0
	for n != 0 {
		n /= 10
		l++
	}
	return l
}

func CalculateNumberLenByRecursive(n int) int {
	var f func(int) int
	f = func(n int) int {
		if n == 0 {
			return 0
		}
		return f(n/10) + 1
	}
	return f(n)
}

func CalculateNumberLenByLog(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func CalculateDigits(v int) {
	digits := int(math.Log10(float64(v))) + 1
	fmt.Println("v", v, "digits is", digits)
	fmt.Println("Pow10", int64(math.Pow10(digits)))

	digits1 := int64(math.Pow10(int(math.Log10(float64(16))) + 1))
	digits2 := int64(math.Pow10(int(math.Log10(float64(500))) + 1))

	fmt.Printf("digits1 = %v, digits2 = %v\n", digits1, digits2)

	ID := int64(202410241702)
	ServerID := int64(2)
	RobotIndex := int64(3)
	fmt.Printf("%v\n", ID*digits1*digits2+ServerID*digits2+RobotIndex)
}
