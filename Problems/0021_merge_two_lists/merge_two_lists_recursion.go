package problem0021

import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

func mergeTwoListsByRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsByRecursion(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoListsByRecursion(list1, list2.Next)
	return list2
}
