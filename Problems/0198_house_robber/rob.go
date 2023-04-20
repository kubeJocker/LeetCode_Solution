package _198_house_robber

func rob(nums []int) int {
	n := len(nums)
	dp_i_plus_1, dp_i_plus_2 := 0, 0
	dp_i := 0
	for i := n - 1; i >= 0; i-- {
		dp_i = max(dp_i_plus_1, nums[i]+dp_i_plus_2)
		dp_i_plus_2 = dp_i_plus_1
		dp_i_plus_1 = dp_i
	}
	return dp_i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
