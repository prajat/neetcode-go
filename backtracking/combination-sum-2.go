package main

import "slices"

/*
You are given an array of integers candidates, which may contain duplicates, and a target integer target. Your task is to return a list of all unique combinations of candidates where the chosen numbers sum to target.
Each element from candidates may be chosen at most once within a combination. The solution set must not contain duplicate combinations.
You may return the combinations in any order and the order of the numbers in each combination can be in any order.
*/
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	slices.Sort(candidates)

	path := make([]int, 0)

	var dfs func(int, int)
	dfs = func(start, sum int) {
		if sum == target {
			combo := make([]int, len(path))
			copy(combo, path)
			result = append(result, combo)
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			nextSum := sum + candidates[i]
			if nextSum > target {
				break
			}

			path = append(path, candidates[i])
			dfs(i+1, nextSum)
			path = path[:len(path)-1]
		}

	}
	dfs(0, 0)
	return result
}
