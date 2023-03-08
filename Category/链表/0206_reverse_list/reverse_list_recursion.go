package _206_reverse_list

import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lastNode := reverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return lastNode
}
