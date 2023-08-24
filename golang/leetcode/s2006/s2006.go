package main

import (
	"fmt"
)

func main() {
	fmt.Println(countKDifference([]int{1, 2, 2, 1}, 1))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func countKDifference(nums []int, k int) int {
	ans := 0

	for i := 0; i < len(nums)-1; i++ {
		targetA := abs(nums[i] - k)
		for j := i + 1; j < len(nums); j++ {
			if targetA == nums[j] {
				ans++
			}
		}
	}

	return ans
}
