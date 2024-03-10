package leetcode

import (
	"reflect"
	"testing"
)

func Test_s349(t *testing.T) {
	type Pair struct {
		got      []int
		expected []int
		note     string
	}

	Func := intersection
	paris := []Pair{
		{Func([]int{1, 2, 2, 1}, []int{2, 2}), []int{2}, ""},
		{Func([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{9, 4}, ""},
	}

	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	ht := make(map[int]int)
	ans := []int{}

	for i := 0; i < len(nums1); i++ {
		ht[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		if ht[nums2[i]] == 1 {
			ht[nums2[i]] = 2
			ans = append(ans, nums2[i])
		}
	}

	return ans
}
