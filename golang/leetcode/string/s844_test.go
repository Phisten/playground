package leetcode

import (
	"reflect"
	"testing"
)

func Test_s844(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := backspaceCompare
	paris := []Pair{
		{Func("ab##", "a#b##"), true, ""},
		{Func("bxj##tw", "bxj###tw"), false, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

// 由後往前檢查相等, 碰到#就+1忽略字元字數
func backspaceCompare(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)

	sIdx := sLen - 1
	tIdx := tLen - 1

	sSkip := 0
	tSkip := 0

	for {
		if sIdx >= 0 && s[sIdx] == '#' {
			sSkip++
			sIdx--
		} else if sSkip > 0 {
			sSkip--
			sIdx--
		} else {
			// s ready to match
			if tIdx >= 0 && t[tIdx] == '#' {
				tSkip++
				tIdx--
			} else if tSkip > 0 {
				tSkip--
				tIdx--
			} else {
				// try match
				if sIdx < 0 || tIdx < 0 {
					return sIdx < 0 && tIdx < 0
				} else if s[sIdx] == t[tIdx] {
					sIdx--
					tIdx--
				} else {
					return false
				}
			}
		}
	}
}
