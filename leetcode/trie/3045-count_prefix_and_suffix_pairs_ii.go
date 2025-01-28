package trie

func Challenge3045(words []string) int64 {
	return countPrefixSuffixPairs(words)
}

// use dual trie (prefix and suffix) to track
func countPrefixSuffixPairs(words []string) int64 {
	var rs int64
	dt := DualTrie{
		preChild: &Trie{},
		sufChild: &Trie{},
	}
	for i := range words {
		rs += dt.Insert(words[i])
	}
	return rs
}

type DualTrie struct {
	// tree for prefix & suffix
	preChild, sufChild *Trie
}

type Trie struct {
	children map[byte]*Trie

	// track number of words having prefix/suffix at this child
	freq int

	isEnd bool

	// use this to check prefix and suffix trie is pointing to the same word
	mapped *Trie
}

// insert word into trie, and return number of new valid pairs after adding new word
func (this *DualTrie) Insert(word string) int64 {
	var nbPairs int64
	preNode, sufNode := this.preChild, this.sufChild
	for i := range word {
		left, right := word[i], word[len(word)-i-1]
		if preNode.children[left] != nil && sufNode.children[right] != nil && preNode.children[left].isEnd && sufNode.children[right].isEnd && preNode.children[left].mapped == sufNode.children[right] {
			nbPairs += int64(preNode.children[left].freq)
		}
		if preNode.children == nil {
			preNode.children = make(map[byte]*Trie)
		}
		if preNode.children[left] == nil {
			newNode := Trie{}
			preNode.children[left] = &newNode
		}
		if sufNode.children == nil {
			sufNode.children = make(map[byte]*Trie)
		}
		if sufNode.children[right] == nil {
			newNode := Trie{}
			sufNode.children[right] = &newNode
		}
		preNode = preNode.children[left]
		sufNode = sufNode.children[right]
	}
	preNode.isEnd = true
	sufNode.isEnd = true
	preNode.mapped = sufNode
	preNode.freq++
	return nbPairs
}

// awesome trick to hash the prefix and suffix into a single key
// credits: https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/solutions/4744547/java-c-python-trie
func countPrefixSuffixPairsWithSingleTrie(words []string) int64 {
	var rs int64
	root := trie3045{
		children: make(map[int]*trie3045),
	}
	for i := range words {
		cur := &root
		for j := range words[i] {
			key := int(words[i][j]-'a')*26 + int(words[i][len(words[i])-j-1]-'a')
			if cur.children[key] == nil {
				cur.children[key] = &trie3045{
					children: make(map[int]*trie3045),
				}
			}
			cur = cur.children[key]
			rs += int64(cur.freq)
		}
		cur.freq++
	}
	return rs
}

type trie3045 struct {
	children map[int]*trie3045
	freq     int
}
