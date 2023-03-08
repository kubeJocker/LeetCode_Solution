package problem0167

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	if length == 0 {
		return numbers
	}
	left, right := 0, length-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return nil
}
