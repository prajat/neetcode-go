package main

import "fmt"

func main() {
	fmt.Println(FindMin([]int{3, 4, 5, 6, 1, 2}))
	fmt.Println(FindMin([]int{4, 5, 6, 7}))
	fmt.Println(FindMin([]int{4, 5, 0, 1, 2, 3}))
	fmt.Println(Search([]int{3, 4, 5, 6, 1, 2}, 1))
	fmt.Println(Search([]int{3, 5, 6, 0, 1, 2}, 4))
}
