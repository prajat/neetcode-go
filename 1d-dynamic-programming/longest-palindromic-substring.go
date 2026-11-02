package main

/*
Given a string s, return the longest substring of s that is a palindrome.
A palindrome is a string that reads the same forward and backward.
If there are multiple palindromic substrings that have the same length, return any one of them.
*/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// two pointer approach
	start, end := 0, 0

	for i := range s {
		// odd-length palindrome: center is s[i]
		oddLength := expandFromCenter(s, i, i)

		// even-length palindromwe: center is between s[i] and s[i+1]
		evenLength := expandFromCenter(s, i, i+1)

		length := max(oddLength, evenLength)
		if length > end-start+1 {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	// at this point, left and right are one position outside the palindrome
	return right - left - 1
}
