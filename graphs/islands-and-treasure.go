package main

/*
You are given a m×n 2D grid initialized with these three possible values:
-1 - A water cell that can not be traversed.
0 - A treasure chest.
INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.
Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest then the value should remain INF.
Assume the grid can only be traversed up, down, left, or right.
Modify the grid in-place.
*/

/*We insert all the cells (row, col) that represent the treasure chests into the queue. Then, we process the cells level by level, handling all the current cells in the queue at once. For each cell, we mark it as visited and store the current level value as the distance at that cell. We then try to add the neighboring cells (adjacent cells) to the queue, but only if they have not been visited and are land cells. */
type cell struct {
	row, col int
}

func islandsAndTreasure(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	queue := make([]cell, 0)
	visited := make([][]bool, len(grid))
	for r := range visited {
		visited[r] = make([]bool, len(grid[0]))
	}

	// insert all cells that represent treasure chest
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				queue = append(queue, cell{row: row, col: col})
				visited[row][col] = true
			}
		}
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	distance := 0

	// Process the queue
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			grid[curr.row][curr.col] = distance

			for _, d := range directions {
				nr := curr.row + d[0]
				nc := curr.col + d[1]

				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
					continue
				}
				if visited[nr][nc] || grid[nr][nc] == -1 {
					continue
				}

				visited[nr][nc] = true
				queue = append(queue, cell{row: nr, col: nc})
			}
		}
		distance++
	}
}
