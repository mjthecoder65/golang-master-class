package leetcode_50

// URL: https://leetcode.com/problems/binary-tree-postorder-traversal/description/

func PostorderTraversal(root *TreeNode) []int {
	// Postorder: left, right, root;
	var postorder func(root *TreeNode)

	values := []int{}

	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		values = append(values, root.Val)
	}

	postorder(root)

	return values
}
