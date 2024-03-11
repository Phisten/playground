package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func Test_s791(t *testing.T) {
	fmt.Printf("Test_s791\n")

	type FuncPair struct {
		call func(s string, t string) string
		note string
	}

	funcs := []FuncPair{
		{func(s string, t string) string {
			return customSortString(s, t)
		}, "target"},
	}

	type Pair struct {
		got      string
		expected string
		note     string
	}

	for _, Func := range funcs {
		paris := []Pair{
			{Func.call("kqep", "pekeq"), "kqeep", "3"},
			{Func.call("bcafg", "abcd"), "bcad", "2"},
			{Func.call("cba", "abcd"), "cbad", "1"},
		}

		for _, v := range paris {
			if v.expected != v.got {
				t.Errorf("func:%v Expected %v, but got %v, note:%v", Func.note, v.expected, v.got, v.note)
			}
		}
	}
}

func customSortString(order string, s string) string {
	ht1, ht2 := make(map[rune]int), make(map[rune]int)
	ans := ""

	for _, v := range order {
		ht1[v] += 1
	}
	for _, v := range s {
		ht2[v] += 1
	}

	for _, v := range order {
		if count, ok := ht2[v]; ok {
			ans += strings.Repeat(string(v), count)
		}
	}

	for _, v := range s {
		if _, ok := ht1[v]; !ok {
			ans += string(v)
		}
	}

	return ans
}
