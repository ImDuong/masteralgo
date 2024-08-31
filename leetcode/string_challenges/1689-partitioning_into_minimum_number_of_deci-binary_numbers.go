package string_challenges

// just return the biggest digit in the string
func MinPartitions(n string) int {
	rs := 0
	for _, c := range n {
		rs = max(rs, int(c-'0'))
	}
	return rs
}
