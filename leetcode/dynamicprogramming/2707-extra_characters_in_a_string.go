package dynamicprogramming

func Challenge2707(s string, dictionary []string) int {
	return minExtraChar(s, dictionary)
}

// call f[i + 1] as the minimum number of extra characters left over if you break up s[0: i+1] optimally
// e.g., abcd -> s[0:4]
// f[0] = 0 because no extra characters
//
// f[4] should be minimum of these 2 cases
// case 1: 'd' is the extra character, hence f[4] = f[3] + 1
// case 2: the dictionary may have word (called as wj. E.g., 'bcd') appeared as suffix of 'abcd'
//
//	        -> the string needs to find is only 'a' (which is s[0]), while no extra characters in 'bcd'
//			-> this means we just need to find f[1]
//	        -> f[4] = f[1] + 0  (zero means no extra characters in s[1:4])
//
// Hence, f[4] = min(f[3] + 1, f[1])
// -> f[i + 1] = min(f[i] + 1, f[i + 1 - len(wj)])
//
// Final questions!
// Why the minExtraChar of s (having length as n) lies on f[n] but not f[n - 1]?
// This question can be rephrased as "Why not f[i], but f[i + 1] is for s[i]?"
// To answer this, let's use Reductio ad absurdum
// Assumption: if we set up f[i] is for s[i], this means when we are at s[0]
//   - if len(wj) = 1 -> need to find f[0 - 1]
//   - as we can see, f[-1] is complicated to implement
//   - hence, we should setup f[i + 1] is for s[i]
//   - that's why our final result for s lies on f[n]
func minExtraChar(s string, dictionary []string) int {
	w := make(map[string]bool)
	for i := range dictionary {
		w[dictionary[i]] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		f[i+1] = f[i] + 1
		for wj := range w {
			nwj := len(wj)
			if i+1-nwj >= 0 && s[i+1-nwj:i+1] == wj {
				f[i+1] = min(f[i+1], f[i+1-nwj])
			}
		}
	}
	return f[n]
}
