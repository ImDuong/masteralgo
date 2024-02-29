package dynamicprogramming

// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// 0 <= n <= 30
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, 31)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	dp[1] = 1
	return calFib(n, dp)
}

func calFib(n int, dp []int) int {
	var first, second int
	if dp[n-1] != -1 {
		first = dp[n-1]
	} else {
		first = calFib(n-1, dp)
	}

	if dp[n-2] != -1 {
		second = dp[n-2]
	} else {
		second = calFib(n-2, dp)
	}

	dp[n] = first + second
	return dp[n]
}
