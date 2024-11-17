package boilerplates

// Template to solve most 'substring' problems
// Ref: https://leetcode.com/problems/minimum-window-substring/solutions/26808/Here-is-a-10-line-template-that-can-solve-most-'substring'-problems
// usually, if the problem is finding minimum -> we need to find the valid solution first, and then find the optimal one
// otherwise, when the problem is finding maximum -> we need to mark the invalidity of the substring -> then try to find the valid substring
func findSubString(s string) int {
	d := make([]int, 128)

	// left, right for the boundary of two pointers
	left, right := 0, 0
	rs := 0

	// counter to determine substring is valid/invalid
	cnt := 0
	for ; right < len(s); right++ {

		// check condition: e.g., > 0
		if d[s[right]] > 0 {
			// modify counter
		}
		d[s[right]]++

		// shrunk the window until we find the valid substring
		for cnt > 0 {
			// update rs if we are finding minimum

			// check condition: e.g., > 1
			if d[s[left]] > 1 {
				// modify counter to make the window valid/invalid
			}

			d[s[left]]--
			// increase left to shrunk the window
			left++
		}

		// update rs if we are finding maximum
	}
	return rs
}
