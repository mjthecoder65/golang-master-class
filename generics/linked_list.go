package generics

import "fmt"

// Implementing a linked list using a normal in Golang.
type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

type LinkedList[T any] struct {
	Head *ListNode[T]
}

func (list *LinkedList[T]) InsertAtHead(element T) {
	newNode := ListNode[T]{
		Val:  element,
		Next: list.Head,
	}
	list.Head = &newNode
}

func (list *LinkedList[T]) RemoveFromHead() (T, bool) {
	if list.Head == nil {
		var zeroValue T
		return zeroValue, false
	}
	removedValue := list.Head.Val
	list.Head = list.Head.Next // Mark and sweep algorithm for garbabe collection will collect the unreferenced nodes.
	return removedValue, true
}

func (list *LinkedList[T]) Print() {
	currentNode := list.Head

	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}
