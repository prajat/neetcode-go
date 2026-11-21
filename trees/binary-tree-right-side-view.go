package main

/*You are given the root of a binary tree. Return only the values of the nodes that are visible from the right side of the tree, ordered from top to bottom. */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := []*TreeNode{root}

	var bfs func(*TreeNode)
	bfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		for len(queue) > 0 {
			levelSize := len(queue)
			level := make([]int, 0, levelSize)

			for i := 0; i < levelSize; i++ {
				currNode := queue[0]
				queue = queue[1:]

				level = append(level, currNode.Val)

				if currNode.Left != nil {
					queue = append(queue, currNode.Left)
				}

				if currNode.Right != nil {
					queue = append(queue, currNode.Right)
				}
			}
			result = append(result, level[levelSize-1])
		}
	}
	bfs(root)
	return result
}
