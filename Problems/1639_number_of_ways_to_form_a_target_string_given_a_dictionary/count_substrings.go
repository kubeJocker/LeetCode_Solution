package _639_number_of_ways_to_form_a_target_string_given_a_dictionary

func countSubstrings(s string, t string) int {
	n, m := len(s), len(t)
	ans := 0
	for d := 1 - m; d < n; d++ {
		i := max(d, 0)
		for k0, k1 := i-1, i-1; i < n && i-d < m; i++ {
			if s[i] != t[i-d] {
				k0 = k1
				k1 = i
			}
			ans += k1 - k0
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
