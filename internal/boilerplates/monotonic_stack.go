package boilerplates

// use increasing monotonic stack
// stack store indices, while result stores the element
func PreviousLessElement(nums []int) []int {
	rs := make([]int, len(nums))
	stk := []int{}
	for i := range nums {
		for len(stk) > 0 && nums[i] < nums[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			rs[i] = 0
		} else {
			rs[i] = nums[stk[len(stk)-1]]
		}
		stk = append(stk, i)
	}
	return rs
}

// use increasing monotonic stack
func NextLessElement(nums []int) []int {
	rs := make([]int, len(nums))
	stk := []int{}
	for i := range nums {
		for len(stk) > 0 && nums[i] < nums[stk[len(stk)-1]] {
			rs[stk[len(stk)-1]] = nums[i]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return rs
}
