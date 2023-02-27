package _114_flatten_binary_tree_to_linked_list

import (
	"Leetcode_Solution/Kit"
)

type TreeNode = Kit.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left, root.Right = nil, left
	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = right
}
