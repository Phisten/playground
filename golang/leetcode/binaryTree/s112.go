package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	lastSum := targetSum - root.Val
	if lastSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, lastSum) || hasPathSum(root.Right, lastSum)
}
