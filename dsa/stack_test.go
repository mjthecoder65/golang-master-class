package dsa

// Testing Stack Data structure.

import (
	"testing"

	"github.com/mjthecoder65/golang-master-class/utils"
	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	stack := CreateStack()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)
	stack.Push(firstElement)
	stack.Push(secondElement)
	require.Equal(t, secondElement, stack.Top())
	require.Equal(t, 2, stack.Size())
}

func TestPop(t *testing.T) {
	stack := CreateStack()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)
	stack.Push(firstElement)
	stack.Push(secondElement)
	require.Equal(t, secondElement, stack.Top())
	require.Equal(t, 2, stack.Size())

	stack.Pop()

	require.Equal(t, firstElement, stack.Top())
	require.Equal(t, 1, stack.Size())

	stack.Pop()
	require.Equal(t, 0, stack.Size())
	// require.PanicsWithError(t, "Stack is empty", func() { stack.Pop() })
}

func TestCle(t *testing.T) {
	stack := CreateStack()

	firstElement := utils.RandomInt(1, 100)
	secondElement := utils.RandomInt(1, 100)
	stack.Push(firstElement)
	stack.Push(secondElement)

	stack.Clear()
	require.Equal(t, 0, stack.Size())
}
