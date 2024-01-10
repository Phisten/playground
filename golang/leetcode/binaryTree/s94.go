package leetcode

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	s := NodeStack{}
	var cur *TreeNode = root

	for !s.isEmpty() || cur != nil {
		for cur != nil {
			s.push(cur)
			cur = cur.Left
		}

		cur = s.pop()
		ans = append(ans, cur.Val)
		cur = cur.Right
	}

	return ans
}
