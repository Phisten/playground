package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	type LvNode struct {
		n  *TreeNode
		lv int
	}
	q := []*LvNode{{n: root, lv: 0}}

	for i := 0; i < len(q); i++ {

		cur := q[i]
		if cur.lv == len(ans) {
			ans = append(ans, []int{cur.n.Val})
		} else {
			if cur.lv%2 == 0 {
				ans[cur.lv] = append(ans[cur.lv], cur.n.Val)
			} else {
				ans[cur.lv] = append([]int{cur.n.Val}, ans[cur.lv]...)
			}
		}

		if cur.n.Left != nil {
			q = append(q, &LvNode{n: cur.n.Left, lv: cur.lv + 1})
		}
		if cur.n.Right != nil {
			q = append(q, &LvNode{n: cur.n.Right, lv: cur.lv + 1})
		}
	}

	return ans
}
