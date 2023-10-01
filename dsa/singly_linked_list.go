package dsa

import "errors"

// SinglyLinkedList is a data structure that contains a head and tail.
// It has two main operations: add and remove.
// Add: adds an element to the end of the list, while remove removes the first element from the list.
// Remove: removes the first element from the list.
// IsEmpty: returns true if the list is empty, otherwise false.
// Size: returns the number of elements in the list.

// SinglyLinkedList interface
type ISinglyLinkedList interface {
	Add(element int)
	Remove() (int, error)
	IsEmpty() bool
	Size() int
}

// Node struct
type Node struct {
	value int
	next  *Node
}

// SinglyLinkedList struct
type SinglyLinkedList struct {
	head *Node
	size int
}

func (s *SinglyLinkedList) IsEmpty() bool {
	return s.head == nil
}

// Add element at the head of the list.
func (s *SinglyLinkedList) Add(element int) {
	node := &Node{value: element}

	if s.IsEmpty() {
		s.head = node
		return
	}
	s.head.next = node
	s.head = node
	s.size++
}

// Remove element from the head of the list.
func (s *SinglyLinkedList) Remove() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("list is empty")
	}

	element := s.head.value
	s.head = s.head.next
	s.size--
	return element, nil
}

// Size returns the number of elements in the list.
func (s *SinglyLinkedList) Size() int {
	return s.size
}

// CreateSinglyLinkedList creates a new singly linked list.

func CreateSinglyLinkedList() ISinglyLinkedList {
	return &SinglyLinkedList{}
}
