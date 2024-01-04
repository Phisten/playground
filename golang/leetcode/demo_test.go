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

func problemsFunc(a string, b string) string {
	return "10110"
}
