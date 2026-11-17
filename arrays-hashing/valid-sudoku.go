package main

/*
You are given a 9 x 9 Sudoku board board. A Sudoku board is valid if the following rules are followed:
Each row must contain the digits 1-9 without duplicates.
Each column must contain the digits 1-9 without duplicates.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
Return true if the Sudoku board is valid, otherwise return false
Note: A board does not need to be full or be solvable to be valid.
*/
func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			digit := int(board[r][c] - '1')
			box := (r/3)*3 + c/3

			if rows[r][digit] || cols[c][digit] || boxes[box][digit] {
				return false
			}

			rows[r][digit] = true
			cols[c][digit] = true
			boxes[box][digit] = true
		}
	}
	return true
}
