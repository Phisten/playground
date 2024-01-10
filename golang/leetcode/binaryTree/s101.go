package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// s100 same tree 很像
func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(lTree *TreeNode, rTree *TreeNode) bool {
	if lTree == nil || rTree == nil {
		return lTree == nil && rTree == nil
	} else if lTree.Val == rTree.Val {
		return isMirror(lTree.Left, rTree.Right) && isMirror(lTree.Right, rTree.Left)
	} else {
		return false
	}
}
