package main

import (
	"strconv"
	"strings"
)

type Solution struct{}

// //naive solution, wont work for all ascii chararcters
// func (s *Solution) Encode(strs []string) string {
// 	return strings.Join(strs, "-")
// }

// func (s *Solution) Decode(encoded string) []string {
// 	return strings.Split(encoded, "-")
// }

// here we use enoding = <lengthOfStr>Str<><"#">
func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length

		result = append(result, encoded[start:end])
		i = end
	}

	return result
}
