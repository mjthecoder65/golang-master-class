package algorithms

type Node struct {
	Value   int
	Edges   []Node
	Visited bool
}

// TODO: To be updated soon....
func DFS(root Node) {
	// Create a stack to store the nodes that need to be visited.
	stack := []Node{root}

	// While the stack is not empty.
	for len(stack) > 0 {
		// Pop a node from the stack.
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Mark the node as visited.
		currentNode.Visited = true

		// Visit the node's children.
		for _, child := range currentNode.Edges {
			if !child.Visited {
				stack = append(stack, child)
			}
		}
	}
}
