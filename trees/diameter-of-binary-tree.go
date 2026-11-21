package main

/*
The diameter of a binary tree is defined as the length of the longest path between any two nodes within the tree. The path does not necessarily have to pass through the root.
The length of a path between two nodes in a binary tree is the number of edges between the nodes. Note that the path can not include the same node twice.
Given the root of a binary tree root, return the diameter of the tree.
*/
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		maxDiameter = max(maxDiameter, leftHeight+rightHeight)

		return 1 + max(leftHeight, rightHeight)
	}
	_ = dfs(root)
	return maxDiameter
}
