package main

/*
You are given two non-empty linked lists, l1 and l2, where each represents a non-negative integer.
The digits are stored in reverse order, e.g. the number 321 is represented as 1 -> 2 -> 3 -> in the linked list.
Each of the nodes contains a single digit. You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Return the sum of the two numbers as a linked list.
*/

/*We track the extra value, carry, here as well. We iterate through the lists l1 and l2 until both lists reach null. We add the values of both nodes as well as the carry. If either of the nodes is null, we add 0 in its place and continue the process while tracking the carry simultaneously. Once we complete the process, if we are left with any carry, we add an extra node with that carry value and return the head of the result list. */
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	// continue while either list has nodes or a carry remains
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next

		carry = sum / 10
	}
	return dummy.Next
}
