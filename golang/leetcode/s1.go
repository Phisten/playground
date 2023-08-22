package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := numToIndex[complement]; ok {
			return []int{index, i}
		}

		numToIndex[num] = i
	}

	return nil
}
