package algorithmfoo

func my_fib(n int) int {
	if n <= 2 {
		return n
	}
	return my_fib(n-1) + my_fib(n-2)
}

func my_fib_cycle(n int) int {
	if n <= 1 {
		return n
	}

	var a, b int = 0, 1
	for index := 1; index <= n; index++ {
		a, b = b, a+b
	}

	return b
}
