package main

/*
You are given an m x n 2-D integer array matrix and an integer target.
Each row in matrix is sorted in non-decreasing order.
The first integer of every row is greater than the last integer of the previous row.
Return true if target exists within matrix or false otherwise.
Can you write a solution that runs in O(log(m * n)) time?
*/
func searchMatrix(matrix [][]int, target int) bool {
	for row := range matrix {
		if binarySearch(matrix[row], target) != -1 {
			return true
		}
	}
	return false
}

/*
Binary search on the flattened matrix
Every row is sorted, and each row starts after the previous one ends. So the matrix behaves like one sorted array of m * n elements. Binary search it, and convert a flat index to a row and column.
Go
*/
func searchMatrixV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2

		value := matrix[mid/cols][mid%cols]

		switch {
		case value == target:
			return true
		case value < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return false
}
