package string_challenges

func Challenge420(password string) int {
	return strongPasswordChecker(password)
}

func strongPasswordChecker(password string) int {
	curPwdLen := len(password)
	// aaa -> aaB1aa: 3 steps
	if curPwdLen <= 3 {
		return 6 - curPwdLen
	}

	nbSatisfiedConditions := 0
	isLower, isUpper, isDigit := false, false, false
	for i := 0; i < curPwdLen; i++ {
		c := password[i]
		if !isLower && c >= 97 && c <= 122 {
			nbSatisfiedConditions++
			isLower = true
		} else if !isUpper && c >= 65 && c <= 90 {
			nbSatisfiedConditions++
			isUpper = true
		} else if !isDigit && c >= 48 && c <= 57 {
			nbSatisfiedConditions++
			isDigit = true
		}
	}
	// 4 <= length < 6:
	// aaaaa -> aaB1aa: -> 2 steps
	// aaBaa -> aaB1aa: 1 step
	// aaaa -> aaB1aa: 2 steps
	// abB1 -> abB1aa: 2 steps
	if curPwdLen < 6 {
		return max(6-curPwdLen, 3-nbSatisfiedConditions)
	}

	// store length of consecutive groups
	consecutiveGrs := []int{}
	endGr := 0
	for i := 0; i <= curPwdLen; i++ {
		if i == curPwdLen || password[i] != password[endGr] {
			consecutiveGrs = append(consecutiveGrs, i-endGr)
			endGr = i
		}
	}
	nbConsecutiveGrs := len(consecutiveGrs)

	nbSteps := 0
	if curPwdLen <= 20 {
		for i := 0; i < nbConsecutiveGrs; i++ {
			if consecutiveGrs[i] >= 3 {
				// aaa -> replace 1
				// aaaa -> replace 1
				// aaaaa -> replace 1
				// -> from 3 to 5: replace 1
				// aaaaaa, aaaaaaa, aaaaaaaa -> from 6 to 8: replace 2
				// case x % 3 == 0 -> need to replace x / 3
				// case x % 3 == r -> need to replace (x - r) / 3
				nbSteps += (consecutiveGrs[i] - (consecutiveGrs[i] % 3)) / 3
			}
		}
		return max(nbSteps, 3-nbSatisfiedConditions)
	}

	// for case > 20, we need to convert to case <= 20
	// 1. remove characters in consecutive groups
	// until running out of consecutive chars or length <= 20
	// 2. if length still > 20, remove again
	// 3. repeat for case <= 20

	// in step 1 (removing steps), we need to prioritize for group that length % 3 == 0 then 1, and finally2
	// why?: aaa -> remove 1, and no longer need to replace
	// aaaa -> remove 1, and need 1 replacement
	// aaaaa -> remove 1, and need 1 replacement
	for j := 0; j < 3; j++ {
		isJustCleaned := true
		for curPwdLen > 20 && isJustCleaned {
			isJustCleaned = false
			for i := 0; i < nbConsecutiveGrs; i++ {
				if consecutiveGrs[i] >= 3 && consecutiveGrs[i]%3 == j {
					// don't need to rush removing all, just remove one by one until the length is enough
					nbRemovingSteps := min(j+1, curPwdLen-20)
					consecutiveGrs[i] -= nbRemovingSteps
					nbSteps += nbRemovingSteps
					curPwdLen -= nbRemovingSteps
					if curPwdLen <= 20 {
						break
					}
					isJustCleaned = true
				}
			}
		}
	}

	if curPwdLen > 20 {
		nbSteps += curPwdLen - 20
		curPwdLen = 20
	}

	additionalSteps := 0
	for i := 0; i < nbConsecutiveGrs; i++ {
		if consecutiveGrs[i] >= 3 {
			// repeat the case length <= 20
			additionalSteps += (consecutiveGrs[i] - (consecutiveGrs[i] % 3)) / 3
		}
	}
	nbSteps += max(additionalSteps, 3-nbSatisfiedConditions)
	return nbSteps
}
