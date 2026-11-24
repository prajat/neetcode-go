package main

import "container/heap"

/*
Given an unsorted array of integers nums and an integer k, return the kth largest element in the array.
By kth largest element, we mean the kth largest element in the sorted order, not the kth distinct element.
Follow-up: Can you solve it without sorting?
*/
func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}
