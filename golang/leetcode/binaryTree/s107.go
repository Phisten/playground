package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}

	type LvNode struct {
		lv int
		n  *TreeNode
	}
	queue := []*LvNode{}
	if root != nil {
		queue = append(queue, &LvNode{lv: 0, n: root})
	}

	i := 0
	for i < len(queue) {
		n, lv := queue[i].n, queue[i].lv

		if len(ans) == lv {
			ans = append(ans, []int{n.Val})
		} else {
			ans[lv] = append(ans[lv], n.Val)
		}

		if n.Left != nil {
			queue = append(queue, &LvNode{n: n.Left, lv: lv + 1})
		}
		if n.Right != nil {
			queue = append(queue, &LvNode{n: n.Right, lv: lv + 1})
		}

		i++
	}

	res := [][]int{}
	for i := 0; i < len(ans); i++ {
		res = append(res, ans[len(ans)-i-1])
	}

	return res
}
