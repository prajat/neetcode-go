package main

/*
Given the root of a binary tree, return its depth.
The depth of a binary tree is defined as the number of nodes along the longest path from the root node down to the farthest leaf node.
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLength := maxDepth(root.Left)
	rightLength := maxDepth(root.Right)
	return max(leftLength, rightLength) + 1
}
