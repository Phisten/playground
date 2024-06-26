package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> [24,12,8,6]  \n", productExceptSelf([]int{1, 2, 3, 4}))
}

// 首先想到將所有數相乘, 然後掃描nums時拿總乘積除以nums[i]得到每一位的答案
//   然後才發現題目要求不能用除法運算子

// s2:將nums[i]左右兩側的值分別乘積並乘進目標ans對應的位置
func productExceptSelf(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)

	left := 1
	right := 1

	for i, v := range nums {
		ans[i] = left
		left *= v
	}

	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}

	return ans
}

// s1 Timeout:O(N^2)乘法掃描方法不符合題目速度要求O(N)
func productExceptSelf_(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)

	for i := range ans {
		ans[i] = 1

		for j, q := range nums {
			if i != j {
				ans[i] *= q
			}
		}
	}

	return ans
}
