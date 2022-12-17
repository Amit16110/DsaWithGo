package linkedlist

import (
	"container/list"
	"fmt"
)

func LinkedListOperations(value int) {
	linkedList := list.New()
	linkedList.PushBack(value)

	fmt.Println("LinkedList==>")
	fmt.Println("LinkedList=from the winwork=>")
}
