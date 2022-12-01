package algoI

// straight forward solution from fabrizio3

func HammingWeight(num uint32) int {
	nb1Bits := 0
	for num != 0 {
		nb1Bits = nb1Bits + int((num & 1))
		num = num >> 1
	}
	return nb1Bits
}

// quoted by steve027, the implementation is discussed in chapter 5 in Hacker's Delight book: https://doc.lagout.org/security/Hackers%20Delight.pdf
func HammingWeightV2(num uint32) int {
	num = num - ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num + (num >> 4)) & 0x0f0f0f0f
	num = num + (num >> 8)
	num = num + (num >> 16)
	return int(num & 0x3f)
}
