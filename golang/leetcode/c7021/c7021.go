package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> true \n", canBeEqual("abcd", "cdab"))
	fmt.Printf("%v -> f \n", canBeEqual("abcd", "dacb"))
}

// 將所有字幕存進hashmap 檢查數量是否一致
func canBeEqual(s1 string, s2 string) bool {
	ans := 0

	pass := [4]bool{}
	for i := 0; i < 4; i++ {
		pass[i] = s1[i] == s2[i]
		if pass[i] {
			ans++
		}
	}

	if !pass[0] && !pass[2] {
		if s1[0] == s2[2] && s1[2] == s2[0] {
			ans += 2
		}
	}

	if !pass[1] && !pass[3] {
		if s1[1] == s2[3] && s1[3] == s2[1] {
			ans += 2
		}
	}

	return ans == 4
}
