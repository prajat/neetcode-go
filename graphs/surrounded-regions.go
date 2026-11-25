package main

/*
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:
Connect: A cell is connected to adjacent cells horizontally or vertically.
Region: To form a region connect every 'O' cell. Regions can have any shape; they do not need to be squares or rectangles.
Surround: A region is surrounded if none of the 'O' cells in that region are on the edge of the board. Such regions are completely enclosed by 'X' cells.
To capture a surrounded region, replace all 'O's with 'X's in-place within the original board. You do not need to return anything.
*/
func solve(board [][]byte) {
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] == 'X' || board[row][col] == '#' {
			return
		}

		board[row][col] = '#'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for col := 0; col < len(board[0]); col++ {
		if board[0][col] == 'O' {
			dfs(0, col)
		}
	}

	for col := 0; col < len(board[0]); col++ {
		row := len(board) - 1
		if board[row][col] == 'O' {
			dfs(row, col)
		}
	}

	for row := 0; row < len(board); row++ {
		if board[row][0] == 'O' {
			dfs(row, 0)
		}
	}

	for row := 0; row < len(board); row++ {
		col := len(board[0]) - 1
		if board[row][col] == 'O' {
			dfs(row, col)
		}
	}

	for r := range board {
		for c := range board[r] {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}
}
