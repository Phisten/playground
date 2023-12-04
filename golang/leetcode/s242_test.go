package leetcode

import (
	"fmt"
	"testing"
)

func Test_s242(t *testing.T) {
	fmt.Printf("Test_s242\n")

	type FuncPair struct {
		call func(s string, t string) bool
		note string
	}

	funcs := []FuncPair{
		{func(s string, t string) bool {
			return isAnagram(s, t)
		}, "target"},
	}

	type Pair struct {
		got      bool
		expected bool
		note     string
	}

	for _, Func := range funcs {
		paris := []Pair{
			{Func.call("anagram", "nagaram"), true, "1"},
			{Func.call("rat", "car"), false, "2"},
			{Func.call("ab", "a"), false, "3"},
		}

		for _, v := range paris {
			if v.expected != v.got {
				t.Errorf("func:%v Expected %v, but got %v, note:%v", Func.note, v.expected, v.got, v.note)
			}
		}
	}
}
func isAnagram(s string, t string) bool {
	hMap := make(map[rune]int, 255)

	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		hMap[v]++
	}
	for _, v := range t {
		hMap[v]--
		if hMap[v] < 0 {
			return false
		}
	}
	return true
}
