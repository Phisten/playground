package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> 4 \n", search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Printf("%v -> -1 \n", search([]int{}, 9))
}

func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	last := -1

	for {
		mid := (left + right) / 2
		if mid == last || left >= n || right < 0 {
			break
		}

		cur := nums[mid]
		if target == cur {
			return mid
		} else if target > cur {
			left = mid + 1
		} else {
			right = mid - 1
		}

		last = mid
	}

	return -1
}
