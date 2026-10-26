## Trees, briefly

A **tree** is a hierarchical structure of nodes. Each node has:

- a value
- zero or more children

There is **one root**. There are **no cycles**.

For coding interviews, you almost always use a **binary tree**:

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

Each node has at most two children: `Left` and `Right`.

A **binary search tree (BST)** is a binary tree with an extra rule:

```text
left subtree < node < right subtree
```

Most NeetCode tree problems are **binary trees**, not BSTs.

---

## Core idea

Almost every tree problem is:

1. Handle `nil` nodes.
2. Look at the current node.
3. Recurse on left and right.
4. Combine the results.

Template:

```go
func dfs(node *TreeNode) {
	if node == nil {
		return
	}

	// do something with node
	dfs(node.Left)
	dfs(node.Right)
}
```

Always check `nil` first. Missing this is the most common panic.

---

## Traversals

### 1. DFS — Preorder: node → left → right

Visit the node first.

```go
func preorder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	preorder(node.Left)
	preorder(node.Right)
}
```

Use for: cloning a tree, serializing, “process parent before children”.

### 2. DFS — Inorder: left → node → right

Visit left, then node, then right.

```go
func inorder(node *TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left)
	fmt.Println(node.Val)
	inorder(node.Right)
}
```

On a **BST**, inorder gives values in **sorted order**.

Use for: BST validation, kth smallest, converting BST to sorted list.

### 3. DFS — Postorder: left → right → node

Visit children first, then the node.

```go
func postorder(node *TreeNode) {
	if node == nil {
		return
	}

	postorder(node.Left)
	postorder(node.Right)
	fmt.Println(node.Val)
}
```

Use for: height, diameter, invert tree, delete nodes, “need child results before deciding”.

### 4. BFS — Level order

Visit level by level using a queue.

```go
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
```

Use for: level-order traversal, right side view, zigzag, min depth, “nearest” questions.

---

## Recursion vs BFS

| Need                             | Use             |
| -------------------------------- | --------------- |
| Compare structure / values       | Recursion (DFS) |
| Height, diameter, invert         | Postorder DFS   |
| Sorted BST values                | Inorder DFS     |
| Process by levels                | BFS queue       |
| Shortest path in unweighted tree | BFS             |

---

## Patterns you will reuse

**Same tree / subtree**

```go
if p == nil || q == nil {
	return p == q
}
return p.Val == q.Val &&
	isSameTree(p.Left, q.Left) &&
	isSameTree(p.Right, q.Right)
```

**Max depth**

```go
return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
```

**Invert tree**

```go
root.Left, root.Right = invert(root.Right), invert(root.Left)
```

**Height from a node** (postorder)

```go
left := height(node.Left)
right := height(node.Right)
return 1 + max(left, right)
```

---

## Complexity

For \(n\) nodes:

- Time is usually \(O(n)\) — visit each node once.
- Recursion space is \(O(h)\), where \(h\) is height.
  - Balanced tree: \(O(\log n)\)
  - Skewed tree (linked-list shaped): \(O(n)\)
- BFS queue space is \(O(w)\), where \(w\) is the widest level.

---

## Mental model

Think of each recursive call as answering one question about a subtree:

- “Are these two subtrees the same?”
- “What is this subtree’s height?”
- “Does this subtree contain the target?”

The current node combines answers from left and right.

That is most of tree interview problems.
