package main

/*
You are given an array prerequisites where prerequisites[i] = [a, b] indicates that you must take course b first if you want to take course a.
The pair [0, 1], indicates that must take course 1 before taking course 0.
There are a total of numCourses courses you are required to take, labeled from 0 to numCourses - 1.
Return true if it is possible to finish all courses, otherwise return false.
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	// [a, b] means take b before a.
	// Therefore add an edge b -> a
	for _, prerequisite := range prerequisites {
		course := prerequisite[0]
		required := prerequisite[1]
		graph[required] = append(graph[required], course)
	}

	state := make([]int, numCourses)

	var dfs func(int) bool
	dfs = func(course int) bool {
		// course is already in current dfs path - cycle
		if state[course] == 1 {
			return false
		}

		// course was completely checked and has no cycle
		if state[course] == 2 {
			return true
		}

		// mark as currently visiting
		state[course] = 1

		for _, nextCourse := range graph[course] {
			if !dfs(nextCourse) {
				return false
			}
		}

		// mark as completely processed
		state[course] = 2

		return true
	}

	for course := range numCourses {
		if state[course] == 0 {
			if !dfs(course) {
				return false
			}
		}
	}

	return true
}
