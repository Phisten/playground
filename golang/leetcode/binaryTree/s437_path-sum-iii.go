package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func s437_pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	queue := []*TreeNode{}

	var dsf func(*TreeNode, int, bool)
	dsf = func(root *TreeNode, curSum int, scanNode bool) {
		if root == nil {
			return
		}
		if scanNode {
			queue = append(queue, root)
		}

		lastSum := curSum - root.Val
		if lastSum == 0 {
			ans++
		}
		dsf(root.Left, lastSum, scanNode)
		dsf(root.Right, lastSum, scanNode)
	}
	dsf(root, targetSum, true)
	for i := 1; i < len(queue); i++ {
		dsf(queue[i], targetSum, false)
	}

	return ans
}
