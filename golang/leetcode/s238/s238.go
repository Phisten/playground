package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> [24,12,8,6]  \n", productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {

	product := 1
	for _, v := range nums {
		product *= v
	}
	for i := range nums {
		if nums[i] != 0 {
			nums[i] = product / nums[i]
		} else {
			nums[i] = product
		}
	}
	return nums
}
