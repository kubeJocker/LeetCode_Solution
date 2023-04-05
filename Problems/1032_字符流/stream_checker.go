package _032_字符流

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type StreamChecker struct {
	root     *TrieNode
	queryStr []byte
}

func Constructor(words []string) StreamChecker {
	streamChecker := StreamChecker{&TrieNode{}, []byte{}}
	for _, word := range words {
		streamChecker.Insert(word)
	}
	return streamChecker
}

func (this *StreamChecker) Query(letter byte) bool {
	this.queryStr = append(this.queryStr, letter)
	return this.Find(this.queryStr)
}

func (this *StreamChecker) Insert(word string) {
	r := this.root
	for i := len(word) - 1; i >= 0; i-- {
		nextIdx := int(word[i] - 'a')
		if r.children[nextIdx] == nil {
			r.children[nextIdx] = &TrieNode{}
		}
		r = r.children[nextIdx]
	}
	r.isEnd = true
}

func (this *StreamChecker) Find(str []byte) bool {
	r := this.root
	for i := len(str) - 1; i >= 0; i-- {
		if r.isEnd {
			return true
		}
		nextIdx := int(str[i] - 'a')
		if r.children[nextIdx] == nil {
			return false
		}
		r = r.children[nextIdx]
	}
	return r.isEnd
}
