package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v==4 \n", countKDifference([]int{1, 1, 2, 2}, 1))
	fmt.Println(countKDifference([]int{10, 2, 10, 9, 1, 6, 8, 9, 2, 8}, 5))
	fmt.Println(countKDifference([]int{1, 2, 3}, 1))
	fmt.Println(countKDifference([]int{7, 7, 8, 3, 1, 2, 7, 2, 9, 5}, 6))
}

func countKDifference(nums []int, k int) int {
	ans := 0

	sort.Ints(nums)

	dict := make(map[int]int, 101)

	for _, value := range nums {
		dict[value]++
	}

	lastNum := nums[len(nums)-1]
	for _, value := range nums {
		targetSubtract := k + value
		if targetSubtract > lastNum {
			break
		}

		if _, exists := dict[targetSubtract]; exists {
			ans += dict[targetSubtract]
		}
	}

	return ans
}
