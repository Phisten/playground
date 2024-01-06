package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//
type Node struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	ans = append(ans, []int{})
	queue := []Node{{root, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.node != nil {
			if cur.level == len(ans) {
				ans = append(ans, []int{})
			}
			ans[cur.level] = append(ans[cur.level], cur.node.Val)

			queue = append(queue,
				Node{node: cur.node.Left, level: cur.level + 1},
				Node{node: cur.node.Right, level: cur.level + 1},
			)
		}
	}

	return ans
}
