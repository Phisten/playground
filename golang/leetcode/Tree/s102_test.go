package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//

/*
3
9 6 8    2 5 7
3
9 2
6 8  5 7
*/
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	arr := treeToArray(root)
	fmt.Println(ans)
	l := len(arr)
	for i := 1; i <= l; i *= 2 {
		curAns := []int{}
		for j := i / 2; j < i; j++ {
			if arr[j] != nil {
				curAns = append(curAns, arr[j].(int))
			}
		}
		ans = append(ans, curAns)
	}

	return ans
}

func treeToArray(root *TreeNode) []interface{} {
	ans := []interface{}{}
	if root == nil {
		ans = append(ans, nil)
		return ans
	}

	ans = append(ans, root.Val)

	left := treeToArray(root.Left)
	ans = append(ans, left...)
	right := treeToArray(root.Right)
	ans = append(ans, right...)

	return ans
}
