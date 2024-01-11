package leetcode

/**
看錯題目, 不是走右邊 而是水平投影由右往左看每個lv第一個node
*/

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	type NodeLv struct {
		node *TreeNode
		lv   int
	}
	queue := []*NodeLv{{
		node: root,
		lv:   0,
	}}

	for i := 0; i < len(queue); i++ {
		cur := queue[i]

		if cur.lv == len(ans) {
			ans = append(ans, cur.node.Val)
		}
		if cur.node.Right != nil {
			queue = append(queue, &NodeLv{node: cur.node.Right, lv: cur.lv + 1})
		}
		if cur.node.Left != nil {
			queue = append(queue, &NodeLv{node: cur.node.Left, lv: cur.lv + 1})
		}
	}

	return ans
}
