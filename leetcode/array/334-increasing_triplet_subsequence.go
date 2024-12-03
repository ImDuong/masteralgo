package array

import "math"

func Challenge334(nums []int) bool {
	return increasingTriplet(nums)
}

func increasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt64, math.MaxInt64
	for i := range nums {
		if m1 >= nums[i] {
			m1 = nums[i]
		} else if m2 >= nums[i] {
			m2 = nums[i]
		} else {
			return true
		}
	}
	return false
}

// use template from boilerplates.lengthOfLIS
func increasingTripletUsingLIS(nums []int) bool {
	lis := []int{nums[0]}
	for i := range nums {
		lb := LowerBound(lis, nums[i])
		if lb >= len(lis) {
			lis = append(lis, nums[i])
		} else {
			// replace with new value to keep the subsequence strictly increasing
			lis[lb] = nums[i]
		}
		if len(lis) == 3 {
			return true
		}
	}
	return false
}

func LowerBound(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
