package main

/*
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.
The input string s is valid if and only if:
Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.
*/

func isValid(s string) bool {
	stack := make([]byte, 0)

	// Map each closing bracket to its matching opening one
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if opening, exists := pairs[ch]; exists {
			// A closing bracket must match the latest opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}

			// remove the matching opening bracket.
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	// Valid only if every opening bracket was matched and stack is empty now
	return len(stack) == 0
}
