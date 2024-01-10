package leetcode

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}

	dst(root, &ans)

	return ans
}

func dst(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	dst(node.Left, ans)

	*ans = append(*ans, node.Val)

	dst(node.Right, ans)
}
