// https://leetcode.com/problems/3sum/description/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("3sum start")
	// solution1: 三層迴圈暴力遍歷
	// s2: 快取2sum三角矩陣 一層迴圈列舉相反數
	// s3: 排序後迴圈遍歷, 以雙指針搜尋總和是否符合範圍
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, 0, 0, -1}))
	fmt.Println("3sum end")
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	target := 0

	numsLen := len(nums)

	ans := [][]int{}
	appendLog := make(map[[3]int]bool)

	for index, value := range nums {
		left := index + 1
		right := numsLen - 1

		// 數字不足 或 首數已過大,後續無解
		if index > numsLen-3 || value+nums[left]+nums[left+1] > target {
			break
		}

		// 首數過小
		if value+nums[numsLen-1]+nums[numsLen-2] < target {
			continue
		}

		for index < left && left < right {
			sum := value + nums[left] + nums[right]
			if sum > target {
				right -= 1
			} else if sum < target {
				left += 1
			} else {
				newAns := [3]int{value, nums[left], nums[right]}
				if !appendLog[newAns] {
					ans = append(ans, []int{value, nums[left], nums[right]})
					appendLog[newAns] = true
				}
				left += 1
				right -= 1
			}
		}

	}

	return ans
}
