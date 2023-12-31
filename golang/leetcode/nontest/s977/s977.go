package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v -> [0,1,9,16,100] \n", sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	left := 0
	right := n - 1
	ans := make([]int, n)
	for i := right; i >= 0; i-- {
		if left == n-1 {
			rightSq := nums[right] * nums[right]
			ans[i] = rightSq
			right--
		} else if right == 0 {
			leftSq := nums[left] * nums[left]
			ans[i] = leftSq
			left++
		} else {

			leftSq := nums[left] * nums[left]
			rightSq := nums[right] * nums[right]
			if leftSq > rightSq {
				ans[i] = leftSq
				left++
			} else {
				ans[i] = rightSq
				right--
			}
		}

	}
	return ans
}

// O(N
func sortedSquares_0(nums []int) []int {

	for i, v := range nums {
		if v >= 0 {
			break
		} else {
			nums[i] = -v
		}
	}

	n := len(nums)
	left := 0
	right := n - 1
	ans := make([]int, n)
	for i := right; i >= 0; i-- {
		if left == n-1 {
			ans[i] = nums[right] * nums[right]
			right--
		} else if right == 0 {
			ans[i] = nums[left] * nums[left]
			left++
		} else {
			if nums[left] > nums[right] {
				ans[i] = nums[left] * nums[left]
				left++
			} else {
				ans[i] = nums[right] * nums[right]
				right--
			}
		}
	}
	return ans
}

// O(NlogN)
func sortedSquares_2(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)

	return nums
}
