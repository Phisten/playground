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
	Func := minSteps
	paris := []Pair{
		{Func("bab", "aba"),  1, ""},
		{Func("leetcode", "practice"),  5, ""},
		{Func("anagram", "mangaar"),  0, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func minSteps(s string, t string) int {
	
	target := map[rune]int{}

	ans := 0

	for _, v := range t {
		target[v] += 1
	}
	for _, v := range s {
		if  target[v]  > 0 {
			target[v] -= 1
		}
	}
	for _, v := range target {
		ans += v
	}
	return ans

}
