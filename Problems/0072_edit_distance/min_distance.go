package _072_edit_distance

// if s1[i] == s2[j]:
//
//	啥都别做（skip）
//	i, j 同时向前移动
//
// else:
//
//	三选一：
//	    插入（insert）
//	    删除（delete）
//	    替换（replace）
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 定义 word1[0..i] 和 word2[0..j] 的最小编辑距离是 dp[i+1][j+1]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)

	}
	// base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	// 自底向上求解
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,
					dp[i][j-1]+1,
					dp[i-1][j-1]+1,
				)
			}
		}
	}
	// 储存着整个 word1 和 word2 的最小编辑距离
	return dp[m][n]
}

func min(a, b, c int) int {
	minN := a
	if minN > b {
		minN = b
	}
	if minN > c {
		minN = c
	}
	return minN
}
