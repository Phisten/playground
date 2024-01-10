package leetcode

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(inorder)
	if length == 0 {
		return nil
	}

	mid := slices.Index(inorder, postorder[length-1])

	return &TreeNode{
		Val:   postorder[length-1],
		Left:  buildTree(inorder[:mid], postorder[:mid]),
		Right: buildTree(inorder[mid+1:], postorder[mid:length-1]),
	}
}
