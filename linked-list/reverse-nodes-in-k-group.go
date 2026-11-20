package main

/*
You are given the head of a singly linked list head and a positive integer k.
You must reverse the first k nodes in the linked list, and then reverse the next k nodes, and so on. If there are fewer than k nodes left, leave the nodes as they are.
Return the modified list after reversing the nodes in each group of k.
You are only allowed to modify the nodes' next pointers, not the values of the nodes.
*/

/*
Iterative solution with a dummy node
For each group:
Check that k nodes remain.
Reverse those k nodes.
Reconnect the previous group's tail to the new head, and the reversed group's tail to the next group.
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy // tail of the previous group

	for {
		// find the k-th node of this group
		kth := getKth(groupPrev, k)
		if kth == nil {
			// fewer than k node remain: leave them as they are
			break
		}

		groupNext := kth.Next

		// Reverse the group. The first node's next becomes groupNext
		prev := groupNext
		current := groupPrev.Next

		for current != groupNext {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
		}

		// Reconnect: the old first node is not the group's tail
		oldGroupHead := groupPrev.Next
		groupPrev.Next = kth // kth is the new head of the group
		groupPrev = oldGroupHead
	}
	return dummy.Next
}

// Returns the k-th node after ₹node, or nil if the list is too short
func getKth(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}
