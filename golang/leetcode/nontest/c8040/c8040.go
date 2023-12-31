package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Printf("%v -> 2 \n", minimumOperations("5047"))
	//                                          0123
	fmt.Printf("%v -> 2 \n", minimumOperations("2245047"))
	fmt.Printf("%v -> 3 \n", minimumOperations("2908305"))
	fmt.Printf("%v -> 1 \n", minimumOperations("10"))
}

// 最大值為lenNum, 若num有0 可以-1
// s1:num字串長度很大 超過int64
func minimumOperations_x(num string) int {
	num1, _ := new(big.Int).SetString("1234567890123456789012345678901234567890", 10)
	num25, _ := new(big.Int).SetString("25", 10)
	a := new(big.Int).Mod(num1, num25)
	ans, _ := strconv.Atoi(a.String())
	return ans
}

// s2: 尾數必須是 0 25 50 75 00
// 可以從後指針開始找0或5結尾  把後面都去掉
// 然後再往前找一個數字必須是0257 找到就是答案
// 都沒有就要刪到剩下0 或是刪全部數字變成0
func minimumOperations(num string) int {
	numLen := len(num)
	ans := numLen
	for _, char := range num {
		if char == '0' {
			ans = numLen - 1
		}
	}

	for i := numLen - 1; i > 0; i-- {
		rightNum := num[i]
		if rightNum == '0' || rightNum == '5' {
			for j := i - 1; j >= 0; j-- {
				preNum := num[j]
				if (preNum == '0' && rightNum == '0') || (preNum == '2' && rightNum == '5') || (preNum == '7' && rightNum == '5') || (preNum == '5' && rightNum == '0') {
					newAns := numLen - j - 2
					if newAns < ans {
						ans = newAns
					}
				}
			}
		}
	}

	return ans
}
