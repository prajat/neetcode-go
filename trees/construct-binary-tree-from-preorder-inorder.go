package main

/*
You are given two integer arrays preorder and inorder.
preorder is the preorder traversal of a binary tree
inorder is the inorder traversal of the same tree
Both arrays are of the same size and consist of unique values.
Rebuild the binary tree from the preorder and inorder traversals and return its root.
*/
func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// make hashSet from inorder arr to store index
	inorderIndex := make(map[int]int, len(inorder))
	for i := range inorder {
		inorderIndex[inorder[i]] = i
	}

	// first element of preorder gives us the root node
	// the root node indexed gives us two half subtree nodes in inorder arr.
	preorderIndex := 0

	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// The next preorder value is the root
		rootValue := preorder[preorderIndex]
		preorderIndex++

		root := &TreeNode{Val: rootValue}

		// find root's position in inorder
		mid := inorderIndex[rootValue]

		// build left subtree
		root.Left = build(left, mid-1)

		// build right subtree
		root.Right = build(mid+1, right)

		return root
	}
	return build(0, len(inorder)-1)
}
