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
