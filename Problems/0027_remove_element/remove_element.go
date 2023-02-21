package problem0027

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	fast, slow := 0, 0
	for fast < length {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
