package main

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of dictionary words.
You are allowed to reuse words in the dictionary an unlimited number of times. You may assume all dictionary words are unique.
*/

// using recursion only
func wordBreak(s string, wordDict []string) bool {
	// build wordSet for easy checking
	wordSet := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		wordSet[wordDict[i]] = true
	}

	var dfs func(int) bool
	dfs = func(index int) bool {
		// base condition - if reached at the end of string.
		if index == len(s) {
			return true
		}

		for end := index + 1; end <= len(s); end++ {
			word := s[index:end]
			if wordSet[word] && dfs(end) {
				return true
			}
		}
		return false
	}
	return dfs(0)
}

// recursion + memoisation
func wordBreakV2(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// -1 means not calculated
	// 0 means false
	// 1 means true
	cache := make([]int8, len(s)+1)
	for i := range cache {
		cache[i] = -1
	}

	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(s) {
			return true
		}

		if cache[index] != -1 {
			return cache[index] == 1
		}

		for end := index + 1; end <= len(s); end++ {
			word := s[index:end]
			if wordSet[word] && dfs(end) {
				cache[index] = 1
				return true
			}
		}

		cache[index] = 0
		return false
	}
	return dfs(0)
}
