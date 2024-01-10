package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 左右中 後序
// 輸出max
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	def := root.Val
	var ans *int = &def

	recursion(root, ans)

	return *ans
}

func recursion(root *TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	}

	left := recursion(root.Left, maxPath)
	right := recursion(root.Right, maxPath)

	curPath := root.Val + max(left, right, 0)
	*maxPath = max(*maxPath, curPath, root.Val+left+right)

	return curPath
}
