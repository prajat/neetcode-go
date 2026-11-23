package main

/*You are given an integer n. Return all well-formed parentheses strings that you can generate with n pairs of parentheses. */
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	path := ""
	var dfs func(int, int)
	dfs = func(open, close int) {
		if len(path) == 2*n {
			result = append(result, path)
			return
		}

		if open < n {
			path += "("
			dfs(open+1, close)
			path = path[:len(path)-1]
		}

		if close < open {
			path += ")"
			dfs(open, close+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return result
}
