package main

import (
	"fmt"
	leetcode "sample-app/leetcode/leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	root := leetcode.ArrayToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Printf("case T -> %v\n", isBalanced(root))
	root = leetcode.ArrayToTreeNode([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})
	fmt.Printf("case F -> %v\n", isBalanced(root))
	root = leetcode.ArrayToTreeNode([]interface{}{1, 2, 2, 3, nil, nil, 3, 4, nil, nil, 4})
	fmt.Printf("case F -> %v\n", isBalanced(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	nums := isBalancedNums(root)

	return nums < 100001
}

func isBalancedNums(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := 0
	if root.Left != nil {
		left = 1 + isBalancedNums(root.Left)
	}
	right := 0
	if root.Right != nil {
		right = 1 + isBalancedNums(root.Right)
	}

	// fmt.Printf("root:%v  left %v, right %v\n", root.Val, left, right)

	big := left
	small := right
	if big < small {
		big = right
		small = left
	}

	if big-small > 1 {
		return 100001
	}
	return big
}
