package _654_maximum_binary_tree

import (
	"Leetcode_Solution/Kit"
)

type TreeNode = Kit.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	best := 0
	for i, num := range nums {
		if num > nums[best] {
			best = i
		}
	}
	return &TreeNode{nums[best], constructMaximumBinaryTree(nums[:best]), constructMaximumBinaryTree(nums[best+1:])}
}
