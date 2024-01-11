package leetcode

func maxAncestorDiff(root *TreeNode) int {
	maxLength := 0

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		minVal, maxVal := node.Val, node.Val

		if node.Left != nil {
			leftMin, leftMax := dfs(node.Left)
			minVal = min(minVal, node.Left.Val, leftMin)
			maxVal = max(maxVal, node.Left.Val, leftMax)
		}
		if node.Right != nil {
			rightMin, rightMax := dfs(node.Right)
			minVal = min(minVal, node.Right.Val, rightMin)
			maxVal = max(maxVal, node.Right.Val, rightMax)
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
