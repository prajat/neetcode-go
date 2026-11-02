package main

/*
Given a string s, return the number of substrings within s that are palindromes.
A palindrome is a string that reads the same forward and backward.
*/
func countSubstrings(s string) int {
	count := 0
	for i := range s {
		// odd-length palindromes
		count += expand(s, i, i)

		// even-length palindromes
		count += expand(s, i, i+1)
	}
	return count
}

func expand(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++

		left--
		right++
	}
	return count
}
