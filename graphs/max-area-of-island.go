package main

/*
You are given a matrix grid where grid[i] is either a 0 (representing water) or 1 (representing land).
An island is defined as a group of 1's connected horizontally or vertically. You may assume all four edges of the grid are surrounded by water.
The area of an island is defined as the number of cells within the island.
Return the maximum area of an island in grid. If no island exists, return 0.
*/
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if row >= len(grid) || row < 0 || col >= len(grid[row]) || col < 0 || grid[row][col] == 0 {
			return 0
		}

		// mark cell visited
		grid[row][col] = 0

		return 1 + dfs(row-1, col) + dfs(row+1, col) + dfs(row, col+1) + dfs(row, col-1)
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				area := dfs(row, col)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
