package main

/*
Given the roots of two binary trees p and q, return true if the trees are equivalent, otherwise return false.
Two binary trees are considered equivalent if they share the exact same structure and the nodes have the same values.
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both nodes are empty
	if p == nil && q == nil {
		return true
	}

	// if any one node is empty
	if p == nil || q == nil {
		return false
	}

	// Values of current node and both subtrees must match
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
