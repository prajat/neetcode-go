package main

/*You are given the head of a linked list of length n. Unlike a singly linked list, each node contains an additional pointer random, which may point to any node in the list, or null.
Create a deep copy of the list.
The deep copy should consist of exactly n new nodes, each including:
The original value val of the copied node
A next pointer to the new node corresponding to the next pointer of the original node
A random pointer to the new node corresponding to the random pointer of the original node
Note: None of the pointers in the new list should point to nodes in the original list.
Return the head of the copied linked list.
In the examples, the linked list is represented as a list of n nodes. Each node is represented as a pair of [val, random_index] where random_index is the index of the node (0-indexed) that the random pointer points to, or null if it does not point to any node. */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// hashMap for mapping the original nodes to their corrsponding copies
	mapping := make(map[*Node]*Node)

	// Pass 1: create a copy of every node (values only)
	for current := head; current != nil; current = current.Next {
		mapping[current] = &Node{Val: current.Val}
	}

	// Pass 2: connect the copies' pointers
	for current := head; current != nil; current = current.Next {
		copied := mapping[current]
		copied.Next = mapping[current.Next]
		copied.Random = mapping[current.Random]
	}
	return mapping[head]
}
