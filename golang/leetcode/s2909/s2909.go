package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}), "13")
	fmt.Println(minimumSum([]int{1, 1, 2, 1}), "4")
	fmt.Println(minimumSum([]int{1, 2, 1}), "4")
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}), "9")
	fmt.Println(minimumSum([]int{1, 2, 1, 2}), "4")
}

// 最初左右指針想法是正確的
// 但預處理的資料錯了, 應儲存往左,往右的最小值
func minimumSum(nums []int) int {

	ans := 2147483647
	n := len(nums)

	dpL := make([]int, n)
	dpR := make([]int, n)

	dpL[0] = nums[0]
	for i := 1; i < n; i++ {
		dpL[i] = min(nums[i], dpL[i-1])
	}

	dpR[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		dpR[i] = min(nums[i], dpR[i+1])
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > dpL[i] && nums[i] > dpR[i] {
			cur := nums[i] + dpL[i] + dpR[i]
			ans = min(ans, cur)
		}
	}

	if ans == 2147483647 {
		ans = -1
	}
	return ans
}

func min(l int, r int) int {
	if l < r {
		return l
	}
	return r
}

// 還是TLE
func minimumSum_xx(nums []int) int {
	ans := 2147483647
	n := len(nums)

	dpL := make([]int, n)
	dpR := make([]int, n)

	// for i := 0; i < n; i++ {
	// 	dpL[i] = 0
	// 	dpR[i] = 0
	// }

	for midI := 1; midI < n-1; midI++ {
		mid := nums[midI]
		for l := 0; l < midI; l++ {
			if nums[l] < mid {
				if dpL[midI] == 0 || nums[l] < dpL[midI] {
					dpL[midI] = nums[l]
				}
			}
		}
		for r := midI + 1; r < n; r++ {
			if nums[r] < mid {
				if dpR[midI] == 0 || nums[r] < dpR[midI] {
					dpR[midI] = nums[r]
				}
			}
		}
	}

	// fmt.Println(dpL)
	// fmt.Println(dpR)

	// mid i
	for i := 1; i < n-1; i++ {
		if dpL[i] > 0 && dpR[i] > 0 {
			cur := dpL[i] + dpR[i] + nums[i]
			if cur < ans {
				ans = cur
			}
		}
	}

	if ans == 2147483647 {
		ans = -1
	}
	return ans
}

// 迴圈暴力解 TLE
func minimumSum_x(nums []int) int {
	ans := 2147483647
	n := len(nums)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				valL := nums[i] + nums[j]
				for k := j + 1; k < n; k++ {
					if nums[j] > nums[k] {
						cur := valL + nums[k]
						if cur < ans {
							ans = cur
						}
					}
				}
			}
		}
	}

	if ans == 2147483647 {
		ans = -1
	}
	return ans
}
