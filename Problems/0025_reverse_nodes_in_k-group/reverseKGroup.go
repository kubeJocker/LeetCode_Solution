package _025_reverse_nodes_in_k_group

import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverseList(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverseList(head, last *ListNode) *ListNode {
	var pre *ListNode
	for head != last {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
