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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)

	twoSum := []int{}
	twoSumMap := make(map[int]([][2]int))

	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			curSum := nums[i] + nums[j]
			twoSum = append(twoSum, curSum)
			twoSumMap[curSum] = append(twoSumMap[curSum], [2]int{nums[i], nums[j]})
		}
	}

	ans := [][]int{}
	ansLogMap := make(map[[4]int]bool)

	twoSumLen := len(twoSum)

	for i := 0; i < twoSumLen-1; i++ {
		for j := i + 1; j < twoSumLen; j++ {
			if twoSum[i]+twoSum[j] == target {
				for _, left := range twoSumMap[twoSum[i]] {
					for _, right := range twoSumMap[twoSum[j]] {
						orderCurAns := []int{left[0], left[1], right[0], right[1]}
						sort.Ints(orderCurAns)
						orderCurAns4 := [4]int{orderCurAns[0], orderCurAns[1], orderCurAns[2], orderCurAns[3]}

						if !ansLogMap[orderCurAns4] {
							ansLogMap[orderCurAns4] = true
							ans = append(ans, orderCurAns)
						}
					}
				}
			}
			if j < numsLen-1 && twoSum[j] == twoSum[j+1] {
				j++
			}
		}
		if twoSum[i] == twoSum[i+1] {
			i++
		}
	}

	fmt.Println(twoSum)

	return ans
}
