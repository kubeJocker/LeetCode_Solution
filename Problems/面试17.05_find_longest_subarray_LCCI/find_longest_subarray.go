package 面试17_05_find_longest_subarray_LCCI

func findLongestSubarray(array []string) []string {
	mp := map[int]int{0: -1}
	sum, max, left := 0, 0, -1
	for i := range array {
		if array[i][0] >= '0' && array[i][0] <= '9' {
			sum++
		} else {
			sum--
		}
		if index, ok := mp[sum]; ok {
			temp := i - index
			if temp > max {
				max = temp
				left = index + 1
			}
		} else {
			mp[sum] = i
		}
	}
	if max == 0 {
		return []string{}
	}
	return array[left : left+max]
}
