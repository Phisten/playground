package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minLvSide := 100001
	if root.Left != nil {
		minLvSide = minDepth(root.Left)
	}
	if root.Right != nil {
		minLvSide = min(minLvSide, minDepth(root.Right))
	}

	if minLvSide < 100001 {
		return 1 + minLvSide
	} else {
		return 1
	}
}
