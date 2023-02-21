package problem0001

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		another := target - v
		if _, ok := numMap[another]; ok {
			return []int{numMap[another], i}
		}
		numMap[v] = i
	}
	return nil
}
