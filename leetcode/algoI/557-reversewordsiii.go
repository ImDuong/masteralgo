package algoI

import (
	"masteralgo/pkg/helpers"
	"strings"
)

func ReverseWordsIII(s string) string {
	reverseS := ""
	words := strings.Fields(s)
	for i := range words {
		word := strings.Split(words[i], "")
		helpers.Reverse(word)
		reverseS += strings.Join(word, "") + " "
	}
	return reverseS[:len(reverseS)-1]
}
