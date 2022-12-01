package algoI

func SingleNumber(nums []int) int {
	mapping := make(map[int]bool)
	for i := range nums {
		if _, ok := mapping[nums[i]]; ok {
			mapping[nums[i]] = false
		} else {
			mapping[nums[i]] = true
		}
	}
	for k := range mapping {
		if mapping[k] {
			return k
		}
	}
	return -1
}

// studied from one of the smartest Golang submission in leetcode
// apply XOR on all elements in nums -> the one exist only one will be remained in the end
func SingleNumberV2(nums []int) int {
	uniqueEle := 0
	for i := range nums {
		uniqueEle ^= nums[i]
	}
	return uniqueEle
}
