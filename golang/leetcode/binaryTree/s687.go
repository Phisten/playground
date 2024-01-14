package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 0

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		linkLeft := node.Left != nil && node.Left.Val == node.Val
		linkRight := node.Right != nil && node.Right.Val == node.Val

		left := dfs(node.Left)
		right := dfs(node.Right)

		if linkLeft && linkRight {
			maxLen = max(left+right+1, maxLen)
			return max(left, right) + 1
		} else if linkRight {
			maxLen = max(right+1, maxLen)
			return 1 + right
		} else if linkLeft {
			maxLen = max(left+1, maxLen)
			return 1 + left
		}
		maxLen = max(1, maxLen)
		return 1
	}
	dfs(root)

	return maxLen - 1
}
