package leetcode

// URL: https://leetcode.com/problems/find-mode-in-binary-search-tree/
// Title : Find Mode in Binary Search Tree
// Difficulty: Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// PreOrder: root, left, right
// Inorder : left, root, right (Return elements in ascending order)
// PostOrder: left, right, root

func inorder(root *TreeNode) []int {
	elements := []int{}

	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		elements = append(elements, root.Val)
		inorder(root.Right)
	}

	return elements
}

func findMode(root *TreeNode) []int {
	elements := inorder(root)
	counter := make(map[int]int)
	maxCounter := 0

	for _, element := range elements {
		counter[element]++
		if counter[element] > maxCounter {
			maxCounter = counter[element]
		}
	}

	result := []int{}

	for element, count := range counter {
		if count == maxCounter {
			result = append(result, element)
		}
	}
	return result
}
