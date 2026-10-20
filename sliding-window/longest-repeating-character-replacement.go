package main

/*
You are given a string s consisting of only uppercase English characters and an integer k.
You can choose up to k characters of the string and replace them with any other uppercase
English character.

After performing at most k replacements, return the length of the longest substring which
contains only one distinct character.
*/

/*
Hints - For a given substring or window, it is always optimal to replace characters with the most frequent character in that substring. Why? Because using the most frequent character minimizes the number of replacements required to make all characters in the substring identical. How can you find the number of replacements now?
The number of replacements is equal to the difference between the length of the current substring and the frequency of the most frequent character in that substring. A brute force solution would be to consider all substrings, use a hash map for frequency counting, and return the maximum length of the substring that has at most k replacements. This would be an O(n^2) solution. Can you think of a better way?
We can use the sliding window approach. The window size will be dynamic, and we will shrink the window when the number of replacements exceeds k. The result will be the maximum window size observed at each iteration.
*/

func CharacterReplacement(s string, k int) int {
	frequency := make(map[byte]int)
	left := 0
	right := 0
	maxFrequency := 0
	maxLength := 0

	for right < len(s) {
		frequency[s[right]]++
		if frequency[s[right]] > maxFrequency {
			maxFrequency = frequency[s[right]]
		}

		windowLength := right - left + 1
		replacements := windowLength - maxFrequency

		// If more than k replacements are needed, shrink the window.
		if replacements > k {
			frequency[s[left]]--
			left++
			windowLength = right - left + 1
		}

		if windowLength > maxLength {
			maxLength = windowLength
		}
		right++
	}

	return maxLength
}
