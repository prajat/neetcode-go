package main

/*
Given a 2-D grid of characters board and a list of strings words, return all words that are present in the grid.
For a word to be present it must be possible to form the word with a path in the board with horizontally or vertically neighboring cells. The same cell may not be used more than once in a word.
*/

// Search for each word independently using DFS and backtracking. - Ineffecient
func findWords(board [][]byte, words []string) []string {
	result := make([]string, 0)

	if len(board) == 0 || len(board[0]) == 0 {
		return result
	}

	for i := range words {
		if exists(board, words[i]) {
			result = append(result, words[i])
		}
	}
	return result
}

func exists(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
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

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) || board[row][col] != word[index] {
		return false
	}

	// Mark the current cell as visited
	original := board[row][col]
	board[row][col] = '#'

	found := search(board, word, row-1, col, index+1) || search(board, word, row+1, col, index+1) || search(board, word, row, col-1, index+1) || search(board, word, row, col+1, index+1)

	// Restore the cell for other paths
	board[row][col] = original
	return found
}

// use a Trie plus backtracking. The Trie avoids searching separately from every word.

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWordsV2(board [][]byte, words []string) []string {
	root := &TrieNode{}

	// Build the trie
	for _, word := range words {
		current := root

		for i := range word {
			index := word[i] - 'a'

			if current.children[index] == nil {
				current.children[index] = &TrieNode{}
			}

			current = current.children[index]
		}

		// store the complete word at its terminal node
		current.word = word
	}

	result := make([]string, 0)

	var dfs func(int, int, *TrieNode)
	dfs = func(row, col int, node *TrieNode) {
		// Check boundaries and visited cells
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) || board[row][col] == '#' {
			return
		}

		index := board[row][col] - 'a'
		next := node.children[index]

		// no word has this prefix
		if next == nil {
			return
		}

		// a complete word has been found
		if next.word != "" {
			result = append(result, next.word)

			// prevent duplicate results
			next.word = ""
		}

		// mark the current cell as visited
		original := board[row][col]
		board[row][col] = '#'

		dfs(row-1, col, next)
		dfs(row+1, col, next)
		dfs(row, col-1, next)
		dfs(row, col+1, next)

		// Restore the cell paths
		board[row][col] = original
	}

	for row := range board {
		for col := range board[row] {
			dfs(row, col, root)
		}
	}

	return result
}
