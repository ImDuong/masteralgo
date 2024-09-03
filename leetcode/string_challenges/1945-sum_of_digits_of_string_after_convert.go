package string_challenges

import "strconv"

func Challenge1945(s string, k int) int {
	return getLucky(s, k)
}

func getLucky(s string, k int) int {
	luckyStr := ""
	for i := 0; i < len(s); i++ {
		luckyStr += strconv.Itoa(int(s[i] - 96))
	}
	for i := 0; i < k; i++ {
		luckyStr = strconv.Itoa(gl(luckyStr))
	}
	rs, _ := strconv.Atoi(luckyStr)
	return rs
}

func gl(luckyNum string) int {
	rs := 0
	for i := 0; i < len(luckyNum); i++ {
		ele, _ := strconv.Atoi(string(luckyNum[i]))
		rs += ele
	}
	return rs
}
