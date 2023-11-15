package leetcode

import (
	"testing"
)

func Test(t *testing.T) {

	res, ans := myPow(2.0, 10), 1024.0
	if res != ans {
		t.Errorf("Expected %v, but got %v", ans, res)
	}
	res, ans = myPow(2.1, 3), 9.261
	if res != ans {
		t.Errorf("Expected %v, but got %v", ans, res)
	}
}

func myPow(x float64, n int) float64 {
	return 0
}
