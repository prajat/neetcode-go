package main

/*Given an m x n matrix of integers matrix, return a list of all elements within the matrix in spiral order. */
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	result := make([]int, 0, len(matrix)*len(matrix[0]))

	for top <= bottom && left <= right {
		//left to right along thr top row
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// top to bottom along the right column
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// right to left along the bottom row, if a row remains
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		// Bottom to top along the left column, if a column remains
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}
	return result
}
