package main

import "slices"

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

// Optimal Solution - without extra space and rearranging nodes using merge sort approach
func mergeKListsV2(lists []*ListNode) *ListNode {

}
