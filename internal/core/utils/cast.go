package utils

// map character to 26 integer array
// 'a' -> 0
// 'b' -> 1
// nearly similar to ord() in Python
func LowercaseToInt(chr byte) int {
	return int(chr - 'a')
}

// nearly similar to chr() in Python
func IntToLowercase(ord int) byte {
	return byte(ord + 'a')
}
