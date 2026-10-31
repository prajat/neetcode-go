package main

/*
You are given an integer n representing the number of steps to reach the top of a staircase. You can climb with either 1 or 2 steps at a time.
Return the number of distinct ways to climb to the top of the staircase.
*/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsV2(n int) int {
	cache := make([]int, n+1) // for memoization (1-d dynamic programming) of results

	var dfs func(int) int
	dfs = func(remaining int) int {
		if remaining == 0 {
			return 1
		}

		if remaining < 0 {
			return 0
		}

		if cache[remaining] != 0 {
			return cache[remaining]
		}

		// take either one step or two steps
		cache[remaining] = dfs(remaining-1) + dfs(remaining-2)
		return cache[remaining]
	}
	return dfs(n)
}
