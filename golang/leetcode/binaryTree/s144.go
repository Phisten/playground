package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func s144_preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node != nil {
			ans = append(ans, node.Val)
			recursion(node.Left)
			recursion(node.Right)
		}
	}
	recursion(root)

	return ans
}
