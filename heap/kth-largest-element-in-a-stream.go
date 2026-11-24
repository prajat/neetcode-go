package main

import "container/heap"

/*Design a class to find the kth largest integer in a stream of values, including duplicates. E.g. the 2nd largest from [1, 2, 3, 3] is 3. The stream is not necessarily sorted.
Implement the following methods:
constructor(int k, int[] nums) Initializes the object given an integer k and the stream of integers nums.
int add(int val) Adds the integer val to the stream and returns the kth largest integer in the stream. */

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)

	value := old[n-1]
	*h = old[:n-1]
	return value
}

type KthLargest struct {
	h *minHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := &minHeap{}
	heap.Init(h)

	obj := KthLargest{
		h: h,
		k: k,
	}

	for _, num := range nums {
		obj.Add(num)
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)

	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}

	return (*this.h)[0]
}
