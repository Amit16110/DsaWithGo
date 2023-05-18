package trie

import (
	"fmt"
	"testing"
)

func TestTri(t *testing.T) {
	trie := NewTrie()

	trie.Insert("abcdefgh")
	trie.Insert("abc")
	trie.Insert("ac")

	fmt.Println("trie remove result", trie.Remove("abc"))
	fmt.Println("trie search result", trie.Search("ac"))
	fmt.Println("trie search result", trie.Search("aa"))

}
