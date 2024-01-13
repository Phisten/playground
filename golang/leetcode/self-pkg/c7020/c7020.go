package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v -> 4 \n", countSymmetricIntegers(1200, 1230))
	fmt.Printf("%v -> 9 \n", countSymmetricIntegers(1, 100))
}

func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for i := low; i <= high; i++ {

		strNum := strconv.Itoa(i)

		charCount := len(strNum)
		if charCount%2 == 0 {
			left := 0
			right := 0

			for idx, char := range strNum {
				chOpt, _ := strconv.Atoi(string(char))
				if idx < charCount/2 {
					left += chOpt
				} else {
					right += chOpt
				}
			}
			if left == right {
				ans++
			}
		}
	}

	return ans
}
