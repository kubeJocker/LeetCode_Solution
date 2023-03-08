package _876

import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length%2 == 0 {
		return slow.Next
	}
	return slow
}
