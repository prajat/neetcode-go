package main

import (
	"slices"
)

/*
We can iterate through nums with index i and get nums[i] = -(nums[j] + nums[k]) after rearranging the equation, making -nums[i] = nums[j] + nums[k]. For each index i, we should efficiently calculate the j and k pairs without duplicates.
To efficiently find the j and k pairs, we run the two pointer approach on the elements to the right of index i as the array is sorted. When we run two pointer algorithm, consider j and k as pointers (j is at left, k is at right) and target = -nums[i], if the current sum num[j] + nums[k] < target then we need to increase the value of current sum by incrementing j pointer. Else if the current sum num[j] + nums[k] > target then we should decrease the value of current sum by decrementing k pointer.
When the current sum nums[j] + nums[k] == target add this pair to the result. We can move j or k pointer until j < k and the pairs are repeated. This ensures that no duplicate pairs are added to the result.
*/
func ThreeSum(nums []int) [][]int {
	output := make([][]int, 0)
	slices.Sort(nums) //pattern observe -> -nums[i] = nums[j] + nums[k]
	for i := 0; i < len(nums)-2; i++ {
		// skip duplicate first values
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		target := -nums[i]
		for j < k {
			if nums[j]+nums[k] == target {
				output = append(output, []int{nums[i], nums[j], nums[k]})

				// Skip duplicate second values
				for j < k && nums[j] == nums[j+1] {
					j++
				}

				// Skip duplicate third value
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				//move both pointer after finding a triplet
				j++
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return output
}
