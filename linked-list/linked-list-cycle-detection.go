package main

/*Given the beginning of a linked list head, return true if there is a cycle in the linked list. Otherwise, return false.
There is a cycle in a linked list if at least one node in the list can be visited again by following the next pointer.
Internally, index determines the index of the beginning of the cycle, if it exists. The tail node of the list will set it's next pointer to the index-th node. If index = -1, then the tail node points to null and no cycle exists. */

func hasCycle(head *ListNode) bool {
	hashSet := make(map[int]bool, 0)
	for head != nil {
		if hashSet[head.Val] && head.Next != nil {
			return true
		} else {
			hashSet[head.Val] = true
		}
		head = head.Next
	}
	return false
}
