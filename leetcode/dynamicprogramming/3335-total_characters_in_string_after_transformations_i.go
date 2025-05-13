package dynamicprogramming

func Challenge3335(s string, t int) int {
	return lengthAfterTransformations(s, t)
}

const mod = 1_000_000_007

func lengthAfterTransformations(s string, t int) int {
	// f(i, c) is the nb of occur of char c in str after i transformations, where c in [0, 25]
	// if c = 0, 'a' can be only converted from 'z'. Hence
	//      f(i, 0) = f(i - 1, 25)
	// if c = 1, 'b' can be converted from 'a' or 'z'. Hence
	//      f(i, 1) = f(i - 1, 25) + f(i - 1, 0)
	// if c >= 2, f(i, c) = f(i - 1, c - 1)

	f := make([][]int, t+1)
	for i := range f {
		f[i] = make([]int, 26)
	}

	// init f(0, c)
	for i := range s {
		f[0][int(s[i]-'a')]++
	}

	for i := 1; i <= t; i++ {
		f[i][0] = f[i-1][25]
		f[i][1] = (f[i-1][25] + f[i-1][0]) % mod

		for c := 2; c < 26; c++ {
			f[i][c] = f[i-1][c-1]
		}
	}
	rs := 0
	for c := range f[0] {
		rs = (rs + f[t][c]) % mod
	}
	return rs
}

// optimize: because f(i, c) only require f(i-1, ) hence, we just need two 1-D arrays (both having fixed size) for previous and current frequencies
func lengthAfterTransformationsOptimizeSpace(s string, t int) int {
	prev := make([]int, 26)

	// init f(0, c)
	for i := range s {
		prev[int(s[i]-'a')]++
	}

	for i := 1; i <= t; i++ {
		cur := make([]int, 26)
		cur[0] = prev[25]
		cur[1] = (prev[25] + prev[0]) % mod

		for c := 2; c < 26; c++ {
			cur[c] = prev[c-1]
		}
		prev = cur
	}
	rs := 0
	for c := range prev {
		rs = (rs + prev[c]) % mod
	}
	return rs
}

// optimize 2: we backup freqs of f(i, 1) before update -> we only need one 1-D fixed size array
func lengthAfterTransformationsOptimize2(s string, t int) int {
	freq := make([]int, 26)

	// init f(0, c)
	for i := range s {
		freq[int(s[i]-'a')]++
	}

	for i := 1; i <= t; i++ {
		bFreq := freq[1]
		freq[1] = (freq[25] + freq[0]) % mod
		freq[0] = freq[25]

		for c := 25; c > 2; c-- {
			freq[c] = freq[c-1]
		}
		freq[2] = bFreq
	}
	rs := 0
	for c := range freq {
		rs = (rs + freq[c]) % mod
	}
	return rs
}
