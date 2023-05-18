package tree

import (
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

func NewTree(root int) *Node {
	return &Node{value: root}
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return NewTree(val)
	}

	if val < n.value {
		n.left = n.left.Insert(val)
	} else {
		n.right = n.right.Insert(val)
	}
	return n
}

// Print InOrder traverses the binary tree in order and prints its values
func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}
	n.left.PrintInOrder()
	fmt.Println("Value: ", n.value)
	n.right.PrintInOrder()

}

func (n *Node) Bsf(root *Node) {
	queue := []*Node{}

	queue = append(queue, root)
	queue = append(queue, nil)

	for len(queue) != 0 {
		temp := queue[0]

		queue = queue[1:]

		if temp == nil { // level is completely travers
			if len(queue) != 0 {
				queue = append(queue, nil)

			}
		} else {

			fmt.Println(temp.value)

			if temp.left != nil {
				queue = append(queue, temp.left)
			}

			if temp.right != nil {
				queue = append(queue, temp.right)
			}
		}
	}

}
