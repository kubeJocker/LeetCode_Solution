package _509_fibonacci_number

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp1, dp2 := 0, 1
	for i := 2; i <= n; i++ {
		dp := dp1 + dp2
		dp1 = dp2
		dp2 = dp
	}
	return dp2
}
