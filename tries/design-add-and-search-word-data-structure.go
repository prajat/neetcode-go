package main

/*Design a data structure that supports adding new words and searching for existing words.
Implement the WordDictionary class:
void addWord(word) Adds word to the data structure.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter. */

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	current := this

	for i := range word {
		index := word[i] - 'a'

		if current.children[index] == nil {
			current.children[index] = &WordDictionary{}
		}

		current = current.children[index]
	}
	current.isEnd = true
}

// Use DFS for Search because '.' can match any character.
func (this *WordDictionary) Search(word string) bool {
	var dfs func(*WordDictionary, int) bool

	dfs = func(node *WordDictionary, index int) bool {
		if node == nil {
			return false
		}

		// All characters matched
		if index == len(word) {
			return node.isEnd
		}

		ch := word[index]
		if ch == '.' {
			// Try every possible child
			for _, child := range node.children {
				if child != nil && dfs(child, index+1) {
					return true
				}
			}
			return false
		}

		childIndex := ch - 'a'
		if node.children[childIndex] == nil {
			return false
		}
		return dfs(node.children[childIndex], index+1)
	}
	return dfs(this, 0)
}
