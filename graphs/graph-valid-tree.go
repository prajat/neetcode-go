package main

/*Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree. */
// Your DFS needs an adjacency list and a parent parameter. The parent is necessary because, in an undirected graph, every edge appears in both directions.
func validTree(n int, edges [][]int) bool {
	// a tree with n nodes must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	// build graph as adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// visited map to detect cycle or not
	visited := make(map[int]bool)

	var dfs func(int, int) bool
	dfs = func(node int, parent int) bool {
		if visited[node] {
			return false
		}

		visited[node] = true

		for _, neighbour := range graph[node] {
			// ignore the edge back to the node we came from
			if neighbour == parent {
				continue
			}

			if !dfs(neighbour, node) {
				return false
			}
		}
		return true
	}

	// start dfs from node 0
	if !dfs(0, -1) {
		return false
	}

	// every node must be reachable from node 0 for a undirected tree
	return len(visited) == n
}
