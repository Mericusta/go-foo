package algorithmfoo

import "math"

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
