package leetcode

import (
	"reflect"
	"testing"
)

// Kadane
func Test_s2540(t *testing.T) {
	type Pair struct {
		got      int
		expected int
		note     string
	}

	Func := getCommon
	paris := []Pair{
		{Func([]int{1, 2, 3}, []int{2, 4}), 2, ""},
		{Func([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}), 2, ""},
	}

	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return -1
}
