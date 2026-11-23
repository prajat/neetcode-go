package main

/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard so that no two queens can attack each other.
A queen in a chessboard can attack horizontally, vertically, and diagonally.
Given an integer n, return all distinct solutions to the n-queens puzzle.
Each solution contains a unique board layout where the queen pieces are placed. 'Q' indicates a queen and '.' indicates an empty space.
You may return the answer in any order.
*/
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]byte, n)

	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // row-col + (n-1)
	diag2 := make([]bool, 2*n+1) // row+col

	var dfs func(int)
	dfs = func(col int) {
		if col == n {
			// copy the board to the result
			solution := make([]string, n)
			for i := 0; i < n; i++ {
				solution[i] = string(board[i])
			}
			result = append(result, solution)
			return
		}

		for row := 0; row < n; row++ {
			d1 := row - col + (n - 1)
			d2 := row + col

			if cols[row] || diag1[d1] || diag2[d2] {
				continue
			}

			board[row][col] = 'Q'
			cols[row] = true
			diag1[d1] = true
			diag2[d2] = true

			dfs(col + 1)

			board[row][col] = '.'
			cols[row] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}
	dfs(0)
	return result
}
