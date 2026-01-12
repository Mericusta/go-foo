package algorithmfoo

// 传统的 fibonacci 实现
func FibFoo(n int) int {
	if n <= 1 {
		return n
	}
	return FibFoo(n-1) + FibFoo(n-2)
}

// 优化的 fibonacci 实现，使用记忆化技术
func FibMemoryFoo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = FibMemoryFoo(n-1, memo) + FibMemoryFoo(n-2, memo)
	return memo[n]
}

// 使用动态规划实现 fibonacci
func FibDPFoo(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 使用 for 循环实现
func FibForFoo(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}