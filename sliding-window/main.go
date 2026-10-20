package main

import "fmt"

func main() {
	//fmt.Println(MaxProfit([]int{10, 1, 5, 6, 7, 1}))
	//fmt.Println(MaxProfit([]int{10, 8, 7, 5, 2}))
	//fmt.Println(LengthOfLongestSubstring("zxyzxy"))
	//fmt.Println(LengthOfLongestSubstring("xxxx"))
	//fmt.Println(LengthOfLongestSubstring("bwf"))
	//fmt.Println(CharacterReplacement("XYYX", 2))
	//fmt.Println(CharacterReplacement("AAABABB", 1))
	fmt.Println(MinWindow("OUZODYXAZV", "XYZ"))
	fmt.Println(MinWindow("xyz", "xyz"))
	fmt.Println(MinWindow("x", "xy"))
}
