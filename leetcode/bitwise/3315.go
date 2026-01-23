package bitwise

func Challenge3315(nums []int) []int {
	return minBitwiseArray(nums)
}

// Given n, find x that x | (x + 1) == n
// e.g., we have x = 1 0 0 0 1
// then      x + 1 = 1 0 0 1 0
// -> x | ( x + 1) = 1 0 0 1 1
// This infers that to get smallest x, when we have n, we need to unset the leftmost one in the trailing ones of n. (This means flip bit 1 to bit 0)
//
// AKA: 1 0 0 1 1 -> 1 0 0 0 1
// To change the 3rd bit to 0, there are no easy way to do it. But we can
// 1. find the 2nd bit
// 2. then shift right
//
// For the step 1, we can find the rightmost set bit for the inverse of n
// E.g.,                      n = 1 0 0 1 1
// The inverse of n will be:
//
//	~n = 0 1 1 0 0
//
// To find the rightmost set bit, we can use trick -q & q
//
//	       -(~n) = 1 0 1 0 0
//	-(~n) & (~n) = 0 0 1 0 0
//
// But in this case, we can simply replace -(~n) with (n + 1). So we have
//
//	(n + 1) & (~n) = 0 0 1 0 0
//
// Here, we finish step 1: find the 2nd bit.
//
// For the step 2, we shift right
//
//	((n + 1) & (~n)) >> 1 = 0 0 0 1 0
//
// Then we flip it, so the wanted position become 0
//
//	~(((n + 1) & (~n)) >> 1) = 1 1 1 0 1
//
// In the end, we can AND with n to get x
// ~(((n + 1) & (~n)) >> 1) & n = 1 0 0 0 1
func minBitwiseArray(nums []int) []int {
	rs := make([]int, len(nums))
	for i := range nums {
		n := nums[i]
		if n != 2 {
			rs[i] = ^(((n + 1) & (^n)) >> 1) & n
		} else {
			rs[i] = -1
		}
	}
	return rs
}
