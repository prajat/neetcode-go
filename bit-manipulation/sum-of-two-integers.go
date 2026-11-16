package main

/*Given two integers a and b, return the sum of the two integers without using the + and - operators. */
/*Idea: addition with XOR and AND
For each bit position:
a ^ b is the sum without carry.
(a & b) << 1 is the carry, shifted to the next bit.
Repeat until there is no carry left. */
func getSum(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}
