package _206_reverse_list

func reverseBetween(head *ListNode, m, n int) *ListNode {
	var successor *ListNode
	if m == 1 {
		return reverseN(head, n, successor)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverseN(head *ListNode, n int, successor *ListNode) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	lastNode := reverseN(head.Next, n-1, successor)
	head.Next.Next = head
	head.Next = successor
	return lastNode
}
