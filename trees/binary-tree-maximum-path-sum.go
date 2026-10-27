package main

/*
Given the root of a non-empty binary tree, return the maximum path sum of any non-empty path.
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes has an edge connecting them. A node can not appear in the sequence more than once. The path does not necessarily need to include the root.
The path sum of a path is the sum of the node's values in the path.
*/

/*
At each node, calculate:
The best path going through the node and possibly one child — returned to its parent.
The best path passing through both children — used to update the global answer.
*/
func maxPathSum(root *TreeNode) int {
	result := root.Val

	// using dfs to traverse the tree
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// ignore negative paths as they reduce the total sum
		leftGain := max(dfs(node.Left), 0)
		rightGain := max(dfs(node.Right), 0)

		// a path passing through this node can use both children
		currentPathSum := node.Val + leftGain + rightGain

		result = max(result, currentPathSum)

		// a parent can use only one branch
		return node.Val + max(leftGain, rightGain)
	}

	dfs(root)
	return result
}
