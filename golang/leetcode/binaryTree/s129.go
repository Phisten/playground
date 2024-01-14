package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, prev int) int {
		if node == nil {
			return 0
		}

		curValue := prev*10 + node.Val

		left := 0
		if node.Left != nil {
			left = dfs(node.Left, curValue)
		}
		right := 0
		if node.Right != nil {
			right = dfs(node.Right, curValue)
		}

		if left+right == 0 {
			return curValue
		}
		return left + right
	}
	return dfs(root, 0)
}
