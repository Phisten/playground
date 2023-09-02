package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v -> 4 \n", countKSubsequencesWithMaxBeauty("bcca", 2))
	fmt.Printf("%v -> 2 \n", countKSubsequencesWithMaxBeauty("abbcd", 4))
}

// mod 1000000007
// 1.計算每個字元的數量進hashTable 2.過濾出不連續重複的字元/使用滑動窗？
// 輸出有幾組子字串的分數總和都是最高的
// s1:DP X
// s2:先過濾出最高分的k個字母(同分時會超過k個)
// 同分的相加
// 拿這k個數字全部乘起來應該就是答案
// *來不及完成 , 算法也有問題
func countKSubsequencesWithMaxBeauty(s string, k int) int {
	fmt.Printf("s=%v, k=%v\n", s, k)
	scoreTable := map[byte]int{}
	for i := 0; i < len(s); i++ {
		scoreTable[s[i]]++
	}
	fmt.Println(scoreTable)

	scoreList := []int{}
	for i := 0; i < len(s); i++ {
		if i == 0 || (s[i] != s[i-1]) {
			scoreList = append(scoreList, scoreTable[s[i]])
		}
	}
	sort.Ints(scoreList)

	fmt.Println(scoreList)

	ansPickList := make([]int, k)
	lastK := k
	for i := len(scoreList) - 1; i >= 0; i-- {
		ansPickList[lastK-1] += scoreList[i]
		if i == len(scoreList)-1 || scoreList[i] != scoreList[i+1] {
			lastK--
		}
		if lastK <= 0 {
			break
		}
	}
	fmt.Println(ansPickList)

	// 前k高分的分數乘起來(排列組合數)
	ans := 1
	for i := 0; i < k; i++ {
		if ansPickList[i] > 0 {
			ans *= ansPickList[i]
		}
	}

	return ans
}
