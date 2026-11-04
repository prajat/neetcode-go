package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, -9, -3, 5}))
	fmt.Println(wordBreak("neetcode", []string{"neet", "code"}))
	fmt.Println(wordBreak("catsincars", []string{"cats", "cat", "sin", "in", "car"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen", "ape"}))
}
