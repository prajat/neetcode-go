package main

import "slices"

// Brute Force Method - O(n2)
func longestConsecutive(nums []int) int {
	longest := 0
	for i := range nums {
		current := nums[i]
		length := 1
		for slices.Contains(nums, current+1) {
			length++
			current++
		}

		if length > longest {
			longest = length
		}
	}
	return longest
}

/*
func contains(nums []int, target int) bool {
	for i := range nums {
		if nums[i] == target {
			return true
		}
	}
	return false
}
*/
