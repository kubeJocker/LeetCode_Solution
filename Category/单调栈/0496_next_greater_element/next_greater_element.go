package _496_next_greater_element

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	length := len(nums1)
	records, stack := map[int]int{}, []int{}
	res := make([]int, length)
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			records[num] = stack[len(stack)-1]
		} else {
			records[num] = -1
		}
		stack = append(stack, num)
	}
	for i, num := range nums1 {
		res[i] = records[num]
	}
	return res
}
