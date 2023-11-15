package leetcode

import (
	"math"
	"testing"
)

func Test_s50(t *testing.T) {
	round5 := func(f float64) float64 {
		return math.Round(f*10000.0) / 10000
	}

	if got, exp := round5(myPow(2.0, 10)), 1024.0; got != exp {
		t.Errorf("Expected %v, but got %v", exp, got)
	}
	if got, exp := round5(myPow(2.1, 3)), 9.261; got != exp {
		t.Errorf("Expected %v, but got %v", exp, got)
	}

	if got, exp := round5(myPow_sdk(2.1, 3)), 9.261; got != exp {
		t.Errorf("myPow_sdk Expected %v, but got %v", exp, got)
	}
}

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func myPow_sdk(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
