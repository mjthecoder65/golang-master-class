package leetcode

type Node struct {
	Val  int
	Next *Node
}

func CreateNode(val int) *Node {
	return &Node{Val: val, Next: nil}
}

type LinkedList struct {
	Head *Node
	Size int
}

func (node *LinkedList) FindItem(target int) bool {
	currentNode := node.Head

	for currentNode != nil {
		if currentNode.Val == target {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (node *LinkedList) AddItem(val int) {
	newNode := CreateNode(val)
	if node.Head == nil {
		node.Head = newNode
	} else {
		newNode.Next = node.Head
		node.Head = newNode
	}
}
