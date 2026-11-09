package main

import "sort"

/*
Given an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
You may return the answer in any order.
Note: Intervals are non-overlapping if they have no common point. For example, [1, 2] and [3, 4] are non-overlapping, but [1, 2] and [2, 3] are overlapping.
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// sort the intervals in ascending order based on start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	output := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {

		last := output[len(output)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			// overlap: extend the end of the last interval
			last[1] = max(last[1], current[1])
		} else {
			// no overlap: start a new interval
			output = append(output, current)
		}
	}

	return output
}
