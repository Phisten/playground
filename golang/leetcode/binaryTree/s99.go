package leetcode

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }


BST中序走訪就是遞增序
紀錄前一個點的位置中序走訪就可以找到錯誤的點
找到第一組錯誤時外部緩存兩個節點（第二個不一定錯）
繼續往後找有無第二組錯誤, 有就蓋掉第二個緩存的節點
全部訪問後交換

只有兩個節點錯置
第一組錯誤搜到 第一個節點一定是錯的
若搜到第二組錯誤 第二組的第二格節點是錯的
*/

func recoverTree(root *TreeNode) {
	var prevNode, leftErr, rightErr *TreeNode

	var recursion func(cur *TreeNode)
	recursion = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		recursion(cur.Left)

		if prevNode != nil && prevNode.Val > cur.Val {
			if leftErr == nil {
				leftErr = prevNode
			}
			rightErr = cur
		}
		prevNode = cur

		recursion(cur.Right)
	}

	recursion(root)
	leftErr.Val, rightErr.Val = rightErr.Val, leftErr.Val
}
