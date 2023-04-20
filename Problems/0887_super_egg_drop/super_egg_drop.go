package _887_super_egg_drop

import "math"

//	int dp(K int, N int) {
//		res int
//	 for 1 <= i <= N {
//			res = min(res,
//					max(
//						dp(k-1, i-1) #在第i层扔鸡蛋时碎了
//						dp(k,   n-i) #在第i层扔鸡蛋没碎
//						) + 1
//			)
//		}
//		return res
//	}
func superEggDrop(k int, n int) int {
	memo := map[int]int{}
	return dp(k, n, memo)
}

func dp(k, n int, memo map[int]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if m, ok := memo[110*n+k]; ok {
		return m
	}
	res := math.MaxInt32
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		broken := dp(k-1, mid-1, memo)
		not_broken := dp(k, n-mid, memo)
		if broken > not_broken {
			hi = mid - 1
			res = min(res, broken+1)
		} else {
			lo = mid + 1
			res = min(res, not_broken+1)
		}
	}
	memo[110*n+k] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
