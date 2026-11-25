package main

/*
You are given a 2-D matrix grid. Each cell can have one of three possible values:
0 representing an empty cell
1 representing a fresh fruit
2 representing a rotten fruit
Every minute, if a fresh fruit is horizontally or vertically adjacent to a rotten fruit, then the fresh fruit also becomes rotten.
Return the minimum number of minutes that must elapse until there are zero fresh fruits remaining. If this state is impossible within the grid, return -1.
*/
func orangesRotting(grid [][]int) int {
	queue := make([]cell, 0)

	fresh := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 2 {
				queue = append(queue, cell{row: r, col: c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	minutes := 0
	// bfs
	for len(queue) > 0 {
		size := len(queue)
		rotted := false
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				nr := curr.row + d[0]
				nc := curr.col + d[1]

				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
					continue
				}
				if grid[nr][nc] != 1 {
					continue
				}

				grid[nr][nc] = 2
				fresh--
				rotted = true
				queue = append(queue, cell{row: nr, col: nc})
			}
		}

		if rotted {
			minutes++
		}
	}

	// check back if any fresh fruit is left
	if fresh > 0 {
		return -1
	}
	return minutes
}
