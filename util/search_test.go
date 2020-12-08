package util

import (
	"fmt"
	"testing"
)

func TestBuildTrie_Trie(t *testing.T) {
	params := []string{"uui", "uii", "uim"}
	given := BuildTrie(params)
	root := &TrieNode{children: make(map[rune]*TrieNode)}
	root.children['u'] = &TrieNode{children: make(map[rune]*TrieNode)}
	node := root.children['u']
	node.children['u'] = &TrieNode{children: make(map[rune]*TrieNode)}
	node.children['i'] = &TrieNode{children: make(map[rune]*TrieNode)}
	fnode := node.children['u']
	snode := node.children['i']
	fnode.children['i'] = &TrieNode{isEnd: true, fuuid: "uui", children: make(map[rune]*TrieNode)}
	snode.children['i'] = &TrieNode{isEnd: true, fuuid: "uii", children: make(map[rune]*TrieNode)}
	snode.children['m'] = &TrieNode{isEnd: true, fuuid: "uim", children: make(map[rune]*TrieNode)}
	if !isTriesSame(given, root) {
		t.Error()
	}
}

func isTriesSame(given *TrieNode, expected *TrieNode) bool {
	isRootsSame := isTrieNodesSame(given, expected)
	fmt.Println(given)
	fmt.Println(expected)
	if !isRootsSame {
		return false
	}
	result := true
	for key := range expected.children {
		result = result && isTriesSame(given.children[key], expected.children[key])
	}
	return isRootsSame && result
}

func isTrieNodesSame(node1 *TrieNode, node2 *TrieNode) bool {
	isUIDSame := node1.fuuid == node2.fuuid
	isEndSame := node1.isEnd == node2.isEnd
	isChildrenLengthSame := len(node1.children) == len(node2.children)
	return isUIDSame && isEndSame && isChildrenLengthSame
}
