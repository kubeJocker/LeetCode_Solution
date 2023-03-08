package _236_lowest_common_ancestor_of_a_binary_tree

import (
	"Leetcode_Solution/Kit"
)

type TreeNode = Kit.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, p.Val, q.Val)
}

func find(root *TreeNode, val1 int, val2 int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val1 || root.Val == val2 {
		return root
	}
	left := find(root.Left, val1, val2)
	right := find(root.Right, val1, val2)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
