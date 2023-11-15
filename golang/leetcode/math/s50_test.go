package leetcode

import (
	"math"
	"testing"
)

func round5(f float64) float64 {
	return math.Round(f*10000.0) / 10000
}

func Test(t *testing.T) {

	res, ans := round5(myPow(2.0, 10)), 1024.0
	if res != ans {
		t.Errorf("Expected %v, but got %v", ans, res)
	}
	res, ans = round5(myPow(2.1, 3)), 9.261
	if res != ans {
		t.Errorf("Expected %v, but got %v", ans, res)
	}
}

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
