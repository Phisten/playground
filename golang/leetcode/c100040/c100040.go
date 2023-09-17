package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v -> 3 \n", countWays([]int{6, 0, 3, 3, 6, 7, 2, 7}))
}

func countWays(nums []int) int {
	ans := 0
	lenN := len(nums)

	sort.Ints(nums)
	valueMap := map[int]int{}
	for i := 0; i < lenN; i++ {
		valueMap[nums[i]] += 1
	}
	for selected := 0; selected <= lenN; selected++ {
		if valueMap[selected] == 0 {

			lessLen := sort.Search(len(nums), func(i int) bool {
				return nums[i] > selected
			})
			if lessLen == -1 {
				lessLen = len(nums)
			}

			if lessLen == selected {
				ans++
			}

		}
	}
	return ans
}
