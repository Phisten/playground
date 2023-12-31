package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("%v ->[[1,2],[3,10],[12,16]] \n", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	// fmt.Printf("%v -> [[0 1] [2 3]] \n", insert([][]int{{2, 3}}, []int{0, 1}))
	// fmt.Printf("%v -> [[1,2],[3,10],[12,16]] \n", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Printf("%v -> [[1,2],[3,3],[4,5]] \n", insert([][]int{{1, 2}, {4, 5}}, []int{3, 3}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}

	lenN := len(intervals)
	mergeLeft := newInterval[0]
	margeRight := newInterval[1]

	needInsert := false

	// 插入最左
	if lenN > 0 && newInterval[1] < intervals[0][0] {
		ans = append(ans, newInterval)
	} else if lenN == 0 {
		ans = append(ans, newInterval)
		return ans
	}

	// 混合或穿插在其中
	for i := 0; i < len(intervals); i++ {
		rangeL := intervals[i][0]
		rangeR := intervals[i][1]
		if (rangeL <= newInterval[0] && newInterval[0] <= rangeR) ||
			(rangeL <= newInterval[1] && newInterval[1] <= rangeR) ||
			(newInterval[0] < rangeL && rangeR < newInterval[1]) {
			if rangeL < mergeLeft {
				mergeLeft = rangeL
			}
			if margeRight < rangeR {
				margeRight = rangeR
			}
			needInsert = true
		} else if i > 0 && (intervals[i-1][1] < newInterval[0] && newInterval[1] < rangeL) {
			ans = append(ans, []int{mergeLeft, margeRight})
			ans = append(ans, intervals[i])
		} else {
			if needInsert {
				ans = append(ans, []int{mergeLeft, margeRight})
				needInsert = false
			}
			ans = append(ans, intervals[i])
		}
	}

	// 最後一筆被合併
	if needInsert {
		ans = append(ans, []int{mergeLeft, margeRight})
		needInsert = false
	}

	// 插入最右
	if lenN > 0 && intervals[lenN-1][1] < newInterval[0] {
		ans = append(ans, newInterval)
	}

	return ans
}
