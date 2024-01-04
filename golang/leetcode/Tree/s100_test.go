package leetcode

import (
	"reflect"
	"testing"
)

func Test_s100(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := isSameTree
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

// recursion吧？
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
