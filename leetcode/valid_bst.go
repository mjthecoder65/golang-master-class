package leetcode

import "math"

// URL: https://leetcode.com/problems/validate-binary-search-tree/
// Title : Validate Binary Search Tree

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	// For the left subtree, the maximum value is the current node's value.
	// For the right subtree, the minimum value is the current node's value.
	return isBST(node.Left, min, node.Val) && isBST(node.Right, node.Val, max)
}
