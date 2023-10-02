package algorithms

// TODO: To be updated soon.....
func BFS(root Node) {
	// Create a queue to store the nodes that need to be visited.
	queue := []Node{root}

	// While the queue is not empty.
	for len(queue) > 0 {
		// Pop a node from the queue.
		currentNode := queue[0]
		queue = queue[1:]

		// Mark the node as visited.
		currentNode.Visited = true

		// Visit the node's children.
		for _, child := range currentNode.Edges {
			if !child.Visited {
				queue = append(queue, child)
			}
		}
	}
}
