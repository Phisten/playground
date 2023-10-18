package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("   -3.14159"))
	fmt.Println(myAtoi("   -42word"))
	fmt.Println(myAtoi(" 001 "))
	fmt.Println(myAtoi(" -5- "))
	fmt.Println(myAtoi("   2147483648"))
	fmt.Println(myAtoi("   -2147483649"))
}

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	allowStr := "0123456789"
	startAllowStr := "+-"

	end := -1
	for i, v := range s {
		if strings.ContainsRune(allowStr, v) {
		} else if i == 0 && strings.ContainsRune(startAllowStr, v) {
		} else {
			break
		}
		end = i
	}
	if end >= 0 {
		s = s[:end+1]
	}

	i32, _ := strconv.Atoi(s)
	if i32 > 2147483647 {
		return 2147483647
	} else if i32 < -2147483648 {
		return -2147483648
	} else {
		return i32
	}
}
