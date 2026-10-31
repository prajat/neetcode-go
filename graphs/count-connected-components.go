package main

// Number of Connected Components in an Undirected Graph
/*You have an undirected graph of n nodes labeled from 0 to n - 1. You are given an integer n and an array edges where edges[i] = [aᵢ, bᵢ] indicates that there is an edge between aᵢ and bᵢ in the graph.
Return the number of connected components in the graph. */

// Solution - Using Union-Find (DSU) algorithm
func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)

	// initially, every node is its own parent and component
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	// we have n nodes so n components in start
	components := n

	var find func(int) int
	find = func(node int) int {
		if parent[node] != node {
			// path compression
			parent[node] = find(parent[node])
		}
		return parent[node]
	}

	var union func(int, int)
	union = func(a, b int) {
		rootA := find(a)
		rootB := find(b)

		// already in same component
		if rootA == rootB {
			return
		}

		// attach the smaller tree under the target tree
		if size[rootA] < size[rootB] {
			rootA, rootB = rootB, rootA
		}

		parent[rootB] = rootA
		size[rootA] += size[rootB]

		components--
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	return components
}
