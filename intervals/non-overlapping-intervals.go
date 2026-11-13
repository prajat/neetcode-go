package main

import "sort"

/*
Given an array of intervals intervals where intervals[i] = [start_i, end_i], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
Note: Intervals are non-overlapping even if they have a common point. For example, [1, 3] and [2, 4] are overlapping, but [1, 2] and [2, 3] are non-overlapping.
*/
func eraseOverlapIntervals(intervals [][]int) int {
	erase := 0

	// sort the intervals based on thier start values
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {

		// if the current interval doesn't overlap
		if intervals[i][0] >= prevEnd {
			// update prevEnd to current interval's end
			prevEnd = intervals[i][1]
			continue
		}

		// set prevEnd to the min of prevEnd and the current interval's end
		prevEnd = min(prevEnd, intervals[i][1])

		// greedily removing the interval that ends last to retain as many intervals as possible.
		erase++
	}

	return erase
}
