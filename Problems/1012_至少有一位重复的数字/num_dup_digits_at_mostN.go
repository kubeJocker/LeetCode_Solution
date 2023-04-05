package _012_至少有一位重复的数字

import "strconv"

/*
数位DP
*/
func numDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum {
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return n - f(0, 0, true, false)
}
