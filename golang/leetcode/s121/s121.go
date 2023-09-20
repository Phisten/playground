package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> 5 \n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v -> 2 \n", maxProfit([]int{2, 4, 1}))
	fmt.Printf("%v -> 5 \n", maxProfit([]int{2, 1, 4, 2, 6}))
}

// 先搜最小 然後往右搜最大
func maxProfit(prices []int) int {
	min := 10001
	max := -1
	ans := 0

	for _, value := range prices {
		if value < min {
			if max-min > ans {
				ans = max - min
			}
			min = value
			max = value
		} else if value > max {
			max = value
		}
	}

	if max-min > ans {
		ans = max - min
	}

	return ans
}
