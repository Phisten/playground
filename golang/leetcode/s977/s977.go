package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v -> [0,1,9,16,100] \n", sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)

	return nums
}
