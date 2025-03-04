package math_challenges

func Challenge1780(n int) bool {
	return checkPowersOfThree(n)
}

// Check if a number can be represented as a sum of distinct powers of 3
// Approach: Convert the number to base-3 and check if any digit is '2'
// If any digit is '2', return false. Else, return true. Proof:
// n = 3^x1 + 3^x2 + ... + 3^xk
// If n has a digit '2' in base-3, then n = 3^x1 + 3^x2 + ... + 3^xk = 3^x1 + 3^x2 + ... + 3^(xk-1) + 3^xk + 1
// Time complexity: O(logn)
func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}
