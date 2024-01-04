package leetcode

import (
	"reflect"
	leetcode "sample-app/leetcode/BST"
	"testing"
)

func Test_s104(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := problemsFunc
	paris := []Pair{
		{Func("1000", "1110"), "10110", ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

type TreeNode = leetcode.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
