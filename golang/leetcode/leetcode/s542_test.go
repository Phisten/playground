package leetcode

import (
	"reflect"
	"testing"
)

// Kadane
func Test_s542(t *testing.T) {
	type Pair struct {
		got      [][]int
		expected [][]int
		note     string
	}

	Func := updateMatrix
	paris := []Pair{
		{Func([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}), [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}, ""},
		{Func([][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}), [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}}, ""},
		{Func([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}), [][]int{{4, 3, 2}, {3, 2, 1}, {2, 1, 0}}, ""},
		{Func([][]int{{0, 0, 0}}), [][]int{{0, 0, 0}}, ""},
	}

	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	opt := make([][]int, m)

	// scan 0
	for i := 0; i < m; i++ {
		hasTop := i > 0
		opt[i] = make([]int, n)

		for j := 0; j < n; j++ {
			hasLeft := j > 0

			if mat[i][j] == 0 {
				opt[i][j] = 0
			} else {
				curDist := 10001
				if hasLeft {
					curDist = min(curDist, opt[i][j-1]+1)
				}
				if hasTop {
					curDist = min(curDist, opt[i-1][j]+1)
				}
				opt[i][j] = curDist
			}
		}
	}

	//reverse scan
	bound_m := m - 1
	bound_n := n - 1
	for i := 0; i < m; i++ {
		hasBottom := i > 0

		for j := 0; j < n; j++ {
			hasRight := j > 0

			if mat[bound_m-i][bound_n-j] == 0 {
				opt[bound_m-i][bound_n-j] = 0
			} else {
				curDist := opt[bound_m-i][bound_n-j]
				if hasRight {
					curDist = min(curDist, opt[bound_m-i][bound_n-j+1]+1)
				}
				if hasBottom {
					curDist = min(curDist, opt[bound_m-i+1][bound_n-j]+1)
				}
				opt[bound_m-i][bound_n-j] = curDist
			}
		}
	}

	return opt
}
