package _351_标记所有元素后数组的分数

import "sort"

/*
把 nums[i]\textit{nums}[i]nums[i] 及其下标绑定后，按照元素值从小到大排序，元素值相同的按照下标排序
*/
func findScore(nums []int) (ans int64) {
	type pair struct{ v, i int }
	sorted := make([]pair, len(nums))
	for i, num := range nums {
		sorted[i] = pair{num, i + 1}
	}
	sort.Slice(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		if a.v == b.v {
			return a.i < b.i
		}
		return a.v < b.v
	})
	vis := make([]bool, len(nums)+2)
	for _, p := range sorted {
		if !vis[p.i] {
			vis[p.i-1] = true
			vis[p.i+1] = true
			ans += int64(p.v)
		}
	}
	return
}
