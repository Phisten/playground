// https://leetcode.com/problems/4sum/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	// fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println("4sum end")
}

func threeSum(nums []int, start int, target int) [][]int {
	numsLen := len(nums)

	ans := [][]int{}
	appendLog := make(map[[3]int]bool)

	for i := start; i < numsLen; i++ {
		index := i
		value := nums[index]

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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)

	ans := [][]int{}
	appendLog := make(map[[4]int]bool)

	for i := 0; i < numsLen; i++ {
		for _, item := range threeSum(nums, i+1, target-nums[i]) {
			log := [4]int{nums[i], item[0], item[1], item[2]}
			if !appendLog[log] {
				appendLog[log] = true
				ans = append(ans, []int{nums[i], item[0], item[1], item[2]})
			}
		}
	}

	return ans
}
