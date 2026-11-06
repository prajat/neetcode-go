package main

/*
Given two strings text1 and text2, return the length of the longest common subsequence between the two strings if one exists, otherwise return 0.
A subsequence is a sequence that can be derived from the given sequence by deleting some or no elements without changing the relative order of the remaining characters.
For example, "cat" is a subsequence of "crabt".
A common subsequence of two strings is a subsequence that exists in both strings.
*/

// using recursion only
func longestCommonSubsequence(text1, text2 string) int {
	var dfs func(int, int) int
	dfs = func(index1, index2 int) int {
		if index1 == len(text1) || index2 == len(text2) {
			return 0
		}

		if text1[index1] == text2[index2] {
			return 1 + dfs(index1+1, index2+1)
		}

		return max(dfs(index1+1, index2), dfs(index1, index2+1))
	}
	return dfs(0, 0)
}

// recursion + memoisation
func longestCommonSubsequenceV2(text1, text2 string) int {
	cache := make([][]int, len(text1))
	for i := range cache {
		cache[i] = make([]int, len(text2))
	}

	var dfs func(int, int) int
	dfs = func(index1, index2 int) int {
		if index1 == len(text1) || index2 == len(text2) {
			return 0
		}

		if cache[index1][index2] != 0 {
			return cache[index1][index2]
		}

		if text1[index1] == text2[index2] {
			return 1 + dfs(index1+1, index2+1)
		}

		cache[index1][index2] = max(dfs(index1+1, index2), dfs(index1, index2+1))
		return cache[index1][index2]
	}
	return dfs(0, 0)
}

// bottom-up DP
func longestCommonSubsequenceDP(text1, text2 string) int {
	return 1
}
