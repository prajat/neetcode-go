package main

/*
Given a string s, split s into substrings where every substring is a palindrome. Return all possible lists of palindromic substrings.
You may return the solution in any order.
*/
func partition(s string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)

	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			list := make([]string, len(path))
			copy(list, path)
			result = append(result, list)
			return
		}

		for end := start + 1; end <= len(s); end++ {
			if isPalindrome(s[start:end]) {
				path = append(path, s[start:end])
				dfs(end)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return result
}

func isPalindrome(str string) bool {
	start := 0
	end := len(str) - 1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
