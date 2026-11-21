package main

/*
Given a binary tree, return true if it is height-balanced and false otherwise.
A height-balanced binary tree is defined as a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	isBalanced := true

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := dfs(root.Left)
		rightHeight := dfs(root.Right)

		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			isBalanced = false
		}

		return 1 + max(leftHeight, rightHeight)
	}
	_ = dfs(root)
	return isBalanced
}
