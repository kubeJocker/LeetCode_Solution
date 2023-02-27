package _538_convert_BST_to_greater_tree

import "Leetcode_Solution/Kit"

type TreeNode = Kit.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}
	traverse(root)
	return root
}
