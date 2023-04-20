package _518_coin_change_2

//for (int i = 1; i <= n; i++) {
//    for (int j = 1; j <= amount; j++) {
//        if (j - coins[i-1] >= 0)
//            dp[i][j] = dp[i - 1][j]
//                     + dp[i][j-coins[i-1]];
//return dp[N][W]
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1 // base case
	for i := 0; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
