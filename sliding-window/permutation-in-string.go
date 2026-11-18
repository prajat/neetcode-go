package main

/*
You are given two strings s1 and s2.
Return true if s2 contains a permutation of s1, or false otherwise. That means if a permutation of s1 exists as a substring of s2, then return true.
Both strings only contain lowercase letters.
*/
func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	frequency := make([]int, 26)
	windowSize := len(s1)

	// count the letters s1 needs
	for i := 0; i < windowSize; i++ {
		frequency[s1[i]-'a']++
	}

	for right := 0; right < len(s2); right++ {
		// add new character; it uses us one needed letter
		frequency[s2[right]-'a']--

		// remove the character that just left the window
		if right >= windowSize {
			frequency[s2[right-windowSize]-'a']++
		}

		// all zeroes means the window has the same count as s1
		if right >= windowSize-1 && allZero(frequency) {
			return true
		}
	}
	return false
}

func allZero(frequency []int) bool {
	for _, count := range frequency {
		if count != 0 {
			return false
		}
	}
	return true
}
