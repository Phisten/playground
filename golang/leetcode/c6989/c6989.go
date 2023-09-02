package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> 18 \n", maxSum([]int{2, 6, 7, 3, 1, 7}, 3, 4))
	fmt.Printf("%v -> 23 \n", maxSum([]int{5, 9, 9, 2, 4, 5, 4}, 1, 3))
	fmt.Printf("%v -> 0 \n", maxSum([]int{1, 2, 1, 2, 1, 2, 1}, 3, 3))
	fmt.Printf("%v -> 4 \n", maxSum([]int{1, 1, 1, 3}, 2, 2))
}

// 主迴圈slidingWindow掃描所有組合, 子迴圈檢查不重複字元是否多餘m, 紀錄最大sum
// *超時
func maxSum_TLE(nums []int, m int, k int) int64 {
	var ans int64 = 0

	for i := 0; i <= len(nums)-k; i++ {

		hash1 := map[int]bool{}
		curM := m
		var curSum int64 = 0
		for j := i; j < i+k; j++ {
			if !hash1[nums[j]] {
				hash1[nums[j]] = true
				curM--
			}
			curSum += int64(nums[j])
		}

		if curM <= 0 && curSum > ans {
			ans = curSum
		}
	}

	return ans
}

// 預先跑好一組滑動窗  接下來只輸入變化
func maxSum(nums []int, m int, k int) int64 {
	var ans int64 = 0

	// init
	hash1 := map[int]int{}
	curM := m
	var curSum int64 = 0
	for i := 0; i < k; i++ {
		if hash1[nums[i]] == 0 {
			curM--
		}
		hash1[nums[i]] += 1
		curSum += int64(nums[i])
	}
	if curM <= 0 && curSum > ans {
		ans = curSum
	}

	// sliding
	for i := k; i < len(nums); i++ {
		// remove old
		hash1[nums[i-k]] -= 1
		if hash1[nums[i-k]] == 0 {
			curM++
		}
		curSum -= int64(nums[i-k])

		// add new
		if hash1[nums[i]] == 0 {
			curM--
		}
		hash1[nums[i]] += 1
		curSum += int64(nums[i])

		if curM <= 0 && curSum > ans {
			ans = curSum
		}
	}

	return ans
}
