package leetcode

// URL: https://leetcode.com/problems/find-mode-in-binary-search-tree/description/

// Title: Find Mode in Binary Search Tree
// Difficulty: Easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	counter := make(map[int]int)
	maxCount := 0
	var inorderTraversal func(root *TreeNode)

	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorderTraversal(root.Left)
		counter[root.Val]++
		if counter[root.Val] > maxCount {
			maxCount = counter[root.Val]
		}
		inorderTraversal(root.Right)
	}

	inorderTraversal(root) // Call inorderTraversal to populate counter

	modes := []int{}

	for element, count := range counter {
		if count == maxCount {
			modes = append(modes, element)
		}
	}

	return modes
}
