package LFU_linkedList

import (
	"container/list"
)

type LFUCache struct {
	keyToVal     map[int]*list.Element
	freqToKey    map[int]*list.List
	cap, minFreq int
}

type Node struct {
	key  int
	val  int
	freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyToVal:  make(map[int]*list.Element),
		freqToKey: make(map[int]*list.List),
		cap:       capacity,
		minFreq:   0,
	}
}

func (this *LFUCache) Get(key int) int {
	x, ok := this.keyToVal[key]
	if !ok {
		return -1
	}
	node := x.Value.(Node)
	oldFreq := node.freq
	node.freq++

	preList := this.freqToKey[oldFreq]
	preList.Remove(x)

	var keyList *list.List
	if keyList, ok = this.freqToKey[node.freq]; !ok {
		keyList = list.New()
	}
	this.keyToVal[key] = keyList.PushFront(node)
	this.freqToKey[node.freq] = keyList
	if preList.Len() == 0 {
		if this.minFreq == oldFreq {
			this.minFreq++
		}
	}
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if x, ok := this.keyToVal[key]; ok {
		node := x.Value.(Node)
		node.val = value
		oldFreq := node.freq
		node.freq++

		preList := this.freqToKey[oldFreq]
		preList.Remove(x)

		var nodeList *list.List
		if nodeList, ok := this.freqToKey[oldFreq+1]; !ok {
			nodeList = list.New()
			this.freqToKey[oldFreq+1] = nodeList
		}
		this.keyToVal[key] = nodeList.PushFront(node)
		if preList.Len() == 0 {
			if oldFreq == this.minFreq {
				this.minFreq++
			}
		}
		return
	}

	if this.cap == len(this.keyToVal) {
		minList := this.freqToKey[this.minFreq]
		min := minList.Back()
		minList.Remove(min)
		node := min.Value.(Node)
		delete(this.keyToVal, node.key)
	}
	newNode := Node{key, value, 1}
	this.minFreq = newNode.freq
	if keyList, ok := this.freqToKey[this.minFreq]; !ok {
		keyList = list.New()
	} else {
		this.keyToVal[key] = keyList.PushFront(newNode)
		this.freqToKey[this.minFreq] = keyList
	}
}
