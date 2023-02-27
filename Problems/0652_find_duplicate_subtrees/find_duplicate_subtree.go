package _652_find_duplicate_subtrees

import (
	"Leetcode_Solution/Kit"
	"strconv"
)

type TreeNode = Kit.TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	record := make(map[string]int)
	var traverse func(node *TreeNode) string
	traverse = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := traverse(node.Left)
		right := traverse(node.Right)

		subTree := left + "," + right + "," + strconv.Itoa(node.Val)
		freq := record[subTree]
		if freq == 1 {
			res = append(res, node)
		}
		record[subTree] += 1
		return subTree
	}
	traverse(root)
	return res
}
