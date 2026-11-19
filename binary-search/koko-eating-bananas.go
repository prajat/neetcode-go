package main

/*
You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.
You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.
Return the minimum integer k such that you can eat all the bananas within h hours.
*/

/*
Why binary search works
If Koko can finish at speed k, she can also finish at any faster speed. So the feasible speeds form a pattern like this:
text
k:        1  2  3  4  5  6 ...
feasible: N  N  N  Y  Y  Y ...
You want the first Y, which is the smallest feasible k.
Lowest possible speed: 1
Highest needed speed: max(piles). At that speed, each pile takes exactly 1 hour, which is the fastest possible.
*/
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0

	for _, pile := range piles {
		right = max(right, pile)
	}

	for left < right {
		mid := left + (right-left)/2

		if canFinish(piles, h, mid) {
			// mid works; try slower speed
			right = mid
		} else {
			// mid is too slow
			left = mid + 1
		}
	}
	return left
}

func canFinish(piles []int, h, speed int) bool {
	hours := 0

	for _, pile := range piles {
		// cieling division: hours needed for this pile
		hours += (pile + speed - 1) / speed

		if hours > h {
			return false
		}
	}
	return true
}
