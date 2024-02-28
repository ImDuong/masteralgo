package algoII

func SearchRotatedSortedArray(nums []int, target int) int {
	// NAIVE solution: O(n)
	// TODO: add another solution to gain advantage from sorted array
	for i := range nums {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
