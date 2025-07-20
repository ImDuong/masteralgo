package trie

import (
	"fmt"
	"sort"
	"strings"
)

func Challenge1948(paths [][]string) [][]string {
	return deleteDuplicateFolder(paths)
}

// We break down this problem into 3 smaller steps
// 1. Build the trie
// 2. For each node, we store the serialize version of its subfolder structure, and store the frequency of each serial
// 3. Traverse from root and skip subfolders having serial version with many occurrences
// Now, to do the step 2, we will need:
// - post-order
// - wrap serial with `()` to serialize
// - sort the serial
// Note: serial idea is from `Leetcode 297. Serialize and Deserialize Binary Tree`
func deleteDuplicateFolder(paths [][]string) [][]string {
	// 1. build trie
	root := NewSerialTrie()
	for i := range paths {
		root.Insert(paths[i])
	}

	// 2. construct serial representation for each node
	freq := make(map[string]int)
	root.Serial(freq)

	// 3. traverse from root to add only unique folders
	rs := [][]string{}
	path := []string{}
	var addUniqueFolders func(*SerialTrie)
	addUniqueFolders = func(node *SerialTrie) {
		if freq[node.serial] > 1 {
			return
		}

		if len(path) > 0 {
			tmp := make([]string, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}

		// use backtrack to traverse all subpaths
		for folder, child := range node.children {
			path = append(path, folder)
			addUniqueFolders(child)

			// backtrack
			path = path[:len(path)-1]
		}
	}
	addUniqueFolders(root)
	return rs
}

type SerialTrie struct {
	children map[string]*SerialTrie
	serial   string // serialized ver of cur node
}

func NewSerialTrie() *SerialTrie {
	return &SerialTrie{
		children: make(map[string]*SerialTrie),
	}
}

func (this *SerialTrie) Insert(subFolders []string) {
	node := this
	for i := range subFolders {
		if node.children[subFolders[i]] == nil {
			newNode := NewSerialTrie()
			node.children[subFolders[i]] = newNode
		}
		node = node.children[subFolders[i]]
	}
}

// use post-order to serialize each node
// -> traverse all leaves before its parent
func (this *SerialTrie) Serial(freq map[string]int) {
	if len(this.children) == 0 {
		return
	}

	// use slice instead of string because we'll sort later
	curSerial := make([]string, 0, len(this.children))
	for folder, child := range this.children {
		child.Serial(freq)

		// combine serial of its children
		curSerial = append(curSerial, fmt.Sprintf("%s(%s)", folder, child.serial))
	}

	// make hash persistent
	sort.Strings(curSerial)
	this.serial = strings.Join(curSerial, "")
	freq[this.serial]++
}
