package prefixsum

func Challenge2845(nums []int, modulo int, k int) int64 {
	return countInterestingSubarrays(nums, modulo, k)
}

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	// nb of elements (that nums[i] % modulo == k) in the array nums in [0, i]
	// -> prefix sum: prefix[i]

	// prefix[l, r] = prefix[r] - prefix[l - 1]
	// and prefix[l, r] needs to % modulo == k
	// -> (prefix[r] - prefix[l - 1]) % modulo == k
	// -> (prefix[r] - k + modulo) % modulo == prefix[l - 1] % modulo
	// note: (prefix[r] - k + modulo) part is to avoid negative mods
	// Hence, we iterate r, and find how many occurences of l that satisfy the above equation
	// -> to find occurences quickly, we use map

	// nb occurences of (prefix[i] % modulo)
	freq := make(map[int]int)

	// since traveling once, -> use variable instead of array for prefix sum
	prefix := 0

	// prefix starts with 0, and 0 % modulo = 0
	// so we count 1 for the empty subarray
	freq[0] = 1

	var rs int64
	for i := range nums {
		if nums[i]%modulo == k {
			prefix++
		}
		rs += int64(freq[(prefix-k+modulo)%modulo])
		freq[prefix%modulo]++
	}
	return rs
}
