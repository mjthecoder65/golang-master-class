package generics

import "errors"

type Queue[T any] struct {
	elements []T
}

func (queue *Queue[T]) Enqueue(element T) {
	queue.elements = append(queue.elements, element)
}

func (queue *Queue[T]) Dequeue() (T, error) {
	frontElement, err := queue.Front()
	if err != nil {
		return frontElement, err
	}
	queue.elements = queue.elements[1:]
	return frontElement, nil
}

func (queue *Queue[T]) Empty() bool {
	return len(queue.elements) == 0
}

func (queue *Queue[T]) Front() (T, error) {
	if queue.Empty() {
		var zeroValue T
		return zeroValue, errors.New("error: queue is empty")
	}
	return queue.elements[0], nil
}
