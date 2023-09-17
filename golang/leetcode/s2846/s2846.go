package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Printf("%v -> [0,0,1,3] \n", minOperationsQueries(7, [][]int{
	// 	{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 4, 2}, {4, 5, 2}, {5, 6, 2},
	// }, [][]int{
	// 	{0, 3}, {3, 6}, {2, 6}, {0, 6},
	// }))

	fmt.Printf("%v -> ? \n", minOperationsQueries(8, [][]int{
		{1, 2, 6}, {1, 3, 4}, {2, 4, 6}, {2, 5, 3}, {3, 6, 6}, {3, 0, 8}, {7, 0, 2},
	}, [][]int{
		{4, 6}, {0, 4}, {6, 5}, {7, 4},
	}))

}

// TLE
func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	qLen := len(queries)
	ans := []int{}

	// edges init 排序

	for i := 0; i < qLen; i++ {
		curPoint := queries[i][0]
		endPoint := queries[i][1]
		if curPoint > endPoint {
			curPoint = endPoint
			endPoint = queries[i][0]
		}

		hash1 := map[int]int{}
		list1 := []int{}
		for curPoint != endPoint {
			cost := edges[curPoint][2]
			if hash1[cost] == 0 {
				list1 = append(list1, cost)
			}
			hash1[cost] += 1
			curPoint = edges[curPoint][1]
		}

		for idx, value := range list1 {
			list1[idx] = hash1[value]
		}
		sort.Ints(list1)
		sum := 0
		for i := 0; i < len(list1)-1; i++ {
			sum += list1[i]
		}
		ans = append(ans, sum)
	}

	return ans
}
