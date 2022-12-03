package algoI

// using fixed sliding window technique
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// setup fixed frequency table for s1 and for first window slide of s2
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	for i := range s1 {
		freq1[s1[i]-97]++
		freq2[s2[i]-97]++
	}

	// return true whenever the window slide of s2's freq table
	// is the same as the s1's freq table
	nbEquallyFreq := 0
	for i := range freq1 {
		if freq1[i] == freq2[i] {
			nbEquallyFreq++
		}
	}
	if nbEquallyFreq == 26 {
		return true
	}

	// for each sliding, update nbEquallyFreq & window slide's freq table of s2
	for i := 1; i <= len(s2)-len(s1); i++ {
		// adjust nbEquallyFreq whenever a character is pushed or popped
		popChar := s2[i-1] - 97
		pushChar := s2[i+len(s1)-1] - 97

		if freq2[popChar] == freq1[popChar] {
			nbEquallyFreq--
		}
		freq2[popChar]--
		if freq2[popChar] == freq1[popChar] {
			nbEquallyFreq++
		}

		if freq2[pushChar] == freq1[pushChar] {
			nbEquallyFreq--
		}
		freq2[pushChar]++
		if freq2[pushChar] == freq1[pushChar] {
			nbEquallyFreq++
		}

		if nbEquallyFreq == 26 {
			return true
		}
	}
	return nbEquallyFreq == 26
}
