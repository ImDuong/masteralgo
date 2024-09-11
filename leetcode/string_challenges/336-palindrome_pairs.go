package string_challenges

import "fmt"

func Challenge336(words []string) [][]int {
	return palindromePairs(words)
}

func palindromePairs(words []string) [][]int {
	wordDict := make(map[string]int)
	checked := make(map[int][]int)
	maxLen := 0
	emptyWordIdx := -1
	for i := range words {
		wordDict[words[i]] = i
		wordLen := len(words[i])
		if maxLen < wordLen {
			maxLen = wordLen
		}
		if wordLen == 0 {
			emptyWordIdx = i
		}
	}
	rs := make([][]int, 0)
	fmt.Println("word dict", wordDict)
	for i, word := range words {
		if i == emptyWordIdx {
			continue
		}
		// convert to the problem find the shortest prefix to make a palindome: challenge 214
		{
			lps := getLPS(word + "@" + reverse(word))
			palLen := lps[len(lps)-1]
			if palLen == len(word) && emptyWordIdx != -1 {
				rs = append(rs, []int{emptyWordIdx, i})
				rs = append(rs, []int{i, emptyWordIdx})
			}
			leftOver := word[palLen:]
			prefix := reverse(leftOver)

			if len(prefix) == 0 {
				// for case itself is palindrome, try to duplicate itself and find
				tmpWord := word
				for k := 1; k <= maxLen/len(word); k++ {
					// tmpWord = aba
					// -> find abaaba, abaabaaba, ... (duplicating)
					tmpWord += string(word[0])
					j, ok := wordDict[tmpWord]
					if ok {
						if !isExisted(i, j, checked) {
							rs = append(rs, []int{i, j})
							checked[i] = append(checked[i], j)
						}
						if !isExisted(j, i, checked) {
							rs = append(rs, []int{j, i})
							checked[j] = append(checked[j], i)
						}
					}
				}

				if len(word) > 1 {
					// or find second half of itself and duplicate of half: ba, baba, ...
					tmpWord = ""
					for k := 1; k <= maxLen/(len(word)/2); k++ {
						// tmpWord = aba
						// -> find abaaba, abaabaaba, ... (duplicating)
						tmpWord += word[len(word)/2:]
						if tmpWord == word {
							continue
						}
						j, ok := wordDict[tmpWord]
						if ok {
							if !isExisted(i, j, checked) {
								rs = append(rs, []int{i, j})
								checked[i] = append(checked[i], j)
							}
						}
					}

					// for the case: aaaa -> it cannot match with other strings
					continue
				}
			}

			// word: abcd
			// leftOver: bcd -> prefix: dcb -> searchWord = dcba
			searchWord := prefix + word[:palLen]
			isSearchFound := findAndAdd(i, word, searchWord, wordDict, checked, &rs, false)

			searchWord = prefix
			findAndAdd(i, word, searchWord, wordDict, checked, &rs, false)

			for k := 1; k <= maxLen-len(prefix); k++ {
				if len(word) == 0 || len(prefix) == 0 {
					break
				}
				searchWord += string(word[0])

				// avoid duplicate when the first search word
				// (i.e., prefix + word[:palLen]) has been added already
				if isSearchFound && palLen == 1 && k == 1 {
					continue
				}
				findAndAdd(i, word, searchWord, wordDict, checked, &rs, false)
			}
		}

		// now do the same thing but to find the shortest suffix to make a palindrome
		{
			word = reverse(word)
			lps := getLPS(word + "@" + reverse(word))
			palLen := lps[len(lps)-1]
			leftOver := word[palLen:]
			prefix := reverse(leftOver)

			searchWord := prefix + word[:palLen]
			isSearchFound := findAndAdd(i, word, reverse(searchWord), wordDict, checked, &rs, true)

			searchWord = prefix
			findAndAdd(i, word, reverse(searchWord), wordDict, checked, &rs, true)

			for k := 1; k <= maxLen-len(prefix); k++ {
				if len(word) == 0 || len(prefix) == 0 {
					break
				}
				searchWord += string(word[0])

				// avoid duplicate when the first search word
				// (i.e., prefix + word[:palLen]) has been added already
				if isSearchFound && palLen == 1 && k == 1 {
					continue
				}
				findAndAdd(i, word, reverse(searchWord), wordDict, checked, &rs, true)
			}
		}
	}
	return rs
}

func isExisted(i, j int, checked map[int][]int) bool {
	if len(checked[i]) == 0 {
		return false
	}
	for k := range checked[i] {
		if checked[i][k] == j {
			return true
		}
	}
	return false
}

func findAndAdd(i int, word, searchWord string, wordDict map[string]int, checked map[int][]int, rs *[][]int, isIFirst bool) bool {
	if searchWord == word || len(searchWord) == 0 {
		return false
	}
	j, ok := wordDict[searchWord]
	if ok {
		if isIFirst {
			if !isExisted(i, j, checked) {
				*rs = append(*rs, []int{i, j})
				checked[i] = append(checked[i], j)
			}
		} else {
			if !isExisted(j, i, checked) {
				*rs = append(*rs, []int{j, i})
				checked[j] = append(checked[j], i)
			}
		}
	}
	return ok
}

func getLPS(pat string) []int {
	if len(pat) == 0 {
		return []int{}
	}
	patLen := len(pat)
	lps := make([]int, patLen)

	// always start at 0, because there is no proper prefix
	// for string having length = 1
	lps[0] = 0

	maxMatch := 0
	for i := 1; i < patLen; i++ {
		if pat[i] == pat[maxMatch] {
			maxMatch++
			lps[i] = maxMatch
			continue
		} else if maxMatch > 0 {
			maxMatch = lps[maxMatch-1]
			i--
		}
	}

	return lps
}
