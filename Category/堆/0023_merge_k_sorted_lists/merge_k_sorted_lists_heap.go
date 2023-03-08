package problem0023

import (
	. "Leetcode_Solution/Kit"
	"container/heap"
)

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKListsByQueue(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	mheap := new(minHeap)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(mheap, lists[i])
		}
	}

	for mheap.Len() > 0 {
		temp := heap.Pop(mheap).(*ListNode)
		if temp.Next != nil {
			heap.Push(mheap, temp.Next)
		}
		curr.Next = temp
		curr = curr.Next
	}
	return dummy.Next
}
