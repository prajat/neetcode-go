package main

/*
Given an integer array nums, find a subarray that has the largest product, and return the product.
A subarray is a contiguous non-empty sequence of elements within an array.
You can assume the output will fit into a 32-bit integer.
Note that the product of an array with a single element is the value of that element.
*/

// kadane's algorithm
func maxProduct(nums []int) int {
	currentMax, currentMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// if num is negative number, swap the currentMax and currentMin
		if num < 0 {
			currentMax, currentMin = currentMin, currentMax
		}

		currentMax = max(num, currentMax*num)
		currentMin = min(num, currentMin*num)

		result = max(result, currentMax)
	}
	return result
}
