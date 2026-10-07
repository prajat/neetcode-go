package main

import "slices"

func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		key := string(b)
		m[key] = append(m[key], s)
	}

	output := make([][]string, 0, len(m))
	for _, group := range m {
		output = append(output, group)
	}
	return output
}
