package problem0086

import (
	"Leetcode_Solution/Kit"
)

type ListNode = Kit.ListNode

func partition(head *ListNode, x int) *ListNode {
	dummyLessThanX, dummyMoreThanX := &ListNode{}, &ListNode{}
	p1, p2 := dummyLessThanX, dummyMoreThanX
	curr := head

	for curr != nil {
		if curr.Val < x {
			p1.Next = curr
			p1 = p1.Next
		} else {
			p2.Next = curr
			p2 = p2.Next
		}
		tempNode := curr.Next
		curr.Next = nil
		curr = tempNode
	}

	p1.Next = dummyMoreThanX.Next
	return dummyLessThanX.Next
}
