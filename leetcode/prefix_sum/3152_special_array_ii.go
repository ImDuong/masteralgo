package prefixsum

func Challenge3152(nums []int, queries [][]int) []bool {
	return isArraySpecial(nums, queries)
}

func isArraySpecial(nums []int, queries [][]int) []bool {
	// prefix sum: record how many violate indices before
	ps := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if isDifParity(nums[i-1], nums[i]) {
			ps[i] = ps[i-1]
		} else {
			ps[i] = ps[i-1] + 1
		}
	}
	rs := make([]bool, len(queries))
	for i := range queries {
		if ps[queries[i][0]] == ps[queries[i][1]] {
			rs[i] = true
		}
	}
	return rs
}

func isDifParity(a, b int) bool {
	return (a%2 == 0 && b%2 != 0) || (a%2 != 0 && b%2 == 0)
}
