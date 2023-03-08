package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import "Leetcode_Solution/Kit"

type TreeNode = Kit.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	index := 0
	for ; index < len(preorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
