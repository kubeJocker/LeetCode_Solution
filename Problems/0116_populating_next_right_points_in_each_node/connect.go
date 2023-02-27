package _116_populating_next_right_points_in_each_node

import "Leetcode_Solution/Kit"

type Node = Kit.Node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

func traverse(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	traverse(node1.Right, node2.Left)
}
