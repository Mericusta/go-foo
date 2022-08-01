package builtinfoo

func NumberEvenOddCheck(n int) bool {
	return n&1 == 1
}

func ZoomInAndOutInMultiplesOf2(n int, in bool) int {
	if in {
		return n >> 1
	}
	return n << 1
}

func ZoomOutInMultiplesOf10(n int, origin bool) int {
	if origin {
		n = n * 10
	} else {
		n = (n << 1) + (n << 3)
	}
	return n
}
