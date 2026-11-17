package main

/*
You are given an array of non-negative integers height which represent an elevation map. Each value height[i] represents the height of a bar, which has a width of 1.
Return the total amount of water that can be trapped between the bars.
*/

// The water above a bar is min(maxLeft, maxRight) - height[i]. Track the left and right maximums with two pointers, and always move the side with the smaller maximum. That side's water level is already decided by its own maximum.
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	water := 0

	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			water += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			water += rightMax - height[right]
		}
	}
	return water
}
