package main

/*Given the head of a linked list and an integer n, remove the nth node from the end of the list and return its head. */

// BruteForce Method - extra O(n) space
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := make([]int, 0)
	resultArr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	nthIndex := len(arr) - n
	for i := range arr {
		if i == nthIndex {
			continue
		}
		resultArr = append(resultArr, arr[i])
	}
	dummy := &ListNode{}
	current := dummy
	for i := range resultArr {
		current.Next = &ListNode{Val: resultArr[i]}
		current = current.Next
	}
	return dummy.Next
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	// use two pass approach to find the length of Linked list(N)
	length := 0
	for current := head; current != nil; current = current.Next {
		length++
	}

	// Dummy handles removing the head node
	dummy := &ListNode{Next: head}

	// The node before the target node
	previous := dummy

	// Move the pointer to nth index from end by going to N-n index from start
	for i := 0; i < length-n; i++ {
		previous = previous.Next
	}
	// Remove the nth node from end
	previous.Next = previous.Next.Next

	return dummy.Next
}
