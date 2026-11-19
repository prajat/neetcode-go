package main

import (
	"slices"
)

/*
You are given two integer arrays nums1 and nums2 of size m and n respectively, where each is sorted in ascending order. Return the median value among all elements of the two arrays.
Your solution should run in O(log(m+n)) time
*/

// Not optimal - O((m+n)log(m+n))
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	nums3 := make([]int, 0, len(nums1)+len(nums2))
	nums3 = append(nums3, nums1...)
	nums3 = append(nums3, nums2...)

	slices.Sort(nums3)

	n := len(nums3)

	if n%2 == 0 {
		return float64(nums3[n/2-1]+nums3[n/2]) / 2.0
	} else {
		return float64(nums3[n/2])
	}
}
