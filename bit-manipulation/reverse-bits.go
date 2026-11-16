package main

import "math/bits"

/*Given a 32-bit unsigned integer n, reverse the bits of the binary representation of n and return the result. */
func reverseBits(n int) int {
	result := 0

	for i := 1; i <= 32; i++ {
		// take the lowest bit of n and append it to the result
		bit := n & 1
		result = (result << 1) | bit

		// Move to the next bit of n
		n >>= 1
	}
	return result
}

// Using math/bits
func reverseBitsV2(n int) int {
	return int(bits.Reverse32(uint32(n)))
}
