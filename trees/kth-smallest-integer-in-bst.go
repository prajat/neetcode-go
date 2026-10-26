package main

/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) in the tree.
A binary search tree satisfies the following constraints:
The left subtree of every node contains only nodes with keys less than the node's key.
The right subtree of every node contains only nodes with keys greater than the node's key.
Both the left and right subtrees are also binary search trees.
*/

// Not Optimal Solution - Used inorder to convert into sorted array and then returned the answer
func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	inorderArr(root, &arr)
	return arr[k-1]
}

func inorderArr(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorderArr(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorderArr(root.Right, arr)
}

// Optimal Solution - Using DFS
func kthSmallestV2(root *TreeNode, k int) int {
	result := 0
	count := 0

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count > k {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}
