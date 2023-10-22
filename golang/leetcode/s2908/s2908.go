package main

import (
	"fmt"
)

func main() {
	// fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}), "13")
	// fmt.Println(minimumSum([]int{1, 1, 2, 1}), "4")
	// fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}), "9")
	fmt.Println(minimumSum([]int{1, 2, 1, 2}), "4")

}

// 迴圈暴力解 TLE
func minimumSum(nums []int) int {
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

// 左右找最小數
// 邏輯太複雜, 放棄
func minimumSum_xx(nums []int) int {
	ans := 2147483647
	n := len(nums)

	idxL := 0
	idxR := n - 1

	ans = minimumSumReg(nums, idxL, idxR, idxR)

	if ans == 2147483647 {
		ans = -1
	}

	return ans
}

func minimumSumReg(nums []int, idxL int, idxR int, idxM int) int {
	ans := 2147483647

	valL := nums[idxL]
	valR := nums[idxR]

	if idxM == idxL || idxM == idxR {
		newValM := 2147483647
		for i := idxL + 1; i < idxR; i++ {
			cur := nums[i]
			if cur < newValM && cur > valL && cur > valR {
				newValM = cur
				idxM = i
			}
		}
	}

	valM := nums[idxM]

	if valM > valL && valM > valR {
		sum := valL + valM + valR
		if sum < ans {
			ans = sum
		}
	}

	newIdxL := -1
	for i := idxL + 1; i <= idxM; i++ {
		if nums[i] < valL {
			newIdxL = i
			break
		}
	}
	if newIdxL > -1 {
		ansL := 2147483647
		ansL = minimumSumReg(nums, newIdxL, idxR, idxM)
		if ansL < ans {
			ans = ansL
		}
	}

	newIdxR := -1
	for i := idxR - 1; i >= idxM; i-- {
		if nums[i] < valR {
			newIdxR = i
			break
		}
	}
	if newIdxR > -1 {
		ansR := 2147483647
		ansR = minimumSumReg(nums, idxL, newIdxR, idxM)
		if ansR < ans {
			ans = ansR
		}
	}

	return ans
}

// c1c2可以一起解, 只差在量級
// c2沒超過2^31 不用處理大數
// kernel window 掃瞄後取非0最小值
// -搞錯題意, 組合可以跳號 不能用window掃瞄
func minimumSum_x(nums []int) int {
	ans := 2147483647
	n := len(nums)

	for i := 1; i < n-1; i++ {
		cur := nums[i] + nums[i-1] + nums[i+1]
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			if ans > cur {
				ans = cur
			}
		}
	}

	if ans == 2147483647 {
		ans = -1
	}

	return ans
}
