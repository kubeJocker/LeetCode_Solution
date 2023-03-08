package _090_subsets_2

import "sort"

func subsetsWithDup(nums []int) [][]int {
	track, res := []int{}, [][]int{}
	sort.Ints(nums)
	backtrack(nums, 0, track, &res)
	return res
}

func backtrack(nums []int, start int, track []int, res *[][]int) {

	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
	return
}
