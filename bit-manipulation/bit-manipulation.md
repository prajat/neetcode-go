## Core operators in Go

| Operator   | Meaning                         | Example                    |
| ---------- | ------------------------------- | -------------------------- |
| `&`        | AND                             | `6 & 3` = `2`              |
| `\|`       | OR                              | `6 \| 3` = `7`             |
| `^`        | XOR (or bitwise NOT when unary) | `6 ^ 3` = `5`, `^0` = `-1` |
| `&^`       | AND NOT (clear bits)            | `6 &^ 2` = `4`             |
| `<<`, `>>` | Shift left / right              | `1 << 3` = `8`             |

Go's `&` binds tighter than `|` and `^`, so `a | b & c` means `a | (b & c)`. Parenthesise anyway to make it clear.

## Essential patterns

### Check, set, clear, toggle bit `i`

```go
bitIsSet := (n>>i)&1 == 1
setBit    := n | (1 << i)
clearBit  := n &^ (1 << i)
toggleBit := n ^ (1 << i)
```

### Lowest set bit and clearing it

```go
lowest := n & -n      // isolate lowest 1 bit
n = n & (n - 1)       // clear lowest 1 bit (Brian Kernighan)
```

Repeating `n & (n-1)` until `n == 0` counts set bits.

### Power of two

```go
isPowerOfTwo := n > 0 && n&(n-1) == 0
```

### XOR properties

- `a ^ a = 0`
- `a ^ 0 = a`
- XOR is commutative and associative

So XOR cancels duplicates. This is the basis of "find the element that appears once."

## Standard library: `math/bits`

```go
import "math/bits"

bits.OnesCount(uint(n))   // number of 1 bits
bits.TrailingZeros(uint(n)) // zeros after the lowest 1
bits.LeadingZeros(uint(n))  // zeros before the highest 1
bits.Len(uint(n))          // bits needed to represent n
```

Use the `uint` versions for bit-count functions.

## Worked examples

### 1. Single Number (XOR everything)

```go
func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n // duplicates cancel out
	}
	return result
}
```

Time \(O(n)\), space \(O(1)\).

### 2. Number of 1 bits

```go
func hammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1 // remove lowest 1 bit
		count++
	}
	return count
}
```

### 3. Counting Bits (for 0 to n)

Use the fact that `i >> 1` removes the last bit:

```go
func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}
```

### 4. Missing Number (values 0 to n, one missing)

XOR the indices with the values. Pairs cancel, leaving the missing number.

```go
func missingNumber(nums []int) int {
	result := len(nums)
	for i, n := range nums {
		result ^= i ^ n
	}
	return result
}
```

### 5. Hamming distance

Differing bits are the 1s in `a ^ b`:

```go
distance := bits.OnesCount(uint(a ^ b))
```

### 6. Enumerate all subsets

Each bit in `mask` decides whether an element is included:

```go
func subsets(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	for mask := 0; mask < 1<<n; mask++ {
		subset := []int{}
		for j := 0; j < n; j++ {
			if (mask>>j)&1 == 1 {
				subset = append(subset, nums[j])
			}
		}
		result = append(result, subset)
	}
	return result
}
```

This gives \(2^n\) subsets.

## Pitfalls in Go

- **Signed shifts:** `>>` on a signed `int` is arithmetic, so negative numbers keep their sign bit. Use `uint` or `uint32` when you want logical shifts (e.g., reverse bits).
- **Overflow:** `1 << 63` overflows an `int` on 64-bit systems. Use `uint64` or `1 << uint(i)` with a bounded `i`.
- **Negative shift counts panic** at runtime, so keep shift amounts non-negative.
- **`&^` is AND-NOT**, a convenient way to clear bits. Go doesn't have `~` as binary NOT; use `^x` for unary complement.

## Quick reference

| Goal              | Expression                |
| ----------------- | ------------------------- |
| Is bit `i` set?   | `(n>>i)&1 == 1`           |
| Set bit `i`       | `n \| (1<<i)`             |
| Clear bit `i`     | `n &^ (1<<i)`             |
| Remove lowest 1   | `n & (n-1)`               |
| Power of two?     | `n > 0 && n&(n-1) == 0`   |
| Count 1s          | `bits.OnesCount(uint(n))` |
| Cancel duplicates | XOR all values            |
