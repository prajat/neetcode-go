package main

import (
	"strconv"
	"strings"
)

/*Implement an algorithm to serialize and deserialize a binary tree.
Serialization is the process of converting an in-memory structure into a sequence of bits so that it can be stored or sent across a network to be reconstructed later in another computer environment.
You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure. There is no additional restriction on how your serialization/deserialization algorithm should work. */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	values := make([]string, 0)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			values = append(values, "#")
			return
		}

		values = append(values, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return strings.Join(values, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index >= len(values) || values[index] == "#" {
			index++
			return nil
		}

		value, _ := strconv.Atoi(values[index])
		index++

		node := &TreeNode{Val: value}
		node.Left = dfs()
		node.Right = dfs()

		return node
	}
	return dfs()
}
