package util

import (
	"fmt"
)

// TrieNode ...
type TrieNode struct {
	isEnd    bool
	children map[rune]TrieNode
	fuuid    string
}

/*
Use this trie to understand given uuid is unique or not
if the uuid is unique then get the actual uuid and use
the asked todo or command.
*/
func buildTrie(list []string) *TrieNode {
	root := TrieNode{children: make(map[rune]TrieNode)}
	node := root
	for _, str := range list {
		node = root
		for _, char := range str {
			if _, ok := node.children[char]; !ok {
				node.children[char] = TrieNode{children: make(map[rune]TrieNode)}
			}
			node = node.children[char]
		}
		node.isEnd = true
		node.fuuid = str
		fmt.Println(node)
	}
	fmt.Println(root)
	return &root
}

// FindUUIDFromPrefix ...
func FindUUIDFromPrefix(uid string) (bool, string) {
	root := *buildTrie([]string{})
	for _, char := range uid {
		if _, ok := root.children[char]; ok {
			root = root.children[char]
		} else {
			return false, ""
		}
	}

	// Go if there is only one children exists.
	for len(root.children) == 1 {
		for key := range root.children {
			root = root.children[key]
		}
	}
	// Return the uuid if it is end and there is no children.
	if root.isEnd && len(root.children) == 0 {
		return true, root.fuuid
	}
	return false, ""
}
