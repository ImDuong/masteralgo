package bitwise

import "masteralgo/internal/core/utils"

func Challenge2429(num1 int, num2 int) int {
	return minimizeXor(num1, num2)
}

func minimizeXor(num1 int, num2 int) int {
	n1, n2 := utils.GetNbSetBits(num1), utils.GetNbSetBits(num2)
	if n1 == n2 {
		return num1
	}
	rs := num1
	for n1 > n2 {
		// unset rightmost set bit
		rs = rs & (rs - 1)
		n1--
	}

	for n1 < n2 {
		// set rightmost unset bit
		rs = rs | (rs + 1)
		n1++
	}
	return rs
}
