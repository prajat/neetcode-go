package main

/*Given a binary tree root, return the level order traversal of it as a nested list, where each sublist contains the values of nodes at a particular level in the tree, from left to right. */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// initialise result array and the queue(slice) with root added as first element
	result := make([][]int, 0)
	queue := []*TreeNode{root}

	// Loop till queue is non empty
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		// pop out the first element from the queue  in a levelSize loop
		for i := 0; i < levelSize; i++ {
			currNode := queue[0]
			queue = queue[1:]
			// append the current node val in the level array
			level = append(level, currNode.Val)

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
