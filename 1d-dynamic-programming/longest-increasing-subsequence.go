package main

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.
A subsequence is a sequence that can be derived from the given sequence by deleting some or no elements without changing the relative order of the remaining characters.
For example, "cat" is a subsequence of "crabt".
*/

// using recursion only
func lengthOfLIS(nums []int) int {
	var dfs func(int, int) int
	dfs = func(index, previousIndex int) int {
		if index == len(nums) {
			return 0
		}

		// option-1: skip the current number.
		skip := dfs(index+1, previousIndex)

		// option-2: include the current number if it is increasing
		include := 0
		if previousIndex == -1 || nums[index] > nums[previousIndex] {
			include = 1 + dfs(index+1, index)
		}

		return max(skip, include)
	}
	return dfs(0, -1)
}

// recursion + memoisation
func lengthOfLISV2(nums []int) int {
	n := len(nums)

	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, n+1)

		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(index, previousIndex int) int {
		if index == n {
			return 0
		}

		cacheIndex := previousIndex + 1
		if cache[index][cacheIndex] != -1 {
			return cache[index][cacheIndex]
		}

		skip := dfs(index+1, previousIndex)

		include := 0
		if previousIndex == -1 || nums[index] > nums[previousIndex] {
			include = 1 + dfs(index+1, index)
		}

		cache[index][cacheIndex] = max(skip, include)
		return cache[index][cacheIndex]
	}
	return dfs(0, -1)
}
