package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func s236_lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var dfs func(node *TreeNode) (*TreeNode, bool, bool)
	dfs = func(node *TreeNode) (*TreeNode, bool, bool) {
		if node == nil {
			return nil, false, false
		}

		left, pInLeft, qInLeft := dfs(node.Left)
		if pInLeft && qInLeft {
			return left, pInLeft, qInLeft
		}
		right, pInRight, qInRight := dfs(node.Right)
		if pInRight && qInRight {
			return right, pInRight, qInRight
		}

		if (pInLeft || pInRight) && (qInLeft || qInRight) {
			return node, true, true
		}

		var ans *TreeNode
		if p == node || q == node {
			ans = node
		}

		return ans, p == node || pInLeft || pInRight, q == node || qInLeft || qInRight
	}
	lca, _, _ := dfs(root)
	return lca
}
