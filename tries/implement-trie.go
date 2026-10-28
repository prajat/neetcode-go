package main

/*A prefix tree (also known as a trie) is a tree data structure used to efficiently store and retrieve keys in a set of strings. Some applications of this data structure include auto-complete and spell checker systems.
Implement the PrefixTree class:
PrefixTree() Initializes the prefix tree object.
void insert(String word) Inserts the string word into the prefix tree.
boolean search(String word) Returns true if the string word is in the prefix tree (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/

type PrefixTree struct {
	// children for each possible character
	// end - whether a complete word ends at that node
	children [26]*PrefixTree
	isEnd    bool
}

func Constructor() PrefixTree {
	return PrefixTree{}
}

// O(L) where L = length of the word
func (this *PrefixTree) Insert(word string) {
	current := this

	for i := range word {
		index := word[i] - 'a'

		if current.children[index] == nil {
			current.children[index] = &PrefixTree{}
		}

		current = current.children[index]
	}
	current.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	current := this

	for i := range word {
		index := word[i] - 'a'

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}
	return current.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	current := this

	for i := range prefix {
		index := prefix[i] - 'a'

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}
	return true
}
