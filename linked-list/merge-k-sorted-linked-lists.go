package main

import (
	"container/heap"
	"slices"
)

/*You are given an array of k linked lists lists, where each list is sorted in ascending order.
Return the sorted linked list that is the result of merging all of the individual linked lists. */

// Brute Force Method - extra space - O(n)
func mergeKLists(lists []*ListNode) *ListNode {
	// Collect all the node in an array and sort it
	arr := make([]int, 0)
	for i := range lists {
		current := lists[i]
		for current != nil {
			arr = append(arr, current.Val)
			current = current.Next
		}
	}
	slices.Sort(arr)
	// build the list back from arr
	return buildList(arr)
}

// Optimal Solution - without extra space and rearranging nodes using minHeap as helper
type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)

	node := old[n-1]
	*h = old[:n-1]

	return node
}

func mergeKLis	tsV2(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Add the head of every non-empty list.
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	for minHeap.Len() > 0 {
		// Get the node with the smallest value.
		node := heap.Pop(minHeap).(*ListNode)

		// Attach it to the result.
		tail.Next = node
		tail = tail.Next

		// Add the next node from the same list.
		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
	}

	return dummy.Next
}
