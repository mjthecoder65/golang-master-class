package leetcode_50

// URL: https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/1365377101/

func PreorderTraversal(root *TreeNode) []int {
	var preorder func(root *TreeNode)

	values := []int{}
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		values = append(values, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)

	return values
}
