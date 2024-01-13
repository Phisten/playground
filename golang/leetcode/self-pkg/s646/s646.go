package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v -> 4 \n", findLongestChain([][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}}))
	fmt.Printf("%v -> 2 \n", findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Printf("%v -> 3 \n", findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))

}

func findLongestChain(pairs [][]int) int {

	len1 := len(pairs)

	sort.Slice(pairs,
		func(i, j int) bool {
			return pairs[i][0] < pairs[j][0]
		})

	fmt.Println(pairs)

	dp := make([]int, len1)
	max := 1
	for i := 0; i < len1; i++ {
		dp[i] = 1
	}
	for i := 1; i < len1; i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = dp[j] + 1
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
		fmt.Printf("%v ", dp[i])
	}
	fmt.Printf("\n")

	return max
}
