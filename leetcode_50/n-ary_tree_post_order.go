package leetcode_50

// URL: https://leetcode.com/problems/n-ary-tree-postorder-traversal

func Postorder(root *Node) []int {
	result := []int{}
	var postorder func(root *Node)

	postorder = func(root *Node) {
		if root == nil {
			return
		}

		for _, node := range root.Children {
			postorder(node)
		}
		result = append(result, root.Val)
	}
	postorder(root)

	return result
}
