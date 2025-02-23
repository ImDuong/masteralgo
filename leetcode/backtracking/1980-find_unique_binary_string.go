package backtracking

func Challenge1980(nums []string) string {
	return findDifferentBinaryString(nums)
}

// The itituition is just make sure our string is different from each given string by only one bit
// So we can just iterate through the given strings and flip bit at i-th position
// -> This is Cantor's diagonal argument
func findDifferentBinaryString(nums []string) string {
	rs := make([]byte, len(nums[0]))
	for i := range len(nums[0]) {
		rs[i] = flip(nums[i][i])
	}
	return string(rs)
}

func flip(a byte) byte {
	if a == '0' {
		return '1'
	}
	return '0'
}
