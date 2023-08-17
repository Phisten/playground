package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
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
