package problem0141

import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
