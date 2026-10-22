package main

/*You are given the head of a singly linked-list.
The positions of a linked list of length = 7 for example, can intially be represented as:
[0, 1, 2, 3, 4, 5, 6]
Reorder the nodes of the linked list to be in the following order:
[0, 6, 1, 5, 2, 4, 3]
In the general case, label the nodes by their original zero-based positions from 0 to n - 1. After reordering, those original positions appear in this order:
[0, n-1, 1, n-2, 2, n-3, ...]
These numbers represent node positions, not the values stored in the nodes.
You may not modify the values in the list's nodes, but instead you must reorder the nodes themselves. */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle of the list.
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split the list into two halves.
	second := slow.Next
	slow.Next = nil

	// Reverse the second half.
	second = reverseList(second)

	// Merge both halves alternately.
	first := head

	for second != nil {
		firstNext := first.Next
		secondNext := second.Next

		first.Next = second
		second.Next = firstNext

		first = firstNext
		second = secondNext
	}
}
