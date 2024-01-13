package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var dsf func(root *TreeNode, targetSum int) [][]int
	dsf = func(root *TreeNode, targetSum int) [][]int {
		var path [][]int = [][]int{}
		if root == nil {
			return path
		}

		lastSum := targetSum - root.Val
		// is leaf
		if lastSum == 0 && root.Left == nil && root.Right == nil {
			return append(path, []int{root.Val})
		} else {

			left := dsf(root.Left, lastSum)
			for _, v := range left {
				path = append(path, append([]int{root.Val}, v...))
			}

			right := dsf(root.Right, lastSum)
			for _, v := range right {
				path = append(path, append([]int{root.Val}, v...))
			}
		}

		return path
	}

	return dsf(root, targetSum)
}
