package utils

// get ceiling round of a / b
func DivCeil(dividend, divisor int) int {
	return (dividend + divisor - 1) / divisor
}
