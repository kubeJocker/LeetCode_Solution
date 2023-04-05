package _357_使数组元素全部相等的最少操作

import "sort"

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q*j - sum[j]
		right := sum[n] - sum[j] - q*(n-j)
		ans[i] = int64(left + right)
	}
	return ans
}
