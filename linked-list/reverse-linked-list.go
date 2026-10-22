package main

/*Given the beginning of a singly linked list head, reverse the list, and return the new beginning of the list. */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Optimal Solution - reverse in place - O(n)
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
