package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("%v -> 2 \n", countPairs([][]int{{1, 2}, {4, 2}, {1, 3}, {5, 2}}, 5))
	fmt.Printf("%v -> 5 \n", countPairs([][]int{{27, 94}, {61, 68}, {47, 0}, {100, 4}, {127, 89}, {61, 103}, {26, 4}, {51, 54}, {91, 26}, {98, 23}, {80, 74}, {19, 93}}, 95))
}

func countPairs(coordinates [][]int, k int) int {
	ans := 0

	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {

			left := coordinates[i][0] ^ coordinates[j][0]
			right := coordinates[i][1] ^ coordinates[j][1]
			// fmt.Printf("pointI:%v,pointJ:%v ,left:%v right:%v \n", i, j, left, right)
			if left+right == k {
				ans++
			}
		}
	}

	return ans
}
