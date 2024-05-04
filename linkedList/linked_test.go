package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := ListNode{}

	ll.insertFront(5)
	ll.insertFront(15)
	ll.insertFront(25)
	ll.insertFront(35)
	ll.insertBack(40)
	ll.insertBack(50)
	ll.display()

	// ll.reverseLL()
	// ll.display()
}
