package boilerplates

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			newNode := Constructor()
			node.children[index] = &newNode
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.children[index] != nil {
			node = node.children[index]
		} else {
			return false
		}
	}

	return node.isEnd == true
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if node.children[index] != nil {
			node = node.children[index]
		} else {
			return false
		}
	}
	return true
}
