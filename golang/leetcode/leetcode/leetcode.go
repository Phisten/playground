package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 錯的
func ArrayToTreeNode(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{}
	queue := []*TreeNode{root}
	index := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if arr[index] != nil {
			val, _ := arr[index].(int)
			node.Val = val
			index++
			node.Left = &TreeNode{}
			queue = append(queue, node.Left)
		}

		if index < len(arr) && arr[index] != nil {
			val, _ := arr[index].(int)
			node.Val = val
			index++
			node.Right = &TreeNode{}
			queue = append(queue, node.Right)
		}
	}

	return root
}
