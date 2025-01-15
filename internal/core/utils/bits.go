package utils

// Awesome materials: https://graphics.stanford.edu/~seander/bithacks.html

func GetNbSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func UnsetRightMostSetBit(n int) int {
	return n & (n - 1)
}

// set the bit at pos in n to 1
// where pos is the position of bit starting from right
func SetFromRight(n, pos int) int {
	return n | (1 << pos)
}

func SetRightMostUnsetBit(n int) int {
	return n | (n + 1)
}

func UnsetFromRight(n, pos int) int {
	return n ^ (1 << pos)
}
