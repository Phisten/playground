package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	stack := []*TreeNode{root}
	var cur *TreeNode
	for len(stack) > 0 {
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if cur == nil {
			continue
		}

		ans = append(ans, cur.Val)
		stack = append(stack, cur.Right, cur.Left)
	}

	return ans
}
