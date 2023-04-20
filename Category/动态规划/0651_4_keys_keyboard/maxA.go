package _651_4_keys_keyboard

// 定义：dp[i] 表示 i 次操作后最多能显示多少个 A
//for (int i = 0; i <= N; i++)
//    dp[i] = max(
//            这次按 A 键，
//            这次按 C-V
//        )
func MaxA(N int) int {
	dp := make([]int, N+1)
	dp[0] = 0
	for i := 1; i <= N; i++ {
		//按 A 键
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			//全选 & 复制 dp[j-2]，连续粘贴 i - j 次
			//屏幕上共 dp[j - 2] * (i - j + 1) 个 A
			dp[i] = Max(dp[i], dp[j-2]*(i-j+1))
		}
	}
	//N 次按键之后最多有几个 A？
	return dp[N]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
