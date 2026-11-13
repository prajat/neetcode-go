package main

import "sort"

/*
Given an array of meeting time interval objects consisting of start and end times [[start_1,end_1],[start_2,end_2],...] (start_i < end_i), determine if a person could add all meetings to their schedule without any conflicts. The intervals may be provided in any order.
Note: (0,8),(8,10) is not considered a conflict at 8
*/

type Interval struct {
	start int
	end   int
}

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) < 2 {
		return true
	}

	// sort the intervals in ascending order based on thier start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	prevEnd := intervals[0].end
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < prevEnd {
			// overlapped intervals
			return false
		}

		prevEnd = intervals[i].end
	}

	return true
}
