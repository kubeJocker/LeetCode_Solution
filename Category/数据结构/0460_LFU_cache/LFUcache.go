package _460_LFU_cache

import "container/heap"

/*
使用堆来维护kv以及操作频率，使用hash表来定位节点中的索引
节点同时保存在堆中的位置索引，以便heap在堆整理时使用
缓存替换时，堆弹出操作频率最小的节点并执行堆整理
*/
type LFUCache struct {
	nheap     *NodeHeap
	nodeMap   map[int]*Node
	size, cap int
	opcnt     int
}

type Node struct {
	freq, lastop int
	key, val     int
	index        int
}
type NodeHeap []*Node

func (nheap NodeHeap) Len() int {
	return len(nheap)
}
func (nheap NodeHeap) Swap(i, j int) {
	nheap[i], nheap[j] = nheap[j], nheap[i]
	nheap[i].index, nheap[j].index = nheap[j].index, nheap[i].index
}

func (nheap NodeHeap) Less(i, j int) bool {
	if nheap[i].freq < nheap[j].freq {
		return true
	} else if nheap[i].freq > nheap[j].freq {
		return false
	} else {
		return nheap[i].lastop < nheap[j].lastop
	}
}

func (nheap *NodeHeap) Push(x interface{}) {
	*nheap = append(*nheap, x.(*Node))
}

func (nheap *NodeHeap) Pop() interface{} {
	x := (*nheap)[len(*nheap)-1]
	*nheap = (*nheap)[:len(*nheap)-1]
	return x
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{&NodeHeap{}, map[int]*Node{}, 0, capacity, 0}
	heap.Init(cache.nheap)
	return cache
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	if node, ok := this.nodeMap[key]; !ok {
		return -1
	} else {
		x := node.val
		node.freq++
		node.lastop = this.opcnt
		this.opcnt++
		heap.Fix(this.nheap, node.index)
		return x
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if node, ok := this.nodeMap[key]; ok {
		node.freq++
		node.lastop = this.opcnt
		this.opcnt++
		node.val = value
		heap.Fix(this.nheap, node.index)
	} else {
		if this.size == this.cap {
			kicked := heap.Pop(this.nheap)
			delete(this.nodeMap, kicked.(*Node).key)
		} else {
			this.size++
		}
		newNode := &Node{1, this.opcnt, key, value, this.nheap.Len()}
		this.opcnt++
		heap.Push(this.nheap, newNode)
		this.nodeMap[key] = newNode
	}
}
