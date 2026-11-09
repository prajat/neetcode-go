package main

/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [start_i, end_i] represents the start and the end time of the ith interval. intervals is initially sorted in ascending order by start_i.
You are given another interval newInterval = [start, end].
Insert newInterval into intervals such that intervals is still sorted in ascending order by start_i and also intervals still does not have any overlapping intervals. You may merge the overlapping intervals if needed.
Return intervals after adding newInterval.
Note: Intervals are non-overlapping if they have no common point. For example, [1,2] and [3,4] are non-overlapping, but [1,2] and [2,3] are overlapping.
*/

/*
Linear scan in three phases
The intervals are already sorted, so one pass is enough:
Add every interval that ends before newInterval starts.
Merge every interval that overlaps newInterval.
Add every interval that starts after newInterval ends.
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	output := make([][]int, 0, len(intervals)+1)
	i := 0
	n := len(intervals)

	// 1. Intervals completely before newInterval
	for i < n && intervals[i][1] < newInterval[0] {
		output = append(output, intervals[i])
		i++
	}

	// 2. Merge all overlapping intervals
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	output = append(output, newInterval)

	// 3. Intervals completely after newInterval
	for i < n {
		output = append(output, intervals[i])
		i++
	}

	return output
}
