package leetcode

func countNodes(root *TreeNode) int {
	fullNodeCount := 0
	if root == nil {
		return 0
	}

	// 1+2+4 = 7 = 2^3-1
	// 1+2+4+8 = 15 = 2^4 -1
	fullNodeCount = 1
	depth := 0
	cur := root
	for {
		if cur != nil {
			depth++
			cur = cur.Left
			fullNodeCount *= 2
		} else {
			fullNodeCount--
			break
		}
	}

	endDfs := false
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, curDepth int) int {
		if endDfs {
			return 0
		}
		if node == nil {
			if depth == curDepth {
				return 1
			} else {
				endDfs = true
				return 0
			}
		}

		right := dfs(node.Right, curDepth+1)
		left := dfs(node.Left, curDepth+1)
		return right + left
	}
	sub := dfs(root, 1)

	return fullNodeCount - sub
}

func s222_On_countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := countNodes(root.Right)
	l := countNodes(root.Left)
	return l + r + 1
}
