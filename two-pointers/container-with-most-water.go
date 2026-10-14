package main

// Brute Force Method - O(n2)
func maxArea(heights []int) int {
	result := 0
	for i := range heights {
		for j := range heights {
			width := j - i
			height := min(heights[i], heights[j])
			area := width * height
			if area > result {
				result = area
			}
		}
	}
	return result
}

// two-pointer approach - O(n)
func MaxArea(heights []int) int {
	result := 0
	start := 0
	end := len(heights) - 1
	for start < end {
		width := end - start
		height := min(heights[start], heights[end])
		area := height * width
		result = max(area, result)
		if heights[start] < heights[end] {
			start++
		} else {
			end--
		}
	}
	return result
}
