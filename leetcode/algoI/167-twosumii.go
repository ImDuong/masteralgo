package algoI

func TwoSumII(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return numbers
}

func TwoSumIIV2(numbers []int, target int) []int {
	headP, tailP := 0, len(numbers)-1
	for len(numbers) > 0 && headP < len(numbers) && tailP >= 0 {
		if numbers[headP]+numbers[tailP] > target {
			tailP--
		} else if numbers[headP]+numbers[tailP] < target {
			headP++
		} else {
			return []int{headP + 1, tailP + 1}
		}
	}
	return []int{}
}
