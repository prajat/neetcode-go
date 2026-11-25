package main

/*
You are given an array of integers cost where cost[i] is the cost of taking a step from the ith floor of a staircase. After paying the cost, you can step to either the (i + 1)th floor or the (i + 2)th floor.
You may choose to start at the index 0 or the index 1 floor.
Return the minimum cost to reach the top of the staircase, i.e. just past the last index in cost.
*/
func minCostClimbingStairs(cost []int) int {
	cache := make([]int, len(cost))

	var dfs func(int) int
	dfs = func(index int) int {
		if index >= len(cost) {
			return 0
		}

		if cache[index] != 0 {
			return cache[index]
		}

		cost := cost[index] + min(dfs(index+1), dfs(index+2))
		cache[index] = cost

		return cache[index]
	}

	return min(dfs(0), dfs(1))
}
