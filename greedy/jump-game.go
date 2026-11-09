package main

/*
You are given an integer array nums where each element nums[i] indicates your maximum jump length at that position.
Return true if you can reach the last index starting from index 0, or false otherwise.
*/

// Simple recursion - exponentioal time complexity
func canJump(nums []int) bool {
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index >= len(nums)-1 {
			return true
		}

		for jump := 1; jump <= nums[index]; jump++ {
			if dfs(index + jump) {
				return true
			}
		}
		return false
	}
	return dfs(0)
}

// recursion + memoisation
func canJumpMemo(nums []int) bool {
	// 0=unknown, 1=true, 2=false
	cache := make([]int8, len(nums))

	var dfs func(int) bool
	dfs = func(index int) bool {
		if index >= len(nums)-1 {
			return true
		}

		if cache[index] != 0 {
			return cache[index] == 1
		}

		for jump := 1; jump <= nums[index]; jump++ {
			if dfs(index + jump) {
				cache[index] = 1
				return true
			}
		}

		cache[index] = 2
		return false
	}
	return dfs(0)
}

// greddy - optimal solution
func canJumpGreedy(nums []int) bool {
	farthest := 0

	for i := range nums {
		// current index is unreacheable
		if i > farthest {
			return false
		}

		farthest = max(farthest, i+nums[i])

		if farthest >= len(nums)-1 {
			return true
		}
	}
	return true
}
