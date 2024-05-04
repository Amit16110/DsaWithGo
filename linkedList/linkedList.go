package linkedlist

import (
	"fmt"
)

type Node struct {
	value any
	next  *Node
}

type ListNode struct {
	head, tail *Node
	len        int
}

func (l *ListNode) insertFront(value any) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.len++
	} else {
		current := l.head

		node.next = current
		l.head = node

		l.len++
	}
}

func (l *ListNode) insertBack(value any) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
		l.len++
	} else {
		curr := l.tail

		curr.next = node
		l.tail = node
		l.len++
	}
}

func (l *ListNode) reverse(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return current
}

func (l *ListNode) reverseLL() {

	if l.head == nil {
		return
	}

	var prev *Node
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev

}

func (l *ListNode) display() {
	current := l.head

	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}

	fmt.Println()
}
