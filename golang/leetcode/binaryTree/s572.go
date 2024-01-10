package leetcode

// x先找到subRoot, 在進詳細匹配
// 		沒提到數字會不會重複

// 每個節點都要送詳細匹配
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	} else if root.Val == subRoot.Val {
		if tryMatch(root.Left, subRoot.Left) && tryMatch(root.Right, subRoot.Right) {
			return true
		}
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func tryMatch(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	} else if root.Val == subRoot.Val {
		return tryMatch(root.Left, subRoot.Left) && tryMatch(root.Right, subRoot.Right)
	}
	return false
}
