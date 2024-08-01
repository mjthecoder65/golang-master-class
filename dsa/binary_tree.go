package dsa

// Binary Tree is a data structure that contains a root node and two children nodes.
// It has two main operations: add and remove.
// Add: adds an element to the tree, while remove removes an element from the tree.
// Remove: removes an element from the tree.
// IsEmpty: returns true if the tree is empty, otherwise false.
// Size: returns the number of elements in the tree.

// Binary Tree interface
type IBinaryTree interface {
	Add(element int)
	Remove(element int) (int, error)
	IsEmpty() bool
	Size() int
	Contains(element int) bool
}

// Node struct
type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

// Binary Tree struct
type BinaryTree struct {
	root *BSTNode
	size int
}

func (b *BinaryTree) IsEmpty() bool {
	return b.root == nil
}

// Add element to the tree.
func (b *BinaryTree) Add(element int) {
	// TODO: Implement the logic here.
}

func breadthFirstSearch(root *BSTNode) {
	if root == nil {
		return
	}
	
}
