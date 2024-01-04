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

	Func := climbStairs
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

// 1 2 3 5 8 費氏數列？？？
// dp or recursion
func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i < 3 {
			dp[i] = i
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}
