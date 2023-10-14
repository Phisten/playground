package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v -> [2,1,-1] \n", lastVisitedIntegers([]string{"1", "2", "prev", "prev", "prev"}))
}

func lastVisitedIntegers(words []string) []int {

	ans := []int{}

	ints := []int{}

	prevStream := 0

	for _, v := range words {
		if v != "prev" {
			marks, _ := strconv.Atoi(v)
			ints = append(ints, marks)
			prevStream = 0
		} else {
			prevStream++
			n := len(ints)
			if n >= prevStream {
				ans = append(ans, ints[len(ints)-prevStream])
			} else {
				ans = append(ans, -1)
			}
		}
	}

	return ans
}
