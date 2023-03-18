package _380M_链表随机节点

import "math/rand"
import "Leetcode_Solution/Kit"

type ListNode = Kit.ListNode

// 蓄水池抽样法
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	ans := 0
	for node, i := this.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 { // 1/i 的概率选中（替换为答案）
			ans = node.Val
		}
		i++
	}
	return ans
}
