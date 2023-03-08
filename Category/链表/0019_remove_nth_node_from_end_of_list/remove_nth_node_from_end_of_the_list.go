package problem0019

import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	curr, slow, fast := dummy, head, head
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	for fast != nil {
		curr = slow
		slow = slow.Next
		fast = fast.Next
	}
	curr.Next = slow.Next
	return dummy.Next
}
