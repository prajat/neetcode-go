package main

/*
Given an m x n matrix of integers matrix, if an element is 0, set its entire row and column to 0's.
You must update the matrix in-place.
Follow up: Could you solve it using O(1) space?
*/
// Simple Version : O(m+n) space
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zeroRows, zeroCols := make([]bool, m), make([]bool, n)

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 0 {
				zeroRows[row], zeroCols[col] = true, true
			}
		}
	}

	for row := range matrix {
		for col := range matrix[row] {
			if zeroRows[row] || zeroCols[col] {
				matrix[row][col] = 0
			}
		}
	}
}
