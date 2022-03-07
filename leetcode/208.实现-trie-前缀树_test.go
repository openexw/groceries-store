package leetcode

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := Constructor()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("hello")
	//trie.Print()
}

func TestTrie_Search(t *testing.T) {
	trie := Constructor()
	trie.Insert("app")
	trie.Insert("apple")
	//trie.Insert("hello")
	t.Log(trie.Search("app"))
}
