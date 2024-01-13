package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遞迴做左右值交換
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invertTree(root.Left)
		invertTree(root.Right)

		temp := root.Left
		root.Left = root.Right
		root.Right = temp
	}

	return root
}
