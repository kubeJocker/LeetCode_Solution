package problem0005

/*
遍历每个字符，并向两端扩散，查找是否是回文(同时考虑奇数与偶数情况)
*/

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}
