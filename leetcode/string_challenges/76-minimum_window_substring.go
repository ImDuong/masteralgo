package string_challenges

// use the boilerplates.findSubString template
func minWindow(s string, t string) string {
	d := make([]int, 128)
	for i := range t {
		d[t[i]]++
	}

	// left, right for the boundary of two pointers
	left, right := 0, 0

	const MaxInt = int(^uint(0) >> 1)
	minLen := MaxInt
	head := 0

	// if cnt reaches 0, it means current sliding window contains t -> we have a valid substring
	cnt := len(t)
	for ; right < len(s); right++ {
		if d[s[right]] > 0 {
			cnt--
		}
		d[s[right]]--

		for cnt == 0 {
			// update rs if we are finding minimum
			if minLen > right-left+1 {
				head = left
				minLen = right - left + 1
			}

			if d[s[left]] == 0 {
				cnt++
			}

			d[s[left]]++
			left++
		}
	}
	if minLen == MaxInt {
		return ""
	}
	return s[head : head+minLen]
}
