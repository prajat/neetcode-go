package main

/*
You are given an integer array nums where nums[i] represents the amount of money the ith house has. The houses are arranged in a straight line, i.e. the ith house is the neighbor of the (i-1)th and (i+1)th house.
You are planning to rob money from the houses, but you cannot rob two adjacent houses because the security system will automatically alert the police if two adjacent houses were both broken into.
Return the maximum amount of money you can rob without alerting the police.
*/

// Recursive solution
func rob(nums []int) int {
	var robFrom func([]int, int) int
	robFrom = func(nums []int, index int) int {
		if index >= len(nums) {
			return 0
		}

		robCurrent := nums[index] + robFrom(nums, index+2)
		skipCurrent := robFrom(nums, index+1)

		return max(robCurrent, skipCurrent)
	}
	return robFrom(nums, 0)
}

// 1-d Dynamic programming
// dp[i] = maximum money we can rob from the first i houses
// skip current house: dp[i-1]
// rob current house: nums[i-1] + dp[i-2]
func robDP(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n+1)

	// no house means no money.
	dp[0] = 0

	// one house means rob that house
	dp[1] = nums[0]

	for i := 2; i <= n; i++ {
		robCurrent := nums[i-1] + dp[i-2]
		skipCurrent := dp[i-1]

		dp[i] = max(robCurrent, skipCurrent)
	}
	return dp[n]
}
