package leetcode

// URL: https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/description/
// Title: Count Nodes Equal to Average of Subtree
// Difficulty: Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func averageOfSubtree(root *TreeNode) int {
	count := 0

	var getSubtreeAverage func(node *TreeNode) (int, int)

	getSubtreeAverage = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftSum, leftCount := getSubtreeAverage(node.Left)
		rightSum, rightCount := getSubtreeAverage(node.Right)

		totalSum := leftSum + rightSum + node.Val
		totalCount := leftCount + rightCount + 1
		subtreeAverage := totalSum / totalCount

		if node.Val == subtreeAverage {
			count++
		}

		return totalSum, totalCount
	}

	getSubtreeAverage(root)

	return count
}
