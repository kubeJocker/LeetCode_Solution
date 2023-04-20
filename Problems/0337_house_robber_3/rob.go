package _337_house_robber_3

import "Leetcode_Solution/Kit"

type TreeNode = Kit.TreeNode

func rob(root *TreeNode) int {
	res := dp(root)
	return max(res[0], res[1])
}

func dp(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dp(root.Left)
	right := dp(root.Right)
	//抢这家
	rob_this := root.Val + left[0] + right[0]
	not_rob := max(left[0], left[1]) + max(right[0], right[1])
	return []int{not_rob, rob_this}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}