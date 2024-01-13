package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

func Test_s1704(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := halvesAreAlike
	paris := []Pair{
		{Func("book"), true, ""},
		{Func("textbook"), false, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

// 可以優化 hashtable+鉗形指標
func halvesAreAlike(s string) bool {
	count := len(s) / 2
	left, right := s[:count], s[count:]
	target := "aeiouAEIOU"
	valid := 0
	for i := 0; i < count; i++ {
		if strings.IndexByte(target, left[i]) > -1 {
			valid++
		}
		if strings.IndexByte(target, right[i]) > -1 {
			valid--
		}
	}
	return valid == 0
}
