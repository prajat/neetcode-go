package main

/*
Given an array nums containing n integers in the range [0, n] without any duplicates, return the single number in the range that is missing from nums.
Follow-up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?
*/

// Using Hash Set O(n) extra space
func missingNumber(nums []int) int {
	n := len(nums)
	hashSet := make(map[int]bool, n)
	for i := range nums {
		hashSet[nums[i]] = true
	}

	for i := 0; i <= n; i++ {
		if hashSet[i] == false {
			return i
		}
	}
	return 0
}

// Using Bitwise XOR - O(1) extra space
func missingNumberV2(nums []int) int {
	n := len(nums)
	result := 0
	for i := 0; i <= n; i++ {
		result = result ^ i
	}

	for i := range nums {
		result = result ^ nums[i]
	}
	return result
}
