package main

import "container/heap"

/*
You are given an array of integers nums and an integer k. There is a sliding window of size k that starts at the left edge of the array. The window slides one position to the right until it reaches the right edge of the array.
Return a list that contains the maximum element in the window at each step.
*/

/*
Max-heap with lazy deletion
Store (value, index) pairs in a max-heap. The top is the largest value, but it may have slid out of the window. When you read the max, pop the top while its index is outside the window (index <= right - k).
*/

type Item struct {
	value int
	index int
}

type MaxHeap []Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]
	return item
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	h := &MaxHeap{}
	result := make([]int, 0, len(nums)-k+1)

	for right := 0; right < len(nums); right++ {
		heap.Push(h, Item{value: nums[right], index: right})

		// window is not full yet
		if right < k-1 {
			continue
		}

		// Remove maximum that are outside the window
		for (*h)[0].index <= right-k {
			heap.Pop(h)
		}

		result = append(result, (*h)[0].value)
	}
	return result
}
