package main

/*
Given a 2-D grid of characters board and a string word, return true if the word is present in the grid, otherwise return false.
For the word to be present it must be possible to form it with a path in the board with horizontally or vertically neighboring cells. The same cell may not be used more than once in a word.
*/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for row := range board {
		for col := range board[row] {
			if board[row][col] == word[0] && search(board, word, row, col, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, row, col, index int) bool {
	if index == len(word) {
		return true
	}

	// check boundaries
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) || board[row][col] != word[index] {
		return false
	}

	// mark visited
	original := board[row][col]
	board[row][col] = '#'

	found := search(board, word, row-1, col, index+1) || search(board, word, row+1, col, index+1) || search(board, word, row, col-1, index+1) || search(board, word, row, col+1, index+1)

	// restore visted cell
	board[row][col] = original

	return found
}
