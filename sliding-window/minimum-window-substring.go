package main

/*
Given two strings s and t, return the shortest substring of s such that every character in t, including duplicates, is present in the substring. If such a substring does not exist, return an empty string "".
You may assume that the correct output is always unique.
*/

/*
HINTS:
We need to find substrings in s that should have atleast the characters of t. We can use hash maps to maintain the frequencies of characters. It will be O(1) for lookups.
We can use a dynamically sized sliding window approach on s. We iterate through s while maintaining a window. If the current window contains at least the frequency of characters from t, we update the result and shrink the window until it is valid.
We should ensure that we maintain the result substring and only update it if we find a shorter valid substring. Additionally, we need to keep track of the result substring's length so that we can return an empty string if no valid substring is found.
*/

func MinWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Count the characters required from t.
	required := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}

	window := make(map[byte]int)

	// Number of unique characters whose required count is satisfied.
	formed := 0
	need := len(required)

	left := 0
	bestStart := 0
	bestLength := len(s) + 1

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if required[ch] > 0 {
			window[ch]++

			if window[ch] == required[ch] {
				formed++
			}
		}

		// The current window contains all characters from t.
		for formed == need {
			windowLength := right - left + 1

			if windowLength < bestLength {
				bestLength = windowLength
				bestStart = left
			}

			leftChar := s[left]

			if required[leftChar] > 0 {
				window[leftChar]--

				if window[leftChar] < required[leftChar] {
					formed--
				}
			}

			left++
		}
	}

	if bestLength == len(s)+1 {
		return ""
	}

	return s[bestStart : bestStart+bestLength]
}
