package main

/*
There is an m x n grid where you are allowed to move either down or to the right at any point in time.
Given the two integers m and n, return the number of possible unique paths that can be taken from the top-left corner of the grid (grid[0][0]) to the bottom-right corner (grid[m - 1][n - 1]).
You may assume the output will fit in a 32-bit integer.
*/

// using recursion(dfs) only
func uniquePaths(m, n int) int {
	var dfs func(int, int) int
	dfs = func(row, col int) int {
		// out of the grid
		if row >= m || col >= n {
			return 0
		}

		// reached the destination
		if row == m-1 && col == n-1 {
			return 1
		}

		return dfs(row+1, col) + dfs(row, col+1)
	}
	return dfs(0, 0)
}

// recursion + memoisation
func uniquePathsV2(m, n int) int {
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if row >= m || col >= n {
			return 0
		}

		if row == m-1 && col == n-1 {
			return 1
		}

		if cache[row][col] != 0 {
			return cache[row][col]
		}

		cache[row][col] = dfs(row+1, col) + dfs(row, col+1)
		return cache[row][col]
	}
	return dfs(0, 0)
}

// bottom-up DP
func uniquePathsDP(m, n int) int {
	dp := make([]int, n)

	for i := range dp {
		// first row: only one way to reach each cell
		dp[i] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[col] += dp[col-1]
		}
	}
	return dp[n-1]
}
