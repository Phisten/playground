package leetcode

import "slices"

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }

3,9,20,15,7
9,3,15,20,7

	      3
	   9    20
	1      15 7

pre:3,9,1,20,15,7
in :1,9,3,15,20,7
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}

	midIdx := slices.Index(inorder, preorder[0])
	leftLen := midIdx + 1

	buildNode := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:leftLen], inorder[:midIdx]),
		Right: buildTree(preorder[midIdx+1:], inorder[midIdx+1:]),
	}

	return buildNode
}
