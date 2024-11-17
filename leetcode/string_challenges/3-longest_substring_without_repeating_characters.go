package string_challenges

func Challenge3(s string) int {
	return lengthOfLongestSubstring(s)
}

// use the boilerplates.findSubString template
func lengthOfLongestSubstring(s string) int {
	d := make([]int, 128)

	// left, right for the boundary of two pointers
	left, right := 0, 0
	rs := 0

	// counter to determine substring is valid/invalid
	cnt := 0
	for ; right < len(s); right++ {

		// check condition: e.g., > 0
		if d[s[right]] > 0 {
			cnt++
		}
		d[s[right]]++

		// shrunk the window until we find the valid substring
		for cnt > 0 {
			if d[s[left]] > 1 {
				cnt--
			}

			d[s[left]]--

			// increase left to shrunk the window
			left++
		}
		rs = max(rs, right-left+1)
	}
	return rs
}
