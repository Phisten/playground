package leetcode

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("pow 3^2=%v  \n", math.Pow(3, 2))

	// 指定長度
	a := [2]int{1, 2}
	b := a

	b[0] = 3

	fmt.Printf("a=%v b=%v \n", a, b)

	// 不指定長度
	a2 := []int{1, 2}
	b2 := a2

	b2[0] = 3

	fmt.Printf("a2=%v b2=%v \n", a2, b2)
}
