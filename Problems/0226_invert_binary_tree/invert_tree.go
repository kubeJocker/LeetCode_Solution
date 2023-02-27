package _226_invert_binary_tree

import (
	"Leetcode_Solution/Kit"
)

type TreeNode = Kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	traverse(root.Left)
	traverse(root.Right)
}
