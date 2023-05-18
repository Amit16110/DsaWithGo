package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	// create a new tree
	tree := NewTree(6)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(8)
	tree.Insert(9)
	// fmt.Println("ans", ans)
	tree.Bsf(tree)

	fmt.Printf("height of tree: %d\n", MaxheightOfTree(tree))
}
