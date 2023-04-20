package _787_cheapest_flight_within_k_stops

// f[t][i]f表示通过恰好 t次航班从出发城市 src 到达城市 i 需要的最小花费
// f[t][i]= (j,i)∈flights min {f[t−1][j]+cost(j,i)}
func findCheapestPrice_dp(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	f := make([][]int, k+2)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			f[t][i] = min(f[t][i], f[t-1][j]+cost)
		}
	}
	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, f[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
