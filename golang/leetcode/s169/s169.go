package main

import "fmt"

func main() {
	moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
	fmt.Printf(" -> %v\n", "[4,2,4,3,5,1,0,0,0,0]")

	moveZeroes([]int{1, 2, 0, 0, 3, 0, 4, 5, 0})
	moveZeroes([]int{0, 0, 1, 2, 3, 0, 0, 0})
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 1, 0, 3, 0})
	moveZeroes([]int{0, 1, 2, 3, 0})

}
func moveZeroes(nums []int) {
	pickIndex := 0
	hasZero := false

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			hasZero = true
		}
		if hasZero {
			if i > pickIndex {
				pickIndex = i + 1
			} else {
				pickIndex++
			}

			for ; ; pickIndex++ {
				if pickIndex >= len(nums) {
					nums[i] = 0
					break
				}
				if nums[pickIndex] != 0 {
					nums[i] = nums[pickIndex]
					break
				}
			}
		}

	}
	fmt.Printf("%v\n", nums)
}
