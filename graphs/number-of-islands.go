package main

/*
Given a 2D grid grid where '1' represents land and '0' represents water, count and return the number of islands.
An island is formed by connecting adjacent lands horizontally or vertically and is surrounded by water. You may assume water is surrounding the grid (i.e., all the edges are water).
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islands := 0

	var dfs func(int, int)
	dfs = func(row, col int) {
		// Stop at boundaries
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == '0' {
			return
		}

		// mark this land as visited
		grid[row][col] = '0'

		// Visit all four direction
		dfs(row-1, col) // up
		dfs(row+1, col) // down
		dfs(row, col-1) // left
		dfs(row, col+1) // right
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				islands++
				dfs(row, col)
			}
		}
	}
	return islands
}
