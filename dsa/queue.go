package dsa

import "errors"

// Queue is a data structure that follows the FIFO (First In First Out) principle.
// It has two main operations: enqueue and dequeue.
// Enqueue: adds an element to the end of the queue, while dequeue removes the first element from the queue.
// Dequeue: removes the first element from the queue.
// Peek: returns the first element from the queue without removing it.
// IsEmpty: returns true if the queue is empty, otherwise false.
// Size: returns the number of elements in the queue.
// Clear: removes all elements from the queue.

// Queue interface.
type IQueue interface {
	Enqueue(element int)
	Dequeue() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
	Clear()
}

// Queue struct
type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]

	return element, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	return q.elements[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Clear() {
	q.elements = make([]int, 0)
}

// Note: Return a pointer to struct is a common design choice in Go for several reasons.
// One of the reasons is that it allows us to modify the struct fields directly without
// having to return the struct again.
func CreateQueue() IQueue {
	return &Queue{}
}
