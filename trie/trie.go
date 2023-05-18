package trie

type TriNode struct {
	children map[rune]*TriNode
	isEnd    bool
}

func newTriNode() *TriNode {
	return &TriNode{
		children: make(map[rune]*TriNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TriNode
}

func NewTrie() *Trie {
	return &Trie{
		root: newTriNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, w := range word {
		if _, ok := node.children[w]; !ok {
			node.children[w] = newTriNode()
		}
		node = node.children[w]
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, w := range word {
		if _, ok := node.children[w]; !ok {
			return false
		}
		node = node.children[w]
	}
	return node.isEnd
}

func (t *Trie) StartWith(prefix string) bool {
	node := t.root

	for _, w := range prefix {
		if _, ok := node.children[w]; !ok {
			return false
		}
		node = node.children[w]
	}

	return true
}

func (t *Trie) Remove(word string) bool {
	node := t.root

	for _, w := range word {
		if _, ok := node.children[w]; !ok {
			return false
		}
		node = node.children[w]

	}
	node.isEnd = false
	return true
}
