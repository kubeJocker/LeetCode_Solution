package 面试61_扑克牌中的顺子

func isStraight(nums []int) bool {
	hashMap := map[int]bool{}
	max, min := 0, 14
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if _, ok := hashMap[num]; ok {
			return false
		}
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		hashMap[num] = true
	}
	return max-min < 5
}
