package _003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	window := map[byte]int{}
	for right < len(s) {
		if index, ok := window[s[right]]; ok && index >= left {
			left = index + 1
		}
		window[s[right]] = right
		right++
		res = max(res, right-left)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
