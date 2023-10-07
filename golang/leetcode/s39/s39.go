package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Printf("%v -> [[2,2,3],[7]] \n", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Printf("%v -> \n[[7,7,2,2],[7,3,3,3,2],[7,3,2,2,2,2],[3,3,3,3,3,3],[3,3,3,3,2,2,2],[3,3,2,2,2,2,2,2],[2,2,2,2,2,2,2,2,2]] \n", combinationSum([]int{7, 3, 2}, 18))
}

// 注意到target的值不大, 可能可以由1開始動態規劃往上算到target
// - 輸出不需排序,candidates可以先排序
// - dp儲存target扣到剩多少時有哪些解法
// - 處理到重複解
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)

	dp := make([][][]int, target+1)

	ans := [][]int{}

	if n == 0 {
		return ans
	}
	for t := candidates[0]; t <= target; t++ {
		dp[t] = [][]int{}
		for _, v := range candidates {
			if v == t {
				dp[t] = append(dp[t], []int{v})
			} else if v < t {
				subValue := t - v

				for _, vj := range dp[subValue] {
					if vj[len(vj)-1] <= v {
						newArr := make([]int, len(vj))
						copy(newArr, vj)
						newArr = append(newArr, v)
						dp[t] = append(dp[t], newArr)
					}
				}
			}
		}
	}

	// test output
	// for i, v := range dp {
	// 	if len(v) > 0 {
	// 		fmt.Println(i, ":", v)
	// 	}
	// }

	return dp[target]
}
