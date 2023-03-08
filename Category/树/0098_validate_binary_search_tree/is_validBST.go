package _098_validate_binary_search_tree

import (
	"Leetcode_Solution/Kit"
	"math"
)

type TreeNode = Kit.TreeNode

func isValidBST(root *TreeNode) bool {
	return traverse(root, math.MinInt64, math.MaxInt64)
}

func traverse(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return traverse(root.Left, lower, root.Val) && traverse(root.Right, root.Val, upper)
}
