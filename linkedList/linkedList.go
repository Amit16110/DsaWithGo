package linkedlist

type Node struct {
	value any
	next  *Node
}

type ListNode struct {
	root *Node
	len  int
}

// func (l *ListNode) init() *ListNode {
// 	l.root.next = &l.root
// 	l.len = 0
// 	return l
// }
// func NewLinkedList() *ListNode {
// 	return new(ListNode).init()
// }

// func (l *ListNode) Front() *Node {
// 	if l.len == 0 {
// 		return nil
// 	}
// 	return l.root.next
// }

func (l *ListNode) insert(value any) *Node {
	node := &Node{value: value}
	if l.root == nil {
		l.root = node
	} else {
		current := l.root

		for current.next != nil {
			current = current.next
		}
		current = current.next
	}
	return nil // just exammple for error
}
