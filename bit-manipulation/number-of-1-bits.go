package main

import (
	"strconv"
)

/*
You are given an unsigned integer n. Return the number of 1 bits in its binary representation.
You may assume n is a non-negative integer which fits within 32-bits.
*/

// naive solution using binary code conversion and counting 1's in it - O(n)
func hammingWeight(n int) int {
	count := 0
	binaryCode := strconv.FormatInt(int64(n), 2)
	//fmt.Println(binaryCode)
	for i := range binaryCode {
		if binaryCode[i] == '1' {
			count++
		}
	}
	return count
}

// bit - manipulation - O(1) space and time
// We iterate 32 times (0 to 31) using index i. The expression (1 << i) creates a bitmask with a set bit at the i-th position.
// Since the mask has a set bit at the i-th position and all 0s elsewhere, we can perform a bitwise-AND with n. If n has a set bit at the i-th position, the result is positive; otherwise, it is 0. We increment the global count if the result is positive and return it after the iteration.
func hammingWeightV2(n int) int {
	count := 0
	for i := 0; i < 32; i++ {
		if (1<<i)&n > 0 {
			count++
		}
	}
	return count
}
