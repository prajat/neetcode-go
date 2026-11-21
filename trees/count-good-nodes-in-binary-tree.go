package main

/*
Within a binary tree, a node x is considered good if the path from the root of the tree to the node x contains no nodes with a value greater than the value of node x
Given the root of a binary tree root, return the number of good nodes within the tree.
*/
func goodNodes(root *TreeNode) int {
	count := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, maxValue int) {
		if node == nil {
			return
		}

		if node.Val >= maxValue {
			count++
		}

		dfs(node.Left, max(maxValue, node.Val))
		dfs(node.Right, max(maxValue, node.Val))
	}
	dfs(root, root.Val)
	return count
}
