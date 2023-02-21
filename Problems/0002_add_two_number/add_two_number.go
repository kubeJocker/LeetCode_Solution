package problem0002

import (
	. "Leetcode_Solution/Kit"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry, curr := 0, dummy
	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry

		carry = sum / 10
		sum %= 10
		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
	}
	return dummy.Next
}
