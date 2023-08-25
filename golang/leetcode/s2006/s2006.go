package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countKDifference([]int{1, 1, 2, 2}, 1))
	fmt.Println(countKDifference([]int{10, 2, 10, 9, 1, 6, 8, 9, 2, 8}, 5))
	fmt.Println(countKDifference([]int{1, 2, 3}, 1))
	fmt.Println(countKDifference([]int{7, 7, 8, 3, 1, 2, 7, 2, 9, 5}, 6))
}

func countKDifference(nums []int, k int) int {
	ans := 0

	sort.Ints(nums)
	fmt.Printf("src: %v , target: %v \n", nums, k)

	dict := make(map[int]int, 101)

	// scan hashmap
	for _, value := range nums {
		dict[value]++
	}

	for i := 0; i < len(nums); i++ {

		if nums[i] <= k/2.0 {
			// 可否同數相加
			if nums[i] == k/2 && k/2*2 == k {
				if value, exists := dict[nums[i]]; exists && value >= 2 {
					dict[nums[i]]--
					ans += dict[nums[i]]
				}
			} else {
				// 找相加
				targetAdd := k - nums[i]
				if value, exists := dict[targetAdd]; exists && value > 0 {
					ans += dict[targetAdd]
				}
			}
		}

		// 找相減
		targetSubtract := k + nums[i]
		if targetSubtract > nums[len(nums)-1] {
			break
		}

		if value, exists := dict[targetSubtract]; exists && value > 0 {
			ans += dict[targetSubtract]
		}

	}

	return ans
}

// [7,7,8,3,1,2,7,2,9,5]
// [1 2 2 3 5 7 7 7 8 9]
/*
 7-1 x2
 8-2 x2
 1-7 x1 =-6
 3-9 x1 =-6
 1+5 x1 =6




*/
