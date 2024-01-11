package leetcode

func maxAncestorDiff(root *TreeNode) int {
	maxLength := 0

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		minVal, maxVal := node.Val, node.Val

		if node.Left != nil {
			leftMin, leftMax := dfs(node.Left)
			minVal = min(minVal, leftMin)
			maxVal = max(maxVal, leftMax)
		}
		if node.Right != nil {
			rightMin, rightMax := dfs(node.Right)
			minVal = min(minVal, rightMin)
			maxVal = max(maxVal, rightMax)
		}

		curCost := max(Abs(node.Val-maxVal), Abs(node.Val-minVal))
		maxLength = max(maxLength, curCost)

		return minVal, maxVal
	}
	dfs(root)

	return maxLength
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
