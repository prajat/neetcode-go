package main

/*
You are given an integer array nums where nums[i] represents the amount of money the ith house has. The houses are arranged in a circle, i.e. the first house and the last house are neighbors.
You are planning to rob money from the houses, but you cannot rob two adjacent houses because the security system will automatically alert the police if two adjacent houses were both broken into.
Return the maximum amount of money you can rob without alerting the police.
*/

// Simple recursive solution for max money without the condition that houses are in circle
func rob2(nums []int) int {
	var dfs func([]int, int) int
	dfs = func(nums []int, index int) int {
		if index >= len(nums) {
			return 0
		}
		robCurrent := nums[index] + dfs(nums, index+2)
		skipCurrent := dfs(nums, index+1)
		return max(robCurrent, skipCurrent)
	}
	return dfs(nums, 0)
}

// recursion with memoisation - top down approach
func rob3(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	/*Because the houses form a circle, the first and last houses are adjacent. Therefore, we have two independent choices:
	Rob houses from index 0 to n-2 — exclude the last house.
	Rob houses from index 1 to n-1 — exclude the first house.
	Take the maximum of both cases. */
	return max(rob3Memo(nums, 0, n-2), rob3Memo(nums, 1, n-1))
}

func rob3Memo(nums []int, start, end int) int {
	memo := make([]int, len(nums))
	visited := make([]bool, len(nums))

	var dfs func(int) int
	dfs = func(index int) int {
		if index > end {
			return 0
		}

		if visited[index] {
			return memo[index]
		}

		skipCurrent := dfs(index + 1)
		robCurrent := nums[index] + dfs(index+2)

		memo[index] = max(skipCurrent, robCurrent)
		visited[index] = true

		return memo[index]
	}
	return dfs(start)
}

// dynamic programming - bottom up approach
func rob4(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(rob4DP(nums, 0, n-2), rob4DP(nums, 1, n-1))
}

// dp[i] = maximum money that can be robbed up to house i
func rob4DP(nums []int, start, end int) int {
	length := end - start + 1
	dp := make([]int, length+1)

	// no money in 0th house
	dp[0] = 0

	if length >= 1 {
		dp[1] = nums[start]
	}

	for i := 2; i <= length; i++ {
		currentHouse := nums[start+i-1]

		skipCurrent := dp[i-1]
		robCurrent := currentHouse + dp[i-2]

		dp[i] = max(skipCurrent, robCurrent)
	}
	return dp[length]
}
