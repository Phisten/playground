package leetcode

import (
	"fmt"
	"testing"
)

// Kadane
func Test_s53(t *testing.T) {
	fmt.Printf("Test_s50\n")

	type Pair struct {
		got      int
		expected int
		note     string
	}

	Func := maxSubArray
	paris := []Pair{
		{Func([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6.0, "[-2, 1, -3, 4, -1, 2, 1, -5, 4]"},
		{Func([]int{1}), 1, "1"},
		{Func([]int{5, 4, -1, 7, 8}), 23, "23"},
	}

	for _, v := range paris {
		if v.expected != v.got {
			t.Errorf("Expected %v, but got %v, note:%v", v.expected, v.got, v.note)
		}
	}
}

func maxSubArray(nums []int) int {
	largestSum := nums[0]
	lastLargestSum := nums[0]
	l := len(nums)

	for i := 1; i < l; i++ {
		curLargestSum := max(nums[i], nums[i]+lastLargestSum)
		largestSum = max(largestSum, curLargestSum)
		lastLargestSum = curLargestSum
	}

	return largestSum
}
