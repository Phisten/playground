package leetcode

import "testing"

func Test_s50(t *testing.T) {

}

func subsets(nums []int) [][]int {
	l := len(nums)
	var ans = [][]int{}

	for i := 0; i <= l; i++ {
		curAns := []int{}
		for j := 0; j < l-i; j++ {
			curAns = append(curAns)
		}

		ans = append(ans, curAns)
	}

	return ans
}

func recursion(nums []int, start int, arrLen int) []int {

}
