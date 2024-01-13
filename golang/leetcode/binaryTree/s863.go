package leetcode

// 先建立parent map 在做三向尋訪
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := map[*TreeNode]*TreeNode{}

	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn != nil {
			if tn.Left != nil {
				parent[tn.Left] = tn
				dfs(tn.Left)
			}
			if tn.Right != nil {
				parent[tn.Right] = tn
				dfs(tn.Right)
			}
		}
	}
	dfs(root)

	ans := []int{}

	var find func(*TreeNode, int, *TreeNode)
	find = func(tn1 *TreeNode, lastStep int, lastNode *TreeNode) {
		if tn1 == nil {
			return
		}
		if lastStep == 0 {
			ans = append(ans, tn1.Val)
		} else {
			for _, v := range []*TreeNode{tn1.Left, tn1.Right, parent[tn1]} {
				if lastNode != v {
					find(v, lastStep-1, tn1)
				}
			}

		}
	}
	find(target, k, nil)

	return ans
}
