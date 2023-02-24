package _104_maximum_depth_of_binary_tree

import "Leetcode_Solution/Kit"

type TreeNode = Kit.TreeNode

func maxDepth(root *TreeNode) int {
	return terverse(root)
}

func terverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := terverse(root.Left)
	right := terverse(root.Right)
	return max(left, right) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
