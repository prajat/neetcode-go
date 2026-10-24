package main

import "fmt"

func main() {
	fmt.Println(IsPalindrome("Was it a car or a cat I saw?"))
	fmt.Println(TwoSum([]int{1, 2, 3, 4}, 3))
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(MaxArea([]int{1, 7, 2, 5, 4, 7, 3, 6}))
	fmt.Println(MaxArea([]int{2, 2, 2}))
	fmt.Println(mergeSort([]int{5, 2, 8, 1, 3}))
}
