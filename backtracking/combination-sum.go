package main

import "slices"

/*
You are given an array of distinct integers nums and a target integer target. Your task is to return a list of all unique combinations of nums where the chosen numbers sum to target.
The same number may be chosen from nums an unlimited number of times. Two combinations are the same if the frequency of each of the chosen numbers is the same, otherwise they are different.
You may return the combinations in any order and the order of the numbers in each combination can be in any order.
*/
func combinationSum(nums []int, target int) [][]int {
	// Sort the incoming array
	slices.Sort(nums)

	result := make([][]int, 0)
	current := make([]int, 0)

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			// store current as copy to result as it will get changed in backtracking calls
			result = append(result, slices.Clone(current))
			return
		}

		for i := start; i < len(nums); i++ {
			// since nums is sorted, later values will also be large
			if nums[i] > remaining {
				break
			}

			current = append(current, nums[i])

			// pass i, not i+1 as same number can be reused
			backtrack(i, remaining-nums[i])

			// undo the choice
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target)
	return result
}
