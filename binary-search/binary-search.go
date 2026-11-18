package main

/*You are given an array of distinct integers nums, sorted in ascending order, and an integer target.
Implement a function to search for target within nums. If it exists, then return its index, otherwise, return -1.
*/

// Binary Search - O(logn)
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
