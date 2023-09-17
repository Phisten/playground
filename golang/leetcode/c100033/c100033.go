package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("%v -> 2 \n", maxNumberOfAlloys(2, 3, 10, [][]int{{2, 1}, {1, 2}, {1, 1}}, []int{1, 1}, []int{5, 5}))
	// fmt.Printf("%v -> 5 \n", maxNumberOfAlloys(3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 100}, []int{1, 2, 3}))
	// fmt.Printf("%v -> \n", maxNumberOfAlloys(1, 1, 2, [][]int{{2}}, []int{1}, []int{1}))
	//
	// fmt.Printf("%v -> 1\n", maxNumberOfAlloys(4, 9, 55, [][]int{{8, 3, 4, 2}, {3, 9, 5, 5}, {1, 7, 9, 8}, {7, 6, 5, 1}, {4, 6, 9, 4}, {6, 8, 7, 1}, {5, 10, 3, 4}, {10, 1, 2, 4}, {10, 3, 7, 2}}, []int{9, 1, 10, 0}, []int{3, 4, 9, 9}))
	fmt.Printf("%v -> 3492\n", maxNumberOfAlloys(12, 17, 62283032, [][]int{{2, 3, 41, 74, 33, 10, 83, 73, 35, 20, 20, 98}, {75, 93, 91, 27, 76, 67, 86, 45, 90, 92, 20, 67}, {26, 54, 67, 39, 65, 75, 24, 21, 35, 94, 87, 46}, {4, 80, 58, 44, 66, 70, 22, 61, 20, 4, 1, 2}, {89, 75, 42, 92, 58, 84, 64, 77, 57, 22, 7, 58}, {99, 47, 20, 81, 88, 61, 54, 29, 29, 79, 19, 37}, {93, 23, 82, 85, 57, 29, 70, 73, 12, 73, 87, 25}, {11, 95, 5, 65, 69, 22, 94, 25, 43, 94, 78, 82}, {42, 38, 10, 42, 96, 64, 20, 32, 41, 73, 24, 59}, {75, 80, 65, 99, 32, 81, 20, 12, 70, 39, 3, 47}, {93, 95, 5, 36, 91, 71, 31, 40, 42, 88, 42, 24}, {23, 7, 99, 96, 52, 65, 29, 17, 81, 12, 38, 35}, {64, 82, 93, 73, 81, 93, 98, 98, 20, 49, 49, 10}, {91, 100, 61, 13, 7, 60, 78, 52, 23, 2, 97, 61}, {38, 90, 79, 19, 88, 94, 98, 10, 60, 67, 42, 56}, {30, 36, 31, 24, 48, 13, 41, 21, 19, 54, 65, 6}, {27, 42, 37, 77, 54, 75, 41, 26, 5, 5, 26, 64}}, []int{97, 49, 16, 58, 20, 16, 50, 41, 82, 14, 35, 76}, []int{2, 24, 71, 81, 3, 41, 79, 24, 83, 32, 74, 50}))
}

// n合金類型 k機器 計算效率最高的機器 返回最大可製造數量 不可超預算
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {

	failedCount := 0
	buildCostK := make([]int, k)
	stockByK := make([][]int, k)
	for i := 0; i < k; i++ {
		stockByK[i] = stock
	}

	ansCount := 0
	for tryCount := 1; failedCount < k; tryCount++ {
		for mcIdx := 0; mcIdx < k; mcIdx++ {
			if buildCostK[mcIdx] <= budget {
				curCost := buildCostK[mcIdx]

				for metIdx := 0; metIdx < n; metIdx++ {
					need := tryCount * composition[mcIdx][metIdx]
					if stock[metIdx] < need {
						newCost := 0
						lastNeed := need - stock[metIdx]
						if lastNeed < composition[mcIdx][metIdx] {
							newCost = lastNeed * cost[metIdx]
						} else {
							newCost = composition[mcIdx][metIdx] * cost[metIdx]
						}
						curCost += newCost
					}
				}
				buildCostK[mcIdx] = curCost
				if curCost <= budget {
					ansCount = tryCount
				} else {
					failedCount++
				}
			}
		}
	}

	return ansCount
}
