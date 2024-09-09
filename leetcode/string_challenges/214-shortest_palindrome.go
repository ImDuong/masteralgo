package string_challenges

import (
	"fmt"
	"masteralgo/pkg/helpers"
)

func shortestPalindrome(s string) string {
	// the problem equivalent to find longest palindrome substring starts from 0
	// -> get the last value of lps obtained from "abcd@dcba"
	lps := helpers.GetLPS(s + "@" + reverse(s))
	palLen := lps[len(lps)-1]
	fmt.Println(palLen)
	leftOver := s[palLen:len(s)]
	return reverse(leftOver) + s
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
