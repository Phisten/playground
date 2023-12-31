package leetcode

import (
	"fmt"
	"testing"
)

func Test_s733(t *testing.T) {
	fmt.Printf("Test_s733\n")

	type Pair struct {
		got      [][]int
		expected [][]int
		note     string
	}

	equal := func(v Pair) bool {
		for i := 0; i < len(v.expected); i++ {
			for j := 0; j < len(v.expected[i]); j++ {
				if v.expected[i][j] != v.got[i][j] {
					return false
				}
			}
		}
		return true
		// return v.expected != v.got
	}

	paris := []Pair{
		{floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, "1"},
		{floodFill([][]int{{0, 0, 0}}, 0, 0, 0), [][]int{{0, 0, 0}}, "2"},
	}

	for _, v := range paris {
		if !equal(v) {
			t.Errorf("Expected %v, but got %v, note:%v", v.expected, v.got, v.note)
		}
	}
}

// 數值範圍不大 看起來可以遞迴
// 1 <= m, n <= 50, image= m*n
// 影像處理注意邊界
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m := len(image)
	if m == 0 {
		return image
	}
	n := len(image[0])

	rc(image[sr][sc], image, sr, sc, color, m, n)

	return image
}

func rc(fromColor int, image [][]int, sr int, sc int, color int, m int, n int) {
	if sr < 0 || sc < 0 || sr >= m || sc >= n ||
		image[sr][sc] == color {
		return
	}

	if image[sr][sc] == fromColor {
		image[sr][sc] = color
		rc(fromColor, image, sr-1, sc, color, m, n)
		rc(fromColor, image, sr+1, sc, color, m, n)
		rc(fromColor, image, sr, sc-1, color, m, n)
		rc(fromColor, image, sr, sc+1, color, m, n)
	}

}
