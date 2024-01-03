package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

		if index < len(arr) && arr[index] != nil {
			val, _ := arr[index].(int)
			node.Val = val
			index++

			if len(arr)-index-1 > len(queue) {
				node.Left = &TreeNode{}
				queue = append(queue, node.Left)
			}
		}

		if index < len(arr) && arr[index] != nil {
			val, _ := arr[index].(int)
			node.Val = val
			index++

			if len(arr)-index-1 > len(queue) {
				node.Right = &TreeNode{}
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
