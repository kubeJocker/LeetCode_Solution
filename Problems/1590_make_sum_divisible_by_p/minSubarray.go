package _590_make_sum_divisible_by_p

// 使用前缀和思想，map存前缀和的余数以及最后出现的index
// 再次遍历当前和的余数-总的余数在map中有值时为成立的一种情况
func minSubarray(nums []int, p int) int {
	hashMap := map[int]int{0: -1}
	sum, res := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	rem := sum % p
	if rem == 0 {
		return 0
	}
	sum = 0
	for i, num := range nums {
		sum += num
		tempRem := sum % p
		k := (tempRem - rem + p) % p
		if _, ok := hashMap[k]; ok {
			res = min(res, i-hashMap[k])
		}
		hashMap[tempRem] = i
	}

	if res >= len(nums) {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
