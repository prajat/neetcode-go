package main

/*
You are given an array of length n which was originally sorted in ascending order. It has now been rotated between 1 and n times. For example, the array nums = [1,2,3,4,5,6] might become:
[3,4,5,6,1,2] if it was rotated 4 times.
[1,2,3,4,5,6] if it was rotated 6 times.
Notice that rotating the array 4 times moves the last four elements of the array to the beginning. Rotating the array 6 times produces the original array.
Assuming all elements in the rotated sorted array nums are unique, return the minimum element of this array.
*/

// Brute Force Method - O(n)
// A brute force solution would be to do a linear search on the array to find the minimum element. This would be an O(n) solution.
func findMin(nums []int) int {
	minimum := nums[0]
	for i := range nums {
		minimum = min(minimum, nums[i])
	}
	return minimum
}

// Optimal Solution - O(logn)
/*Given that the array is rotated after sorting, elements from the right end are moved to the left end one by one. This creates two parts of a sorted array, separated by a deflection point caused by the rotation. For example, consider the array [3, 4, 1, 2]. Here, the array is rotated twice, resulting in two sorted segments: [3, 4] and [1, 2]. And the minimum element will be the first element of the right segment.  */
func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// The min is in the right half
			left = mid + 1
		} else {
			// The min is in the left half
			right = mid
		}
	}
	return nums[left]
}
