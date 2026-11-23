package main

/*
You are given a string digits made up of digits from 2 through 9 inclusive.
Each digit (not including 1) is mapped to a set of characters as shown below:
A digit could represent any one of the characters it maps to.
Return all possible letter combinations that digits could represent. You may return the answer in any order.
*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	hashmap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := make([]string, 0)
	path := make([]byte, 0, len(digits))

	var dfs func(int)
	dfs = func(index int) {
		if index == len(digits) {
			result = append(result, string(path))
			return
		}

		letters := hashmap[digits[index]]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			dfs(index + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}
