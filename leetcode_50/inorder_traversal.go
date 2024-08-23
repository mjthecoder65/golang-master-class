package leetcode_50

// URL: https://leetcode.com/problems/binary-tree-inorder-traversal/

func InorderTraversal(root *TreeNode) []int {
	result := []int{}

	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return result
}
