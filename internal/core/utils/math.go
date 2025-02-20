package utils

// get ceiling round of a / b
func DivCeil(dividend, divisor int) int {
	return (dividend + divisor - 1) / divisor
}

// works with both positive and negative
func Mod(a, b int) int {
	return (a%b + b) % b
}
