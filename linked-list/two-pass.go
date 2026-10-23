package main

func lengthOfListTwoPass(head *ListNode) int {
	// First pass: count the nodes.
	length := 0

	for current := head; current != nil; current = current.Next {
		length++
	}

	// Second pass: traverse the list again.
	current := head

	for i := 0; i < length && current != nil; i++ {
		// Use current here if needed.
		current = current.Next
	}

	return length
}
