package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> 13 \n", sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum1 := 0

	for i := 0; i < len(nums); i++ {
		n := i
		iBit := 0
		for ; n > 0; n /= 2 {
			if n%2 == 1 {
				iBit++
			}
		}
		if iBit == k {
			sum1 += nums[i]
		}
	}
	return sum1
}
