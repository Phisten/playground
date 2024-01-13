package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v -> 2 \n", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

// 題目要求空間O(1)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
