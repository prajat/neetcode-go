package main

/*
There is a new alien language that uses the English alphabet, but the order of the letters is unknown.
You are given a list of strings words from the alien language's dictionary. It is claimed that the strings in words are sorted lexicographically by the rules of this new language.
If this claim is incorrect, and the given arrangement of strings in words cannot correspond to any order of letters, return "".
Otherwise, return a string of the unique letters in the new alien language sorted in lexicographically increasing order by the new language's rules. If there are multiple solutions, return any of them.
A string a is lexicographically smaller than a string b if either of the following is true:
The first letter where they differ is smaller in a than in b.
a is a prefix of b and a.length < b.length.
*/

// Compare each pair of adjacent words. The first different character tells us the ordering between those two alien letters.
func foreignDictionary(words []string) string {
	graph := make(map[byte][]byte)
	indegree := make(map[byte]int)

	// add every character, icluding characters with no edges
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			ch := word[i]
			if _, exists := indegree[ch]; !exists {
				indegree[ch] = 0
			}
		}
	}

	// building ordering constraints from adjacent words
	for i := 0; i < len(words)-1; i++ {
		first := words[i]
		second := words[i+1]

		commonLength := min(len(first), len(second))
		foundDifference := false

		for j := 0; j < commonLength; j++ {
			if first[j] == second[j] {
				continue
			}

			from := first[j]
			to := second[j]

			// avoid counting the same edge more than once
			alreadyExists := false
			for _, neighbour := range graph[from] {
				if neighbour == to {
					alreadyExists = true
					break
				}
			}

			if !alreadyExists {
				graph[from] = append(graph[from], to)
				indegree[to]++
			}

			foundDifference = true
			break
		}

		// Invalid order: longer word appears before its prefix
		if !foundDifference && len(first) > len(second) {
			return ""
		}
	}

	state := make(map[byte]int)
	result := make([]byte, 0, len(indegree))

	var dfs func(byte) bool
	dfs = func(ch byte) bool {
		// currently visiting: cycle detected
		if state[ch] == 1 {
			return false
		}

		// already completely processed
		if state[ch] == 2 {
			return true
		}

		state[ch] = 1

		for _, neighbour := range graph[ch] {
			if !dfs(neighbour) {
				return false
			}
		}

		state[ch] = 2

		// postorder gives reverse topological order
		result = append(result, ch)

		return true
	}

	for ch := range indegree {
		if state[ch] == 0 {
			if !dfs(ch) {
				return ""
			}
		}
	}

	// reverse the postorder result
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return string(result)
}
