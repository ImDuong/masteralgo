package twopointers

func Challenge75(nums []int) {
	sortColors(nums)
}

// dutch partioning problem
// we observer the array as 4 parts: red, white, unclassified, and blue
// at first, the whole array is the unclassified
// we goes from start to end of array to shrink the unclassified
// how?
// we create a white pointer used to iterate, a red pointer to swap red to head, a blue pointer to swap blue to tail
func sortColors(nums []int) {
	red, white, blue := 0, 0, len(nums)-1
	for white <= blue {

		// white pointer is at a red element
		if nums[white] == 0 {
			nums[white], nums[red] = nums[red], nums[white]
			red++

			// because white always moves before red, we always make sure that after this swap, color is correct for both white and red pointers -> we can safely increase white pointer
			white++
			continue
		}

		// white pointer is at the right place
		if nums[white] == 1 {
			white++
			continue
		}

		// white pointer is at a blue element
		nums[white], nums[blue] = nums[blue], nums[white]
		blue--
		// white still stands here for checking the new element
		continue
	}
}
