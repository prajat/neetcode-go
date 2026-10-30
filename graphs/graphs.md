## Graphs, briefly

A **graph** is a set of **nodes (vertices)** connected by **edges**.

Unlike a tree:

- there is no single root
- there can be cycles
- a node can have any number of neighbours
- the graph may be disconnected

A tree is a special graph: connected, acyclic, with one root.

---

## Two kinds you will see

### 1. Explicit graph

Nodes and edges are given directly:

```text
n = 4
edges = [[0,1], [1,2], [2,3]]
```

Usually stored as an **adjacency list**:

```go
graph := make([][]int, n)

for _, edge := range edges {
	a, b := edge[0], edge[1]
	graph[a] = append(graph[a], b)
	graph[b] = append(graph[b], a) // undirected
}
```

### 2. Implicit graph

The graph is not given as edges. Neighbours are computed.

Most common: a **grid**.

```text
grid[r][c]
neighbours: up, down, left, right
```

`Number of Islands` is this type.

---

## Directed vs undirected

| Type       | Meaning             | Coding implication           |
| ---------- | ------------------- | ---------------------------- |
| Undirected | Edge goes both ways | Add both `a → b` and `b → a` |
| Directed   | Edge goes one way   | Add only `a → b`             |

Interview default:

- islands / maze / matrix → undirected
- course schedule / prerequisites → directed
- shortest path in unweighted graph → BFS

---

## DFS vs BFS

### DFS — go deep

Use for:

- connected components
- islands
- cycle detection
- topological sort
- path existence
- backtracking on graphs

Grid DFS template:

```go
func dfs(row, col int) {
	if outOfBounds || visited {
		return
	}

	markVisited()

	dfs(row-1, col)
	dfs(row+1, col)
	dfs(row, col-1)
	dfs(row, col+1)
}
```

### BFS — go level by level

Use for:

- shortest path in an **unweighted** graph
- nearest distance
- rotting oranges
- word ladder
- level-order graph problems

Queue template:

```go
queue := []int{start}
visited[start] = true

for len(queue) > 0 {
	node := queue[0]
	queue = queue[1:]

	for _, neighbour := range graph[node] {
		if visited[neighbour] {
			continue
		}

		visited[neighbour] = true
		queue = append(queue, neighbour)
	}
}
```

Rule of thumb:

- “count connected groups” → DFS or BFS
- “shortest number of steps” → BFS
- “can we finish all courses?” → directed graph + cycle detection / topological sort

---

## The most important idea: visited

Graphs can have cycles. Without `visited`, DFS/BFS can loop forever.

In a grid, two common options:

```go
visited[row][col] = true
```

or mutate in place:

```go
grid[row][col] = '0'
```

In an adjacency-list graph:

```go
visited := make([]bool, n)
```

---

## Common interview patterns

| Pattern                   | Example                                   |
| ------------------------- | ----------------------------------------- |
| Connected components      | Number of Islands                         |
| Multi-source BFS          | Rotting Oranges                           |
| Cycle in undirected graph | Graph Valid Tree                          |
| Cycle in directed graph   | Course Schedule                           |
| Topological sort          | Course Schedule II                        |
| Shortest path, unweighted | Word Ladder, 01 Matrix                    |
| Clone / copy graph        | Clone Graph                               |
| Union-Find                | Number of Provinces, Redundant Connection |

---

## Complexity

For \(V\) nodes and \(E\) edges:

- DFS / BFS time: \(O(V + E)\)
- Adjacency list space: \(O(V + E)\)
- Grid with \(R \times C\) cells: \(O(RC)\)

For a grid, each cell is a node, and each cell has up to 4 edges.

---

## Mental model

A graph problem usually asks one of these:

1. Can I reach B from A?
2. How many separate groups are there?
3. What is the shortest path?
4. Is there a cycle?
5. In what order must I process nodes?

If the input is a grid, think:

```text
each cell = node
4-direction moves = edges
```

If the input is `n` and `edges`, build an adjacency list first, then DFS or BFS.
