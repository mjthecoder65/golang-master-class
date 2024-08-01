package algorithms

// Binary search tree traversal
// Pre-order traversal: root -> left -> right
// In-order traversal: left -> root -> right
// Post-order traversal: left -> right -> root

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

// This returns elements in ascending order.
func inorderTraversal(root *BSTNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(inorderTraversal(root.left), root.value), inorderTraversal(root.right)...)
}

func preorderTraversal(root *BSTNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(preorderTraversal(root.left)), preorderTraversal(root.right)...)
}

func postorderTraversal(root *BSTNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(postorderTraversal(root.left)), postorderTraversal(root.right)...)
}
