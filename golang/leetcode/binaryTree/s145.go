package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node != nil {
			recursion(node.Left)
			recursion(node.Right)
			ans = append(ans, node.Val)
		}
	}
	recursion(root)

	return ans
}
