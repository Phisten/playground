package leetcode

import (
	"reflect"
	"testing"
)

func Test_c713(t *testing.T) {
	type Pair struct {
		got      int
		expected int
		note     string
	}

	Func := numSubarrayProductLessThanK
	paris := []Pair{
		{Func([]int{10, 5, 2, 6}, 100), 8, ""},
		{Func([]int{1, 2, 3}, 0), 0, ""},
	}

	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	l := len(nums)

	ans := 0

	left, right := 0, 0
	product := 1
	for right < l {
		valueR := nums[right]

		product *= valueR

		if product < k {
			ans += right - left + 1
		}

		for product >= k && left <= right {
			valueL := nums[left]
			left++

			product /= valueL

			if product < k {
				ans += right - left + 1
			}
		}

		right++
	}

	return ans
}
