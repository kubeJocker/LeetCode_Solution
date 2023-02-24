package _078_subsets

func subsets_backtrack(nums []int) [][]int {
	res, track := [][]int{}, []int{}
	backtrack(nums, 0, &res, track)
	return res
}

func backtrack(nums []int, start int, res *[][]int, track []int) {
	*res = append(*res, track)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, res, track)
		track = track[:len(track)-1]
	}
}
