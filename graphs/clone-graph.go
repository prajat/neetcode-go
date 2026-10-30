package main

/*Given a node in a connected undirected graph, return a deep copy of the graph.
Each node in the graph contains an integer value and a list of its neighbors.
class Node {
    public int val;
    public List<Node> neighbors;
}
The graph is shown in the test cases as an adjacency list. An adjacency list is a mapping of nodes to lists, used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
For simplicity, nodes values are numbered from 1 to n, where n is the total number of nodes in the graph. The index of each node within the adjacency list is the same as the node's value (1-indexed).
The input node will always be the first node in the graph and have 1 as the value. */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// map original node its cloned node
	clones := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(current *Node) *Node {
		if current == nil {
			return nil
		}

		// return the existing clone if this node was already visited
		if clone, exists := clones[current]; exists {
			return clone
		}

		// create and store the clone before visiting neighbours
		clone := &Node{Val: current.Val}
		clones[current] = clone

		// clone every neighbour
		for _, neighbour := range current.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbour))
		}
		return clone
	}
	return dfs(node)
}
