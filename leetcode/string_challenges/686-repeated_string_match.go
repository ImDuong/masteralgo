package string_challenges

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	dist := (len(a) + len(b) - 1) / len(a)
	fmt.Println(len(a), len(b), dist)
	for i := dist; i < dist+2; i++ {
		ss := strings.Repeat(a, i)
		if strings.Contains(ss, b) {
			return i
		}
	}
	return -1
}
