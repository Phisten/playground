package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodes := []int{}
	var prev *TreeNode
	ans := true

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		nodes = append(nodes, node.Val)
		if prev != nil && prev.Val >= node.Val {
			ans = false
		}
		prev = node
		fmt.Println(node.Val)

		dfs(node.Right)

	}
	dfs(root)

	return ans
}
