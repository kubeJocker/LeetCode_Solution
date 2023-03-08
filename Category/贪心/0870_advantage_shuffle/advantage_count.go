package _870_advantage_shuffle

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	length := len(nums1)
	sort.Ints(nums1)
	sortedNums2 := make([]int, length)
	for i := range sortedNums2 {
		sortedNums2[i] = i
	}
	sort.Slice(sortedNums2, func(i, j int) bool {
		return nums2[sortedNums2[i]] < nums2[sortedNums2[j]]
	})
	i, useless, res := 0, make([]int, 0), make([]int, length)
	for _, index := range sortedNums2 {
		b := nums2[index]
		for i < length && nums1[i] < b {
			useless = append(useless, nums1[index])
			i++
		}
		if i < length {
			res[index] = nums1[i]
			i++
		} else {
			res[index] = useless[0]
			useless = useless[1:]
		}
	}
	return res
}
