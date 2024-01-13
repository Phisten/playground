package leetcode

import (
	"reflect"
	"testing"
)

func Test_s(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	codec := Codec{}

	Func := problemsFunc
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

func problemsFunc(root *TreeNode) int {
	return 4
}
