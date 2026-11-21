package main

/*
Given an array nums of unique integers, return all possible subsets of nums.
The solution set must not contain duplicate subsets. You may return the solution in any order
*/
func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(int)
	dfs = func(start int) {
		subset := make([]int, len(path))
		copy(subset, path)
		res = append(res, subset)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
