package main

/*
A string consisting of uppercase english characters can be encoded to a number using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode a message, digits must be grouped and then mapped back into letters using the reverse of the mapping above. There may be multiple ways to decode a message. For example, "1012" can be mapped into:
"JAB" with the grouping (10 1 2)
"JL" with the grouping (10 12)
The grouping (1 01 2) is invalid because 01 cannot be mapped into a letter since it contains a leading zero.
Given a string s containing only digits, return the number of ways to decode it. You can assume that the answer fits in a 32-bit integer.
*/

// using recursion with memoisation
func numDecodings(s string) int {
	cache := make(map[int]int, 0)

	var dfs func(int) int
	dfs = func(index int) int {
		if index == len(s) {
			return 1
		}

		// substring starts with 0 are invaild
		if s[index] == '0' {
			return 0
		}

		if result, exists := cache[index]; exists {
			return result
		}

		// we can explore all possible decodings by combining one or two consecutive digits.
		// one digit: 1 to 9
		result := dfs(index + 1)

		// two digits: 10 to 26
		if index+1 < len(s) && (s[index] == '1' || (s[index] == '2' && s[index+1] <= '6')) {
			result += dfs(index + 2)
		}

		cache[index] = result
		return cache[index]
	}
	return dfs(0)
}

// only recursion - time limit exceeded when submitted
func numDecodingsV2(s string) int {
	var dfs func(int) int
	dfs = func(index int) int {
		if index == len(s) {
			return 1
		}

		if s[index] == '0' {
			return 0
		}

		// one-digit
		result := dfs(index + 1)

		// two-digits
		if index+1 < len(s) && (s[index] == '1' || (s[index] == '2' && s[index+1] <= '6')) {
			result += dfs(index + 2)
		}

		return result
	}
	return dfs(0)
}
