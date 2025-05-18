package dynamicprogramming

func Challenge1411(n int) int {
	return numOfWays(n)
}

// in each row, there are 2 patterns: 121 and 123
// where 121 meaning first color is the same with the 3rd color
// Init: pattern 121 can be: 121, 131, 212, 232, 313, 323
//
//	pattern 123 can be: 123, 132, 213, 231, 312, 321
//	-> a121, a123 = 6, 6
//
// In the next row
// pattern 121 flw by: 212, 213, 232, 312, 313 -> 3 patterns 121 & 2 patterns 123
// pattern 123 flw by: 212, 231, 232, 312 -> 2 patterns 121 & 2 patterns 123
// -> formula: b121 = 3 * a121 + 2 * a123
//
//	b123 = 2 * a121 + 2 * a123
//
// The final result would be (b121 + b123) % mod
func numOfWays(n int) int {
	a121, a123, b121, b123 := 6, 6, 6, 6
	const mod = 1_000_000_007
	for i := 1; i < n; i++ {
		b121 = (3*a121 + 2*a123) % mod
		b123 = (2*a121 + 2*a123) % mod
		a121, a123 = b121, b123
	}
	return (b121 + b123) % mod
}
