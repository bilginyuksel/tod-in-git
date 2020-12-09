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

func TestBuildTrie_Trie2(t *testing.T) {
	params := []string{"ua", "be", "cu"}
	given := BuildTrie(params)
	root := &TrieNode{children: make(map[rune]*TrieNode)}
	root.children['u'] = &TrieNode{children: make(map[rune]*TrieNode)}
	root.children['b'] = &TrieNode{children: make(map[rune]*TrieNode)}
	root.children['c'] = &TrieNode{children: make(map[rune]*TrieNode)}
	fnode := root.children['u']
	snode := root.children['b']
	tnode := root.children['c']
	fnode.children['a'] = &TrieNode{isEnd: true, fuuid: "ua", children: make(map[rune]*TrieNode)}
	snode.children['e'] = &TrieNode{isEnd: true, fuuid: "be", children: make(map[rune]*TrieNode)}
	tnode.children['u'] = &TrieNode{isEnd: true, fuuid: "cu", children: make(map[rune]*TrieNode)}
	if !isTriesSame(given, root) {
		t.Error()
	}
}

func TestFindUUIDInsideTrie_FindUUID(t *testing.T) {
	params := []string{"uui", "uii", "uim"}
	trie := BuildTrie(params)
	ok, given := trie.FindUUIDFromPrefix("uu")
	if !ok {
		t.Error()
	}
	if given != "uui" {
		t.Error()
	}
}

func TestFindUUIDInsideTrie_UUIDNotFound(t *testing.T) {
	params := []string{"destroy", "destrol", "destool", "desotrol"}
	trie := BuildTrie(params)
	ok, _ := trie.FindUUIDFromPrefix("destro")
	if ok {
		t.Error()
	}
	ok, _ = trie.FindUUIDFromPrefix("dest")
	if ok {
		t.Error()
	}
}

func TestFindUUIDInsideTrie_UUIDFound(t *testing.T) {
	params := []string{"destroy", "destrol", "destool", "desotrol"}
	trie := BuildTrie(params)
	ok, given := trie.FindUUIDFromPrefix("deso")
	if !ok || given != "desotrol" {
		t.Error()
	}
}

func TestFindUUIDInsideTrie_NotRelatedInput(t *testing.T) {
	params := []string{"destroy", "destrol", "destool", "desotrol"}
	trie := BuildTrie(params)
	ok, _ := trie.FindUUIDFromPrefix("seto")
	if ok {
		t.Error()
	}
	ok, _ = trie.FindUUIDFromPrefix("destroyol")
	if ok {
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
