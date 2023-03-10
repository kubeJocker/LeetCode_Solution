package Kit

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func MakeListNodeList(numss [][]int) []*ListNode {
	res := []*ListNode{}

	for _, nums := range numss {
		res = append(res, MakeListNode(nums))
	}

	return res
}

func MakeListNodeListWithCycle(nums []int, pos int) *ListNode {
	head := MakeListNode(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
