package _046_permutations

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, track, result := make([]bool, len(nums)), []int{}, [][]int{}
	backtrace(nums, track, &result, &used)
	return result
}

func backtrace(nums []int, track []int, res *[][]int, used *[]bool) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	for i := range nums {
		if !(*used)[i] {
			(*used)[i] = true
			track = append(track, nums[i])
			backtrace(nums, track, res, used)
			track = track[:len(track)-1]
			(*used)[i] = false
		}
	}
	return
}
