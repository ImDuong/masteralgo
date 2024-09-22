package prefixsum

func Challenge560(nums []int, k int) int {
	return subarraySum(nums, k)
}

func subarraySum(nums []int, k int) int {
	sum := 0
	rs := 0
	// f is mapping prefix sum S to number of subarray having prefix sum as S
	f := make(map[int]int)
	f[0] = 1
	for i := range nums {
		sum += nums[i]

		// assume p[i], p[j] is prefix sum of subarray ending at index i, j, respectively
		// if subarray [j, i] has sum as k => p[i] - p[j - 1] = k
		// Hence, number of subarray having sum as k is equal to number of subarrays having sum as p[j - 1]
		// And instead of using second variable as j, we can shorten j by: f[p[j - 1]] = f[p[i] - k]
		// Hence, in this code, f[p[i] - k] is equivalent to f[sum - k]
		if _, ok := f[sum-k]; ok {
			rs += f[sum-k]
		}
		f[sum]++
	}
	return rs
}
