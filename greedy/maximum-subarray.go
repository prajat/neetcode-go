package main

/*
Given an array of integers nums, find the subarray with the largest sum and return the sum.
A subarray is a contiguous non-empty sequence of elements within an array.
*/

// kadane's algo
func maxSubArray(nums []int) int {
	maxSum, currSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		currSum = max(num, currSum+num)

		maxSum = max(maxSum, currSum)
	}
	return maxSum
}
