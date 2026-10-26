package main

import "fmt"

// DFS - Depth First Search
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// Do something with current node
	// call for left and right subtree recursively
	dfs(root.Left)
	dfs(root.Right)
}

// Preorder Traversal - node -> left -> right
// Visit the node first
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

// Inorder Traversal - left -> node -> right
// Vist left, then node, then right - BST Trees, as inorder will give sorted values
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Println(root.Val)
	inorder(root.Right)
}

// Postorder Traversal - left -> right -> node
// Visit children first , then node. "need childs results before deciding"
func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Println(root.Val)
}
