package main

// Brute Force method - O(n2)
func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	for i := range nums {
		product := 1
		for j := range nums {
			if j == i {
				continue
			}
			product = product * nums[j]
		}
		output[i] = product
	}
	return output
}

// Using Prefix And Suffix Approach O(n)
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	// Store the product of all elements to the left.
	prefix := 1
	for i := range nums {
		output[i] = prefix
		prefix *= nums[i]
	}
	// Multiply by the product of all elements to the right.
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= suffix
		suffix *= nums[i]
	}
	return output
}
