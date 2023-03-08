package _567_permutation_in_string

/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
*/

func checkInclusion(s1 string, s2 string) bool {
	window, need := map[byte]int{}, map[byte]int{}
	for i := range s1 {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		chRight := s2[right]
		right++
		if need[chRight] > 0 {
			window[chRight]++
			if window[chRight] == need[chRight] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			chLeft := s2[left]
			left++
			if need[chLeft] > 0 {
				if window[chLeft] == need[chLeft] {
					valid--
				}
				window[chLeft]--
			}
		}
	}
	return false
}
