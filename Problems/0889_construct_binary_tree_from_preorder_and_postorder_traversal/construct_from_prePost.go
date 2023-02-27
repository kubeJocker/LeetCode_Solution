package _889_construct_binary_tree_from_preorder_and_postorder_traversal

import "Leetcode_Solution/Kit"

type TreeNode = Kit.TreeNode

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postOrderIndex := make(map[int]int, len(postorder))
	for k, v := range postorder {
		postOrderIndex[v] = k
	}
	return build(0, 0, len(preorder), preorder, postOrderIndex)
}

func build(preStart int, postStart int, N int, preOrder []int, postOrderIndex map[int]int) *TreeNode {
	if N == 0 {
		return nil
	}
	root := &TreeNode{Val: preOrder[preStart]}
	if N == 1 {
		return root
	}
	leftNode := preOrder[preStart+1]
	L := postOrderIndex[leftNode] - postStart + 1
	root.Left = build(preStart+1, postStart, L, preOrder, postOrderIndex)
	root.Right = build(preStart+L+1, postStart+L, N-1-L, preOrder, postOrderIndex)
	return root
}
