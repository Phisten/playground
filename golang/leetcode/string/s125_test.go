package leetcode

import (
	"fmt"
	"testing"
)

func Test_s125(t *testing.T) {
	fmt.Printf("Test_s125\n")

	type FuncPair struct {
		call func(s string) bool
		note string
	}

	funcs := []FuncPair{
		{func(s string) bool {
			return isPalindrome(s)
		}, "target"},
	}

	type Pair struct {
		got      bool
		expected bool
		note     string
	}

	for _, Func := range funcs {
		paris := []Pair{
			{Func.call("A man, a plan, a canal:Panama"), true, "1"},
			{Func.call("race a car"), false, "2"},
			{Func.call("0P"), false, "3"},
		}

		for _, v := range paris {
			if v.expected != v.got {
				t.Errorf("func:%v Expected %v, but got %v, note:%v", Func.note, v.expected, v.got, v.note)
			}
		}
	}
}

//	A 到 Z 的 ASCII 值是 65 到 90
//
// a 到 z 的 ASCII 值是 97 到 122。
// 數字從 48 到 57
func isPalindrome(s string) bool {

	l := len(s)
	left := 0
	right := l - 1

	for left < right {
		lc := s[left]

		if lc >= 65 && lc <= 90 {
			lc += 32
		}

		if (lc >= 97 && lc <= 122) || (lc >= 48 && lc <= 57) {
			rc := s[right]

			if rc >= 65 && rc <= 90 {
				rc += 32
			}

			if (rc >= 97 && rc <= 122) || (rc >= 48 && rc <= 57) {
				if lc != rc {
					return false
				} else {
					left++
					right--
				}
			} else {
				right--
			}
		} else {
			left++
		}
	}

	return true
}
