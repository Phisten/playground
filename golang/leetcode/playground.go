package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> 0 \n", someFunc([]int{2, 7, 11, 15}, 9))

	my_arr1 := []string{"Apple", "Mango", "Banana"}
	my_arr2 := append(my_arr1, "123")
	arr3 := [][]string{my_arr1, my_arr2, my_arr2, append(my_arr2, "1234")}
	arr3[2][0] = "11111"
	fmt.Println("The first array, arr1 is:", my_arr1)
	fmt.Println("The array obta   arr2 is:", my_arr2)
	fmt.Println("The array obta   arr2 is:", arr3)
}

func someFunc(nums []int, target int) []int {
	return nil
}
