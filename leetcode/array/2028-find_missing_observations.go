package array

func Challenge2028(rolls []int, mean int, n int) []int {
	return missingRolls(rolls, mean, n)
}

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sumSeenObs := 0
	for i := 0; i < m; i++ {
		sumSeenObs += rolls[i]
	}
	sumMissingObs := mean*(m+n) - sumSeenObs
	if sumMissingObs > n*6 || sumMissingObs < n*1 {
		return []int{}
	}

	avgObValue := sumMissingObs / n
	seenObs := make([]int, n)
	for i := 0; i < n; i++ {
		seenObs[i] = avgObValue
	}
	if avgObValue*n != sumMissingObs {
		remainingValue := sumMissingObs - avgObValue*n
		for i := 0; i < n; i++ {
			if remainingValue == 0 {
				break
			}
			addedValue := min(6-seenObs[i], remainingValue)
			if addedValue == 0 {
				continue
			}
			seenObs[i] += addedValue
			remainingValue -= addedValue
		}
		if remainingValue > 0 {
			return []int{}
		}
	}
	return seenObs
}
