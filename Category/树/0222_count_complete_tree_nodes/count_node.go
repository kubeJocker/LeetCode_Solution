package _222_count_complete_tree_nodes

import (
	"Leetcode_Solution/Kit"
	"math"
)

type TreeNode = Kit.TreeNode

func countNodes(root *TreeNode) int {
	hl, hr := 0, 0
	left, right := root, root
	for left != nil {
		left = left.Left
		hl++
	}
	for right != nil {
		right = right.Right
		hr++
	}
	if hl == hr {
		return (int)(math.Pow(2, float64(hl))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
