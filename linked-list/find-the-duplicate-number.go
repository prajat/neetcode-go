package main

/*
You are given an array of integers nums containing n + 1 integers. Each integer in nums is in the range [1, n] inclusive.
There is exactly one repeated integer in nums, and every other integer appears at most once.
Return the repeated integer.
*/

/*
Floyd's cycle detection, O(1) space
Treat each value as a pointer: index i points to nums[i]. Since there are n + 1 slots and values in 1..n, the duplicate value is where two pointers enter the same index, which forms a cycle.
*/
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// phase 1: find a meeting point inside the cycle
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	// Phase 2: find the cycle entrance, which is the duplicate
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// using hashset and O(n) extra space
func findDuplicateV2(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}
