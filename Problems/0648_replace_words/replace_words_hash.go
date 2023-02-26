package _648_replace_words

import "strings"

func replaceWords_hash(dict []string, sentence string) string {
	roots := make(map[byte][]string)
	for _, root := range dict {
		b := root[0]
		roots[b] = append(roots[b], root)
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		b := []byte(word)
		for j := 1; j < len(b) && j <= 100; j++ {
			if findWord(roots, b[0:j]) {
				words[i] = string(b[0:j])
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func findWord(roots map[byte][]string, word []byte) bool {
	if roots[word[0]] == nil {
		return false
	}
	for _, root := range roots[word[0]] {
		if root == string(word) {
			return true
		}
	}
	return false
}
