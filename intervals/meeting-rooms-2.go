package main

import (
	"slices"
)

/*
Given an array of meeting time interval objects consisting of start and end times [[start_1,end_1],[start_2,end_2],...] (start_i < end_i), find the minimum number of rooms required to schedule all meetings without any conflicts.
Note: (0,8),(8,10) is NOT considered a conflict at 8.
*/
func minMeetingRooms(intervals []Interval) int {
	//We create two arrays, start and end, containing the start and end times of all meetings, respectively. After sorting both arrays, we use a two-pointer based approach
	start, end := make([]int, len(intervals)), make([]int, len(intervals))

	for i, interval := range intervals {
		start[i] = interval.start
		end[i] = interval.end
	}

	// sort both the slices
	slices.Sort(start)
	slices.Sort(end)

	s, e, count, result := 0, 0, 0, 0

	/*We use two pointers, s and e, for the start and end arrays, respectively. We also maintain a variable count to track the current number of active meetings. At each iteration, we increment s while the start time is less than the current end time and increase count, as these meetings must begin before the earliest ongoing meeting ends. */
	/*Then, we increment e and decrement count as a meeting has ended. At each step, we update the result with the maximum value of active meetings stored in count. */
	for s < len(start) {
		if start[s] < end[e] {
			s++
			count++
		} else {
			e++
			count--
		}
		result = max(result, count)
	}
	return result
}
