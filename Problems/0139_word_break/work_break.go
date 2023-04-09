package _139_word_break

// dp[i]表示字符串前i个字符是否能被拆分为字典中的单词
// dp[i]=dp[j] && check(s[j..i−1])
// 其中 check(s[j..i−1]) 表示子串 s[j..i−1] 是否出现在字典中。
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
