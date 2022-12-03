package algoI

// the recursive relation is really important: getMaximumRobbedMoney(P) = getMaximumRobbedMoney(P - 1) || nums[P] + getMaximumRobbedMoney(P - 2)

// recursive solution
func Rob(nums []int) int {
	maxPath := make([]int, len(nums))
	// init default value to faster trace back later
	for i := range maxPath {
		maxPath[i] = -1
	}
	return getMaximumRobbedMoney(nums, len(nums)-1, maxPath)
}

func getMaximumRobbedMoney(nums []int, P int, maxPath []int) int {
	if P < 0 {
		return 0
	}
	if maxPath[P] != -1 {
		return maxPath[P]
	}
	maxAmount := 0
	amount1 := getMaximumRobbedMoney(nums, P-1, maxPath)
	amount2 := getMaximumRobbedMoney(nums, P-2, maxPath) + nums[P]
	if maxAmount < amount1 {
		maxAmount = amount1
	}
	if maxAmount < amount2 {
		maxAmount = amount2
	}
	maxPath[P] = maxAmount
	return maxAmount
}

// iterative solution: heroes3001
func RobV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// we will spare a space for the -1 index in nums
	maxPath := make([]int, len(nums)+1)

	// index 0 in maxPath is equivalent to the index -1 in nums
	// which means index i in maxPath is eqivalent to the index i - 1 in nums
	maxPath[0] = 0
	maxPath[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		maxAmount := 0
		amount1 := maxPath[i]
		amount2 := maxPath[i-1] + nums[i]
		if maxAmount < amount1 {
			maxAmount = amount1
		}
		if maxAmount < amount2 {
			maxAmount = amount2
		}
		maxPath[i+1] = maxAmount
	}
	return maxPath[len(nums)]
}

// iterative optimized solution: heroes3001
func RobV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev1 := 0
	prev2 := 0
	for i := range nums {
		bak := prev1
		amount1 := prev1
		amount2 := prev2 + nums[i]
		if prev1 < amount1 {
			prev1 = amount1
		}
		if prev1 < amount2 {
			prev1 = amount2
		}
		prev2 = bak
	}
	return prev1
}
