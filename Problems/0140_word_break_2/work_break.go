package _140_word_break_2

import "strings"

func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]bool{}
	for _, word := range wordDict {
		wordSet[word] = true
	}
	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordList = append(wordList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordList = append(wordList, []string{word})
		}
		dp[index] = wordList
		return wordList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
