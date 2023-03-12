package _317M_count_the_number_of_beautiful_subarrays

/*
采用前缀异或和的方式计算
*/
func beautifulSubarrays(nums []int) (ans int64) {
	s := 0
	cnt := map[int]int{s: 1} // s[0]
	for _, x := range nums {
		s ^= x
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
