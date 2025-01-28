package stack

// literally the same with challenge 678. Valid Parenthesis String
func Challenge2116(s string, locked string) bool {
	return canBeValid(s, locked)
}

// To determine s is valid, we need to base on 3 conditions:
// 1. length of s is even
// 2. number of unlocked + number of open brackets is >= number of close brackets
// 3. after neutralized with close brackets, we need to check
// 3.1. open brackets occur before unlocked ones
// 3.2. after open brackets neutralized with unlocked ones, number of open brackets must be zeroes
//
//	Note: we don't need to check with number of left over unlocked ones,
//	because if length is s is even, hence, number of left over unlocked ones also even -> still valid
func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}
	openBrackets := []int{}
	unlocked := []int{}

	for i := range s {
		if locked[i] == '0' {
			unlocked = append(unlocked, i)
		} else if s[i] == '(' {
			openBrackets = append(openBrackets, i)
		} else {
			if len(openBrackets) > 0 {
				openBrackets = openBrackets[:len(openBrackets)-1]
			} else if len(unlocked) > 0 {
				unlocked = unlocked[:len(unlocked)-1]
			} else {
				return false
			}
		}
	}

	for len(openBrackets) > 0 && len(unlocked) > 0 && openBrackets[len(openBrackets)-1] < unlocked[len(unlocked)-1] {
		openBrackets = openBrackets[:len(openBrackets)-1]
		unlocked = unlocked[:len(unlocked)-1]
	}

	return len(openBrackets) == 0
}
