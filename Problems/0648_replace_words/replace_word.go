package _648_replace_words

import "strings"

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func replaceWords(dict []string, sentence string) string {
	trie := Constructor()
	for _, v := range dict {
		trie.Insert(v)
	}
	words := strings.Split(sentence, " ")
	var result []string
	word := ""
	i := 0
	for _, value := range words {
		word = ""
		for i = 1; i < len(value); i++ {
			if trie.Search(value[:i]) {
				word = value[:i]
				break
			}
		}

		if len(word) == 0 {
			result = append(result, value)
		} else {
			result = append(result, word)
		}

	}
	return strings.Join(result, " ")
}
