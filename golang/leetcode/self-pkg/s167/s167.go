// 167. Two Sum II - Input Array Is Sorted
package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func binarySearch(nums []int, start int, end int, target int) int {

	for start <= end {
		mid := (end-start)/2 + start
		midValue := nums[mid]

		if midValue == target {
			return mid
		} else if midValue < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func twoSum(numbers []int, target int) []int {

	numbersLen := len(numbers)

	leftIdx := 0
	endIdx := numbersLen - 1

	for leftIdx < endIdx {
		if numbers[leftIdx]+numbers[endIdx] >= target {

			tryFindIdx := binarySearch(numbers, leftIdx+1, endIdx, target-numbers[leftIdx])
			if tryFindIdx >= 0 {
				return []int{leftIdx + 1, tryFindIdx + 1}
			}
		}
		leftIdx += 1
	}
	return []int{}
}
