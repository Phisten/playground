package main

import (
	"fmt"
	leetcode "sample-app/leetcode/leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	root := leetcode.ArrayToTreeNode([]interface{}{1, 2, 3, 4, 5})
	fmt.Printf("case 3 -> %v\n", diameterOfBinaryTree(root))
	root = leetcode.ArrayToTreeNode([]interface{}{1, 2, nil, 4, 5})
	fmt.Printf("case 2 -> %v\n", diameterOfBinaryTree(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	return maxLen(root).diameter
}

type Pair struct {
	depth    int
	diameter int
}

func maxLen(root *TreeNode) *Pair {
	ans := &Pair{
		depth: 0, diameter: 0,
	}

	if root == nil {
		return ans
	}

	left := 0
	if root.Left != nil {
		leftAns := maxLen(root.Left)
		left = 1 + leftAns.depth
		ans.diameter = max(leftAns.diameter, ans.diameter)
	}
	right := 0
	if root.Right != nil {
		rightAns := maxLen(root.Right)
		right = 1 + rightAns.depth
		ans.diameter = max(rightAns.diameter, ans.diameter)
	}

	ans.depth = max(left, right)
	ans.diameter = max(left+right, ans.diameter)

	return ans
}
