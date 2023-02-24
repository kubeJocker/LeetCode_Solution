package _543_diameter_of_binary_tree

import (
	"Leetcode_Solution/Kit"
)

type TreeNode = Kit.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	terverse(root, &res)
	return res
}

func terverse(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := terverse(root.Left, res)
	right := terverse(root.Right, res)
	*res = max(left+right, *res)
	return max(left, right) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
