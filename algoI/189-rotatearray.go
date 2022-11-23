package algoI

import (
	"masteralgo/pkg/helpers"
)

// thank for this exercise, I learn that O(1) space complexity means
// the space required by the algorithm does not grow with the input or size of the data

func NaiveRotateArray(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	for i := range nums {
		if i+len(nums)-k < len(nums) {
			nums[i] = copiedNums[i+len(nums)-k]
		} else {
			nums[i] = copiedNums[i-k]
		}
	}
}

// Step k is the division:
// the last k elements of array will be pushed to the head of array while they maintain the original order
// divide the array to 2 parts, where the second part contains the last k elements
// nums -> [A, B]
// A: A1, A2, ..., Ai
// B: B1, B2, ..., Bi
// our target is: nums -> [B1, B2, ..., Bi, A1, A2, ..., Ai]
// to make B to the head, we simply reverse the array. After the reverse, nums become: [B, A]
// however, the exactly order is: Bi, ..., B2, B1, Ai, ...., A2, A1
// Therefore, to achieve the goal,
// the simple idea is reverse A, reverse B and then reverse both of them
func RotateArray(nums []int, k int) {
	k = k % len(nums)
	helpers.Reverse(nums[0 : len(nums)-k])
	helpers.Reverse(nums[len(nums)-k:])
	helpers.Reverse(nums)
}
