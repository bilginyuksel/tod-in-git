package util

import (
	"fmt"
	"testing"
)

func TestBuildTrie_Trie(t *testing.T) {
	params := []string{"uui", "uii", "uim"}
	given := buildTrie(params)
	root := TrieNode{children: make(map[rune]TrieNode)}
	root.children['u'] = TrieNode{children: make(map[rune]TrieNode)}
	node := root.children['u']
	node.children['u'] = TrieNode{children: make(map[rune]TrieNode)}
	node.children['i'] = TrieNode{children: make(map[rune]TrieNode)}
	fnode := node.children['u']
	snode := node.children['i']
	fnode.children['i'] = TrieNode{isEnd: true, fuuid: "uui", children: make(map[rune]TrieNode)}
	snode.children['i'] = TrieNode{isEnd: true, fuuid: "uii", children: make(map[rune]TrieNode)}
	snode.children['i'] = TrieNode{isEnd: true, fuuid: "uim", children: make(map[rune]TrieNode)}

	if !isTriesSame(given, &root) {
		t.Error()
	}
}

func isTriesSame(given *TrieNode, expected *TrieNode) bool {
	isRootsSame := isTrieNodesSame(given, expected)
	node1 := *given
	node2 := *expected
	fmt.Println(node1)
	fmt.Println(expected)
	for key := range node2.children {
		node1 = node1.children[key]
		node2 = node2.children[key]
		// fmt.Println(node1)
		// if !isTrieNodesSame(&node1, &node2) {
		// 	return false
		// }
	}
	return true && isRootsSame
}

func isTrieNodesSame(node1 *TrieNode, node2 *TrieNode) bool {
	isUIDSame := node1.fuuid == node2.fuuid
	isEndSame := node1.isEnd == node2.isEnd
	isChildrenLengthSame := len(node1.children) == len(node2.children)
	return isUIDSame && isEndSame && isChildrenLengthSame
}
