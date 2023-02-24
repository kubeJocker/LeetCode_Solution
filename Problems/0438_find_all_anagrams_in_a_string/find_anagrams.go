package _438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	window, need := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	res := []int{}
	for right < len(s) {
		chRight := s[right]
		right++
		if need[chRight] > 0 {
			window[chRight]++
			if window[chRight] == need[chRight] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			chLeft := s[left]
			left++
			if need[chLeft] > 0 {
				if window[chLeft] == need[chLeft] {
					valid--
				}
				window[chLeft]--
			}
		}
	}
	return res
}
