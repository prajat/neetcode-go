package main

import "container/heap"

/*
You are given an 2-D array points where points[i] = [xi, yi] represents the coordinates of a point on an X-Y axis plane. You are also given an integer k.
Return the k closest points to the origin (0, 0).
The distance between two points is defined as the Euclidean distance (sqrt((x1 - x2)^2 + (y1 - y2)^2)).
You may return the answer in any order. The answer is guaranteed to be unique(except for the order in which the points are returned.)
*/

type mh [][]int

func (h mh) Len() int {
	return len(h)
}

func (h mh) Less(i, j int) bool {
	// logic to order points based on their distance from origin
	return distanceSq(h[i]) > distanceSq(h[j])
}

func (h mh) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *mh) Push(value any) {
	*h = append(*h, value.([]int))
}

func (h *mh) Pop() any {
	old := *h
	n := len(old)
	value := old[n-1]
	*h = old[:n-1]
	return value
}

func distanceSq(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}

func kClosest(points [][]int, k int) [][]int {
	h := &mh{}
	heap.Init(h)

	for _, p := range points {
		heap.Push(h, p)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([][]int, h.Len())
	for i := 0; i < len(result); i++ {
		result[i] = heap.Pop(h).([]int)
	}

	return result
}
