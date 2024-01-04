package leetcode

import (
	"reflect"
	"testing"
)

func Test_s973(t *testing.T) {
	type Pair struct {
		got      [][]int
		expected [][]int
		note     string
	}

	Func := kClosest
	paris := []Pair{
		// {Func([][]int{{6, 10}, {-3, 3}, {-2, 5}, {0, 2}}, 3), [][]int{{0, 2}, {-3, 3}, {-2, 5}}, ""},
		{Func([][]int{{1, 3}, {-2, 2}}, 1), [][]int{{-2, 2}}, ""},
		{Func([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2), [][]int{{3, 3}, {-2, 4}}, ""},
	}

	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

// sort or heap

// Time Limit Exceeded
func kClosest(points [][]int, k int) [][]int {
	l := len(points)
	minK := min(l, k)

	closestK := make([][]int, minK)
	for i := 0; i < l; i++ {
		curDist := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		points[i] = append(points[i], curDist)
	}

	for j := 0; j < minK; j++ {
		for i := j + 1; i < l; i++ {
			if points[j][2] > points[i][2] {
				points[j], points[i] = points[i], points[j]
			}
		}
		closestK[j] = points[j][:2]
	}

	return closestK
}
