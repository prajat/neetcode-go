package main

import "slices"

/*
You are given an array nums of integers, which may contain duplicates. Return all possible subsets.
The solution must not contain duplicate subsets. You may return the solution in any order.
*/
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(int)
	dfs = func(start int) {
		subset := make([]int, len(path))
		copy(subset, path)
		result = append(result, subset)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}
