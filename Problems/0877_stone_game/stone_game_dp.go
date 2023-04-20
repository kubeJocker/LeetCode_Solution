package _877_stone_game

type Pair struct {
	fir int
	sec int
}

func stoneGame(piles []int) int {
	n := len(piles)
	dp := make([][]Pair, n)
	for i := range dp {
		dp[i] = make([]Pair, n)
		for j := i; j < n; j++ {
			dp[i][j] = Pair{0, 0}
		}
		dp[i][i].fir = piles[i]
		dp[i][i].sec = 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			//先手选择最左边或最右边的分数
			left := piles[i] + dp[i+1][j].sec
			right := piles[j] + dp[i][j-1].sec
			//先手肯定会选择更大的结果，之后先后手转换
			if left > right {
				dp[i][j].fir = left
				dp[i][j].sec = dp[i+1][j].fir
			} else {
				dp[i][j].fir = right
				dp[i][j].sec = dp[i][j-1].fir
			}
		}
	}
	res := dp[0][n-1]
	return res.fir - res.sec
}
