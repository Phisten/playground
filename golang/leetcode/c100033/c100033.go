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
	fmt.Printf("%v -> 1\n", maxNumberOfAlloys(4, 1, 55, [][]int{{8, 3, 4, 2}, {7, 6, 5, 1}}, []int{9, 1, 10, 0}, []int{3, 4, 9, 9}))
	//          																								16,6,8,4 - 9,1,10,0 = 7,5,0,4 * 3499=21+20+0+36 > 55    1060   8+18=26
}

// n合金類型 k機器 計算效率最高的機器 返回最大可製造數量 不可超預算 OLE
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
						fmt.Printf("mcIdx: %v, metIdx %v: newCost %v\n", mcIdx, metIdx, newCost)
						curCost += newCost
					}
				}
				buildCostK[mcIdx] = curCost
				if curCost <= budget {
					fmt.Printf("index: %v\n", mcIdx)
					ansCount = tryCount
				} else {
					failedCount++
				}
			}
		}
	}

	return ansCount
}
