package leetcode

import (
	"testing"
)

func Test_s235(t *testing.T) {
	// type Pair struct {
	// 	got      int
	// 	expected int
	// 	note     string
	// }

	// Func := lowestCommonAncestor
	// paris := []Pair{
	// 	// leetcode.ArrayToTreeNode([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	// 	// {Func(, 2, 8), 6, ""},
	// }

	// for i, v := range paris {
	// 	if !reflect.DeepEqual(v.expected, v.got) {
	// 		f := "i:%v Expected %v, but got %v, note:%v"
	// 		t.Errorf(f, i, v.expected, v.got, v.note)
	// 	}
	// }
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
