package _213_house_robber_2

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
	dp_i_plus_1, dp_i_plus_2 := 0, 0
	dp_i := 0
	for i := end; i >= start; i-- {
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
