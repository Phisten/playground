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

func twoSum(numbers []int, target int) []int {
	numbersLen := len(numbers)

	leftIdx := 0

	for leftIdx < numbersLen-1 {
		if numbers[leftIdx]+numbers[numbersLen-1] >= target {

			for rightIdx := leftIdx + 1; rightIdx < numbersLen; rightIdx += 1 {
				if numbers[leftIdx]+numbers[rightIdx] == target {
					return []int{leftIdx + 1, rightIdx + 1}
				}
			}
		}
		leftIdx += 1
	}
	return []int{}
}
