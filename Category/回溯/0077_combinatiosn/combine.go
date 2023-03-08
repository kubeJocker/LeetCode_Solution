package _077_combinatiosn

//Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
//
//You may return the answer in any order.

func subsets_DFS(nums []int) [][]int {
	track, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		generateSubsets(nums, k, 0, track, &res)
	}
	return res
}

func generateSubsets(nums []int, k, start int, track []int, res *[][]int) {
	if len(track) == k {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums)-(k-len(track))+1; i++ {
		track = append(track, nums[i])
		generateSubsets(nums, k, i+1, track, res)
		track = track[:len(track)-1]
	}
	return
}
