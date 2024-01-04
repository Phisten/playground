package leetcode

import (
	"reflect"
	"testing"
)

func Test_s70(t *testing.T) {
	type Pair struct {
		got      int
		expected int
		note     string
	}

	Func := longestPalindrome
	paris := []Pair{
		{Func(2), 2, ""},
		{Func(3), 3, ""},
		{Func(1), 1, ""},
		// {Func(10), 5, ""},
	}

	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

// 最長迴文, 大小寫敏感
// HT儲存各種字元數量 偶數就輸出（每次偶數+2）
// 加總數量<s的場合+1 (回文正中間可放任意字元)
func longestPalindrome(s string) int {
	l := len(s)
	dp := make(map[byte]int, l)
	longest := 0

	for i := 0; i < l; i++ {
		dp[s[i]] += 1
		if dp[s[i]]%2 == 0 {
			longest += 2
		}
	}
	if longest < l {
		longest += 1
	}
	return longest
}
