package dsa

// Creating tests for queue.go

import (
	"testing"

	"github.com/mjthecoder65/golang-master-class/utils"
	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	queue := CreateQueue()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)
	thirdElement := utils.RandomInt(1, 100)

	queue.Enqueue(firstElement)
	queue.Enqueue(secondElement)
	queue.Enqueue(thirdElement)

	peekElement, _ := queue.Peek()
	require.Equal(t, firstElement, peekElement)
	require.Equal(t, 3, queue.Size())
}

func TestDequeue(t *testing.T) {
	queue := CreateQueue()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)

	queue.Enqueue(firstElement)
	queue.Enqueue(secondElement)

	removedElement, err := queue.Dequeue()
	require.Nil(t, err)
	require.Equal(t, firstElement, removedElement)
	require.Equal(t, 1, queue.Size())

	removedElement, err = queue.Dequeue()
	require.Nil(t, err)
	require.Equal(t, secondElement, removedElement)
	require.Equal(t, 0, queue.Size())

	_, err = queue.Dequeue()
	require.NotNil(t, err)

}

func TestPeek(t *testing.T) {
	queue := CreateQueue()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)

	queue.Enqueue(firstElement)
	queue.Enqueue(secondElement)

	peekElement, _ := queue.Peek()
	require.Equal(t, firstElement, peekElement)
	queue.Dequeue()
	peekElement, _ = queue.Peek()
	require.Equal(t, secondElement, peekElement)
}

func TestSize(t *testing.T) {
	queue := CreateQueue()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)

	queue.Enqueue(firstElement)
	queue.Enqueue(secondElement)
	require.Equal(t, 2, queue.Size())
	queue.Dequeue()
	require.Equal(t, 1, queue.Size())
}

func TestIsEmpty(t *testing.T) {
	queue := CreateQueue()

	require.Equal(t, true, queue.IsEmpty())

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)

	queue.Enqueue(firstElement)
	queue.Enqueue(secondElement)

	require.Equal(t, false, queue.IsEmpty())
}

func TestClear(t *testing.T) {
	queue := CreateQueue()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)

	queue.Enqueue(firstElement)
	queue.Enqueue(secondElement)

	require.Equal(t, false, queue.IsEmpty())
	require.Equal(t, 2, queue.Size())

	queue.Clear()

	require.Equal(t, true, queue.IsEmpty())
	require.Equal(t, 0, queue.Size())
}
