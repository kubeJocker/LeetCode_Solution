package _488H_统计中位数为k的数组

func countSubarrays(nums []int, k int) int {
	kIndex := -1
	for i, num := range nums {
		if num == k {
			kIndex = i
			break
		}
	}
	prefixMap, sum, ans := make(map[int]int), 0, 0
	prefixMap[0] = 1
	for i, num := range nums {
		if num < k {
			sum--
		}
		if num > k {
			sum++
		}
		if i < kIndex {
			prefixMap[sum]++
			continue
		}
		ans += prefixMap[sum] + prefixMap[sum-1]
	}
	return ans
}
