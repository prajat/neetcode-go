package main

import "slices"

/*
Given a square n x n matrix of integers matrix, rotate it by 90 degrees clockwise.
You must rotate the matrix in-place. Do not allocate another 2D matrix and do the rotation.
*/
/*Transpose, then reverse each row
A 90° clockwise rotation is equivalent to:
Transpose the matrix, swapping matrix[i][j] with matrix[j][i].
Reverse each row. */
func rotate(matrix [][]int) {
	n := len(matrix)

	// transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each row
	for i := 0; i < n; i++ {
		for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
}

func rotateV2(matrix [][]int) {
	n := len(matrix)

	// transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each row using slices.reverse
	for i := 0; i < n; i++ {
		slices.Reverse(matrix[i])
	}
}
