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
	stack := []*TreeNode{}

	var cur *TreeNode = root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur, cur.Right)
			cur, cur.Left, cur.Right = cur.Left, nil, nil
		}

		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if cur != nil && cur.Right == nil && cur.Left == nil {
			ans = append(ans, cur.Val)
			cur = nil
		}
	}

	return ans
}
