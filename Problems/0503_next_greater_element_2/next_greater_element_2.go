package _503_next_greater_element_2

func nextGreaterElements(nums []int) []int {
	res := make([]int, 0)
	indexes := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, -1)
	}
	for i := 0; i < len(nums)*2; i++ {
		num := nums[i%len(nums)]
		for len(indexes) > 0 && nums[indexes[len(indexes)-1]] < num {
			index := indexes[len(indexes)-1]
			res[index] = num
			indexes = indexes[:len(indexes)-1]
		}
		indexes = append(indexes, i%len(nums))
	}
	return res
}
