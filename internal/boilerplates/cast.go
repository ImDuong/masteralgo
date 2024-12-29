package boilerplates

// map character to 26 integer array
// 'a' -> 0
// 'b' -> 1
func lowercaseToInt(chr byte) int {
	return int(chr - 'a')
}
