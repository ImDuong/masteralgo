package dynamicprogramming

// LongestPalindrome finds longest palindrome bruteforcely
func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	gMaxStr := s[0:1]
	gMaxLen := 1
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			isP := true
			for k := i; k <= (i+j)/2; k++ {
				// transpose to array where i starts from 0 -> get the formula
				if s[k] != s[j-k+i] {
					isP = false
					break
				}
			}
			if isP && j-i+1 > gMaxLen {
				gMaxLen = j - i + 1
				gMaxStr = s[i : j+1]
			}
		}
	}
	return gMaxStr
}

// LongestPalindromeWithManachers finds longest palindrome using Manachers algorithm
// 1. walk through all characters in string and try to expand from each character as a center
//    -> store the length of longest palindrome for each character
// 2. take shortcut: when we already find a palindrome having center as i, right boundary as r
//    -> apply mirror property of palindrome: skip the right part of found palindrome
//       and continue to expand outside the skipped part
// Note: algo works with string having odd length -> add # interleaved
// Ref: basic idea and why need to add # interleaved: https://cp-algorithms.com/string/manacher.html
// Ref2: implementation: https://www.hackerearth.com/practice/algorithms/string-algorithm/manachars-algorithm/tutorial/
func LongestPalindromeWithManachers(s string) string {
	if len(s) <= 1 {
		return s
	}
	// add `?` so that never worry about out of bound
	q := "?"
	for i := range s {
		q += "#" + string(s[i])
	}
	// why not use `?`, but use another special character: in the case of entire s is a palindrome,
	// needs to make the head and tail of string differently -> escape the algo
	q += "#^"

	// every element in P (aka P[i]) represents the length of palindrome string having center as i
	P := make([]int, len(q))
	gLongestPCenter := 0
	gLongestPLen := 0
	n := len(q)
	// last found center and right limit
	// -> used for applying mirror property of palindrome to speedup the process
	c := 0
	r := 0
	for i := 2; i < n-2; i++ {

		// Step 1: find the initial value of P[i]
		// when i is in the range of last found palindrome string,
		// apply the mirror property for assigning the initial value for P[i] as P[iMirror]
		if i < r {
			// iMirror is the symmetric element of i along the center c.
			// Hence, i - c = c - iMirror
			iMirror := 2*c - i

			// however, i is in range (c, r)
			// -> palindrome string with i as center cannot extend beyond the boundary r. Why?
			//        - because when outside of r, we're not sure the palindrome string with i as center
			//          acting just like the palindrome string with iMirror as center
			// -> P[i]'s initial value cannot exceed r - i
			P[i] = min(P[iMirror], r-i)
		}

		// Step 2: expand around center i. The P[i] now accelerates the process after applying the mirror property
		// - Specifically, instead of checking q[i-1] == q[i+1], we can go to checking q[i-1-P[i]] == q[i+1+P[i]].
		// - Note: because we add ? and ^ around the boundary of string,
		// we will never afraid of (i-1-P[i]) or (i+1+P[i]) out of bound
		for q[i-1-P[i]] == q[i+1+P[i]] {
			// double checking:
			// case original string is an odd string: cac -> #c#a#c# -> P[i] = 3 (length of #c# with a as center)
			// case original string is an even string: caac -> #c#a#a#c# -> P[i] = 4 (length of #c#a with # as center)
			P[i]++
		}

		// Step 3: if the palindrome string centered at i expands past r, update the last center and right limit
		if i+P[i] > r {
			r = i + P[i]
			c = i
		}

		// Step 4: update the global max value
		if gLongestPLen < P[i] {
			gLongestPLen = P[i]
			gLongestPCenter = i
		}
	}

	startIdx := (gLongestPCenter - P[gLongestPCenter]) / 2
	return s[startIdx : startIdx+P[gLongestPCenter]]
}
