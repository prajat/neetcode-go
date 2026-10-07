package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapOfS := make(map[rune]int, 0)
	mapOfT := make(map[rune]int, 0)

	//Iterate over the strings and build maps
	for _, ch := range s {
		mapOfS[ch] = mapOfS[ch] + 1
	}

	for _, ch := range t {
		mapOfT[ch] = mapOfT[ch] + 1
	}

	//Iterate over maps to check if they are equal or not
	for k, v := range mapOfS {
		if mapOfT[k] != v {
			return false
		}
	}

	return true
}
