package leetcode

import (
	"testing"
)

func Test_s235(t *testing.T) {
}

// 其實只要檢測pq是否在全左或全右, 否則就是答案
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	leftN, rightN := p, q
	if p.Val > q.Val {
		leftN, rightN = q, p
	}

	// match
	if leftN.Val <= root.Val && root.Val <= rightN.Val {
		return root
	}

	if root.Left == nil || root.Right == nil {
		return nil
	} else {
		// all left
		leftFind := lowestCommonAncestor(root.Left, leftN, rightN)
		if leftFind != nil {
			return leftFind
		}
		// all right
		rightFind := lowestCommonAncestor(root.Right, leftN, rightN)
		return rightFind
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
