package leetcode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := countNodes(root.Right)
	l := countNodes(root.Left)
	return l + r + 1
}
