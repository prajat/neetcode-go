package main

import "container/heap"

/*
You are given an array of integers stones where stones[i] represents the weight of the ith stone.
We want to run a simulation on the stones as follows:
At each step we choose the two heaviest stones, with weight x and y and smash them togethers
If x == y, both stones are destroyed
If x < y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
Continue the simulation until there is no more than one stone remaining.
Return the weight of the last remaining stone or return 0 if none remain.
*/

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)

	value := old[n-1]
	*h = old[:n-1]
	return value
}

func lastStoneWeight(stones []int) int {
	h := &maxHeap{}
	heap.Init(h)

	// insert all wieghts into max-heap
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	// simulate the process until one or no element remains in the max-heap
	for h.Len() >= 2 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)

		if x == y {
			continue
		}

		heap.Push(h, x-y)
	}

	if h.Len() == 0 {
		return 0
	} else {
		return heap.Pop(h).(int)
	}
}
