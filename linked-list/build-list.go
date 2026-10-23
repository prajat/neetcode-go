package main

// How to make a linked list from array and return the head of the linked list

// Using a dummy Node for help
func buildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	dummy := &ListNode{}
	current := dummy
	for i := range arr {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return dummy.Next
}

// without using helping dummy Node
func buildListV2(arr []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for i := range arr {
		node := &ListNode{Val: arr[i]}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}
