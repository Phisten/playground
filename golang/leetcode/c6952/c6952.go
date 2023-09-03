package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> 3 \n", countInterestingSubarrays([]int{3, 2, 4}, 2, 1))
	fmt.Printf("%v -> 2 \n", countInterestingSubarrays([]int{3, 1, 9, 6}, 3, 0))
	fmt.Printf("%v -> 5 \n", countInterestingSubarrays([]int{11, 12, 21, 31}, 10, 1))

}

// 需要先快取每個i mod 是否是k
// *超時
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	numL := len(nums)
	modSucc := map[int]int{}

	for _, num := range nums {
		if num%modulo == k {
			modSucc[num] = 1
		}
	}

	var ans int64 = 0
	for arrLen := 1; arrLen <= numL; arrLen++ {

		for i := 0; i <= numL-arrLen; i++ {
			cnt := 0
			for j := i; j < i+arrLen; j++ {
				cnt += modSucc[nums[j]]
			}
			if cnt%modulo == k {
				ans++
			}
		}
	}

	return ans

}
