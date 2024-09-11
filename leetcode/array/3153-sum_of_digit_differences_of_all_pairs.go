package array

func Challenge3153(nums []int) int64 {
	return sumDigitDifferences(nums)
}

// the idea is the same with Challenge 477
// we go through each digit in each layer of all numbers
func sumDigitDifferences(nums []int) int64 {
	var rs int64 = 0
	nbEles := len(nums)
	// nums[i] <= 10^9 -> max number of digits are: 10
	for j := 0; j < 10; j++ {
		digitToFreq := make([]int, 10)
		for i := 0; i < nbEles; i++ {
			lastDigit := nums[i] % 10
			digitToFreq[lastDigit]++

			// remove last digit
			nums[i] /= 10
		}

		tmp := 0
		// 2 solutions to find the count for all pairs
		// solution 1: loop through all pairs: O(10^2)
		// solution 2: the same with challenge 477: O(10)
		// 				- remember to divide the final ans with 2, since we count twice for each pair
		for i := 0; i < 10; i++ {
			tmp += digitToFreq[i] * (nbEles - digitToFreq[i])
		}
		rs += int64(tmp)
		if nums[0] == 0 {
			break
		}
	}
	return rs / 2
}
