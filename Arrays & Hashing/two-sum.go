package main

func twoSum(nums []int, target int) []int {
	output := make([]int, 0)
	for i := range len(nums) {
		for j := range len(nums) {
			if nums[i]+nums[j] == target && i != j {
				output = append(output, i)
				output = append(output, j)
				break
			}
		}
	}

	return output
}
