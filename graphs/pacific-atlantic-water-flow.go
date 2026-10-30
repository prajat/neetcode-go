package main

/*
You are given a rectangular island heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
The islands borders the Pacific Ocean from the top and left sides, and borders the Atlantic Ocean from the bottom and right sides.
Water can flow in four directions (up, down, left, or right) from a cell to a neighboring cell with height equal or lower. Water can also flow into the ocean from cells adjacent to the ocean.
Find all cells where water can flow from that cell to both the Pacific and Atlantic oceans. Return it as a 2D list where each element is a list [r, c] representing the row and column of the cell. You may return the answer in any order.
*/
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	rows := len(heights)
	cols := len(heights[0])

	pacific := make(map[[2]int]bool)
	atlantic := make(map[[2]int]bool)

	directions := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	var dfs func(int, int, map[[2]int]bool)
	dfs = func(row, col int, visited map[[2]int]bool) {
		position := [2]int{row, col}

		if visited[position] {
			return
		}

		visited[position] = true

		for _, direction := range directions {
			nextRow := row + direction[0]
			nextCol := col + direction[1]

			if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
				continue
			}

			nextPosition := [2]int{nextRow, nextCol}

			//Reverse water flow condition
			if heights[nextRow][nextCol] >= heights[row][col] {
				if !visited[nextPosition] {
					dfs(nextRow, nextCol, visited)
				}
			}
		}
	}

	// pacific ocean: top row and left column
	for col := range cols {
		dfs(0, col, pacific)
	}

	for row := range rows {
		dfs(row, 0, pacific)
	}

	// atlantic ocean: last row and right column
	for col := range cols {
		dfs(rows-1, col, atlantic)
	}

	for row := range rows {
		dfs(row, cols-1, atlantic)
	}

	// cells reachable from both oceans are the answer.
	result := make([][]int, 0)

	for row := range rows {
		for col := range cols {
			position := [2]int{row, col}
			if pacific[position] && atlantic[position] {
				result = append(result, position[:])
			}
		}
	}
	return result
}
