package main

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	//remove whitespaces from the string s and non-alphanumeric chars
	strArr := strings.Fields(s)
	str := strings.ToLower(strings.Join(strArr, ""))
	str = removeNonAlphanumeric(str)
	//now check if str is palindrome by using two pointer approach
	//start from left and start from right and then check until they reached halfway.
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

func removeNonAlphanumeric(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}

		return -1 // Remove this character.
	}, s)
}

/*
// Ascci version
func isAlphanumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

func removeNonAlphanumeric(s string) string {
	result := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			result = append(result, s[i])
		}
	}

	return string(result)
}
*/
