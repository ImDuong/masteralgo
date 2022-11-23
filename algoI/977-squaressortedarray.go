package algoI

import "math"

func searchAndInsert(nums []int, target int) (int, bool) {
	if nums[0] > target {
		return 0, false
	}
	if nums[len(nums)-1] < target {
		return len(nums), false
	}
	lowB := 0
	upB := len(nums) - 1
	var jmp int
	for upB >= lowB {
		jmp = (upB + lowB) / 2
		if nums[jmp] < target {
			lowB = jmp + 1
		} else if nums[jmp] > target {
			upB = jmp - 1
		} else {
			return jmp, true
		}
	}
	return (upB+lowB)/2 + 1, false
}

// find sign division (zero if there is) then spread doube pointers from that anchor
func SortedSquares(nums []int) []int {
	ln := len(nums)
	numsRs := make([]int, ln)
	if ln == 2 {
		headE := nums[0] * nums[0]
		tailE := nums[1] * nums[1]
		if headE < tailE {
			numsRs[0] = headE
			numsRs[1] = tailE
		} else {
			numsRs[0] = tailE
			numsRs[1] = headE
		}
		return numsRs
	}
	// find sign division
	var rsP, tailP, headP int
	insIdx, isFound := searchAndInsert(nums, 0)
	if !isFound {
		headP = insIdx - 1
		tailP = insIdx
	} else {
		headP = insIdx
		tailP = insIdx + 1
	}
	for ; headP >= 0; headP-- {
		if tailP == ln {
			break
		}
		if math.Abs(float64(nums[tailP])) < math.Abs(float64(nums[headP])) {
			numsRs[rsP] = nums[tailP] * nums[tailP]
			rsP++
			tailP++
			headP++
		} else {
			numsRs[rsP] = nums[headP] * nums[headP]
			rsP++
		}
	}
	for ; tailP < ln; tailP++ {
		numsRs[rsP] = nums[tailP] * nums[tailP]
		rsP++
	}
	for ; headP >= 0; headP-- {
		numsRs[rsP] = nums[headP] * nums[headP]
		rsP++
	}
	return numsRs
}

func SortedSquaresV2(nums []int) []int {
	ln := len(nums)
	numRs := make([]int, ln)
	if ln == 2 {
		headE := nums[0] * nums[0]
		tailE := nums[1] * nums[1]
		if headE < tailE {
			numRs[0] = headE
			numRs[1] = tailE
		} else {
			numRs[0] = tailE
			numRs[1] = headE
		}
		return numRs
	}
	var headP int
	rsP := ln - 1
	tailP := ln - 1
	for headP <= tailP {
		headE := nums[headP] * nums[headP]
		tailE := nums[tailP] * nums[tailP]
		if tailE < headE {
			numRs[rsP] = headE
			headP++
		} else {
			numRs[rsP] = tailE
			tailP--
		}
		rsP--
	}
	return numRs
}
