package algoI

import "unicode"

var letterPermutations = []string{}

func LetterCasePermutation(s string) []string {
	letterPermutations = []string{}
	permutRune := make([]rune, len(s))
	visited := make([]bool, len(s))
	tryLetterPermute(0, s, permutRune, visited)
	return letterPermutations
}

func tryLetterPermute(k int, orginalStr string, permutRune []rune, visited []bool) {
	// only 2 options: lowercase and uppercase
	for v := 0; v < 2; v++ {
		if v == 0 {
			permutRune[k] = unicode.ToLower(rune(orginalStr[k]))
		} else {
			if !unicode.IsLetter(permutRune[k]) {
				continue
			}
			permutRune[k] = unicode.ToUpper(rune(orginalStr[k]))
		}
		if k == len(orginalStr)-1 {
			letterPermutations = append(letterPermutations, string(permutRune))
		} else {
			tryLetterPermute(k+1, orginalStr, permutRune, visited)
		}
	}
}
