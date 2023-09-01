package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> [0,1,1] \n", countBits(2))
	fmt.Printf("%v -> [0,1,1] \n", countBits(1000))
}

/*
0   0   0   0
1   1   1   1
2   10  2   1 2^1
3   11  3   2
4   100 4   1 2^2
5   101 5   2
6   110 6   2
7   111 7   3
8   1000    1 2^3
9   1001    2
10  1010    2
11  1011    3
12  1100    2
13  1101    3
14  1110    3
15  1111    4
16  10000   1 2^4
*/
func countBits(n int) []int {
	ans := []int{0}

	binaryStep := 1
	for i := 1; i <= n; i++ {
		if i == binaryStep*2 {
			binaryStep = i
			ans = append(ans, 1)
		} else {
			ans = append(ans, ans[i-binaryStep]+1)
		}
	}

	return ans

}
