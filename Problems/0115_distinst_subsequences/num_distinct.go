package _115_distinst_subsequences

// 定义：s[i:] 的子序列中 t[j:] 出现的次数为 dp(s, i, t, j)
//
//	func dp(s string, i int, t string, j int) int {
//	   if s[i] == t[j] {
//	       // 匹配，两种情况，累加关系
//	       return dp(s, i + 1, t, j + 1) + dp(s, i + 1, t, j)
//	   } else {
//	       // 不匹配，在 s[i+1:] 的子序列中计算 t[j:] 的出现次数
//	       return dp(s, i + 1, t, j)
//	   }
//	}
func numDistinct(s string, t string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(t))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(s, 0, t, 0, memo)
}

func dp(s string, i int, t string, j int, memo [][]int) int {
	if j == len(t) {
		return 1
	}
	if len(s)-i < len(t)-j {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	res := 0
	if s[i] == t[j] {
		// 匹配，两种情况，累加关系
		res += dp(s, i+1, t, j+1, memo) + dp(s, i+1, t, j, memo)
	} else {
		// 不匹配，在 s[i+1..] 的子序列中计算 t[j..] 的出现次数
		res += dp(s, i+1, t, j, memo)
	}
	memo[i][j] = res
	return res
}
