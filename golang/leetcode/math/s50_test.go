package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func Test_s50(t *testing.T) {
	fmt.Printf("Test_s50\n")
	round5 := func(f float64) float64 {
		return math.Round(f*10000.0) / 10000.0
	}

	type FuncPair struct {
		call func(x float64, n int) float64
		note string
	}

	funcs := []FuncPair{
		{func(x float64, n int) float64 {
			return round5(myPow(x, n))
		}, "target"},
		{func(x float64, n int) float64 {
			return round5(myPow_sdk(x, n))
		}, "sdk"},
	}

	type Pair struct {
		got      float64
		expected float64
		note     string
	}

	for _, Func := range funcs {
		paris := []Pair{
			{Func.call(2.0, 10), 1024.0, ""},
			{Func.call(2.1, 3), 9.261, ""},
			{Func.call(2.0, -2), 0.25, ""},
			{Func.call(2.0, 0), 1.0, ""},
			{Func.call(8.95371, -1), 0.1117, ""},
			{Func.call(1.84364, -14), 0.0002, ""},
			{Func.call(-13.62608, 3), -2529.955, ""},
		}

		for _, v := range paris {
			if v.expected != v.got {
				t.Errorf("func:%v Expected %v, but got %v, note:%v", Func.note, v.expected, v.got, v.note)
			}
		}
	}
}

// 1 2 4 8 16 32 64 128 256 512 1024
// 0 1 2 3 4  5  6  7   8   9   10
// 2^3^2 = 8*8 = 64 = 2^6
//
// 64 16 4 2 1.4
// 1  -1 2 3 4
func myPow(x float64, n int) float64 {
	// fmt.Printf("> x=%v, n=%v\n", x, n)
	if n == 0 {
		return 1
	}
	negative := n < 0
	if negative {
		n = -n
	}
	dp := make([]float64, n+1)

	lastN := n
	ans := x
	dp[1] = x

	for i := 1; i <= n/2; {
		ans = dp[i] * dp[i]

		i *= 2
		dp[i] = ans
		lastN = n - i
		// fmt.Printf("i=%v, lastN=%v, ans=%v\n", i, lastN, ans)
	}

	if n != 1 {
		for i := lastN; i >= 1; {
			if dp[i] > 0 {
				ans *= dp[i]
				lastN -= i
				// fmt.Printf("i=%v, lastN=%v, ans=%v\n", i, lastN, ans)
				i = lastN
			} else {
				i--
			}
		}
	}

	if negative {
		ans = 1.0 / ans
	}

	return ans
}

func myPow_sdk(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
