package leetcode

import (
	"reflect"
	"testing"
)

func Test_s662(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	codec := Codec{}

	Func := widthOfBinaryTree
	paris := []Pair{
		{Func(codec.deserialize("1,3,2,5,3,null,9")), 4, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func widthOfBinaryTree(root *TreeNode) int {
	lvIdx := [][]int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(tn *TreeNode, lv int, curIdx int) {
		if tn == nil {
			return
		}

		if len(lvIdx) == lv {
			lvIdx = append(lvIdx, []int{curIdx, curIdx})
		}
		lvIdx[lv][0] = min(lvIdx[lv][0], curIdx)
		lvIdx[lv][1] = max(lvIdx[lv][1], curIdx)

		dfs(tn.Left, lv+1, curIdx*2-1)
		dfs(tn.Right, lv+1, curIdx*2)
	}
	dfs(root, 0, 1)

	ans := 0
	for _, v := range lvIdx {
		ans = max(ans, v[1]-v[0]+1)
	}

	return ans
}
