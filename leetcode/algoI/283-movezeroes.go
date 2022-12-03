package algoI

func MoveZeroes(nums []int) {
	curP := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[curP] = nums[i]
			curP++
		}
	}
	for i := curP; i < len(nums); i++ {
		nums[i] = 0
	}
}
