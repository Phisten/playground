package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v ->  \n", isValid("()[[]{}"))
}

// 看起來用堆疊處理就解了
func isValid(s string) bool {
	stack := []rune{} // 使用整數切片來實現堆疊

	closeMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, value := range s {
		if value == '[' || value == '{' || value == '(' {
			// 推入元素到堆疊
			stack = append(stack, value)
		} else {
			// 彈出元素
			if len(stack) == 0 {
				return false
			}

			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if closeMap[popped] != value {
				return false
			}
		}
	}

	return len(stack) == 0
}
