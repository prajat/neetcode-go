package main

/*
We keep two pointers, one at the start and the other at the end of the array.
If the sum of the numbers at the two pointers is greater than the target, decrement the right pointer, else increment the left pointer.
Repeat this process until you find a valid pair.
*/

func TwoSum(numbers []int, target int) []int {
	index1 := 0
	index2 := len(numbers) - 1
	for index1 < index2 {
		if numbers[index1]+numbers[index2] == target {
			break
		}
		if numbers[index1]+numbers[index2] > target {
			index2--
		} else {
			index1++
		}
	}
	return []int{index1 + 1, index2 + 1}
}
