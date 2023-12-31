package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]] \n", permute([]int{1, 2, 3}))
	fmt.Printf("%v -> [[1,2],[2,1]] \n", permute([]int{1, 2}))
}

/*
1 2 3

1 2 3
1	3 2

2

3
*/
func permute(nums []int) [][]int {
	n := len(nums)

	return recursion(nums, n)
}

func recursion(nums []int, n int) [][]int {
	ans := [][]int{}

	if n == 1 {
		ans = append(ans, []int{nums[0]})
		return ans
	}

	for i := 0; i < n; i++ {
		newNums := make([]int, n)
		copy(newNums, nums)

		newNums[i], newNums = newNums[n-1], newNums[:n-1]
		recArray := recursion(newNums, n-1)
		recLen := len(recArray)
		for j := 0; j < recLen; j++ {
			var newRow []int
			newRow = append(newRow, recArray[j]...)
			newRow = append(newRow, nums[i])

			ans = append(ans, newRow)
		}
	}

	return ans
}
