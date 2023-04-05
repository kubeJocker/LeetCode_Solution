package _352_美丽子集的数目

func beautifulSubsets(nums []int, k int) int {
	ans := -1 // 去掉空集
	cnt := make([]int, 1001+k*2)
	var dfs func(int)
	dfs = func(i int) { // 从 i 开始选
		ans++ // 合法子集
		if i == len(nums) {
			return
		}
		for j := i; j < len(nums); j++ { // 枚举选哪个
			x := nums[j] + k // 避免负数下标
			if cnt[x-k] == 0 && cnt[x+k] == 0 {
				cnt[x]++ // 选
				dfs(j + 1)
				cnt[x]-- // 恢复现场
			}
		}
	}
	dfs(0)
	return ans
}
