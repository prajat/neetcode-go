package main

/*Given a string s, find the length of the longest substring without duplicate characters.
A substring is a contiguous sequence of characters within a string. */

// Brute Force Method -> O(n2)
// A brute force solution would be to try the substring starting at index i and try to find the maximum length we can form without duplicates by starting at that index.
// We can use a hash set to detect duplicates in O(1) time.
func lengthOfLongestSubstring(s string) int {
	maxLength := 1
	for i := range s {
		hashSet := make(map[byte]bool, 0)
		length := 1
		hashSet[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if hashSet[s[j]] {
				break
			}
			hashSet[s[j]] = true
			length++
		}
		maxLength = max(length, maxLength)
	}
	return maxLength
}

// Using sliding window approach - O(n)
func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0
	right := 0
	lastSeen := make(map[byte]int, 0)
	for right < len(s) {
		if previousIndex, exists := lastSeen[s[right]]; exists {
			left = max(left, previousIndex+1)
		}
		lastSeen[s[right]] = right
		currLength := right - left + 1
		maxLength = max(maxLength, currLength)
		right++
	}
	return maxLength
}
