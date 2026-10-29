package main

import (
	"container/heap"
	"slices"
)

/*The median is the middle value in a sorted list of integers. For lists of even length, there is no middle value, so the median is the mean of the two middle values.
For example:
For arr = [1,2,3], the median is 2.
For arr = [1,2], the median is (1 + 2) / 2 = 1.5
Implement the MedianFinder class:
MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far. */

// Max Heap for the smaller half
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	value := old[n-1]
	*h = old[:n-1]

	return value
}

// Min heap for the larger half
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j] // imp for building max heap
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)

	value := old[n-1]
	*h = old[:n-1]

	return value
}

type MedianFinderV2 struct {
	lower *MaxHeap
	upper *MinHeap
}

func ConstructorMedianFinderV2() MedianFinderV2 {
	lower := &MaxHeap{}
	upper := &MinHeap{}

	heap.Init(lower)
	heap.Init(upper)

	return MedianFinderV2{
		lower: lower,
		upper: upper,
	}
}

func (this *MedianFinderV2) AddNum(num int) {
	if this.lower.Len() == 0 || num <= (*this.lower)[0] {
		heap.Push(this.lower, num)
	} else {
		heap.Push(this.upper, num)
	}

	// move the largest lower-half value to the upper half
	if this.lower.Len() > this.upper.Len()+1 {
		value := heap.Pop(this.lower)
		heap.Push(this.upper, value)
	}

	// move the smallest upper-half calue to the lower half
	if this.upper.Len() > this.lower.Len() {
		value := heap.Pop(this.upper)
		heap.Push(this.lower, value)
	}
}

func (this *MedianFinderV2) FindMedian() float64 {
	if this.lower.Len() > this.upper.Len() {
		return float64((*this.lower)[0])
	}

	left := (*this.lower)[0]
	right := (*this.upper)[0]

	return float64(left+right) / 2.0
}

// Naive Solution using array to store the nums and sort the array and return the array each time
type MedianFinder struct {
	Nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.Nums = append(this.Nums, num)
}

func (this *MedianFinder) FindMedian() float64 {
	var median float64
	if len(this.Nums) == 0 {
		return median
	}

	// sort the Nums first
	slices.Sort(this.Nums)
	size := len(this.Nums)

	if size%2 == 0 {
		// even sized array
		leftMid := this.Nums[size/2-1]
		rightMid := this.Nums[(size / 2)]
		median = (float64(leftMid) + float64(rightMid)) / 2
	} else {
		// odd sized array
		median = float64(this.Nums[size/2])
	}

	return median
}
