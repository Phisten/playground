// https://leetcode.com/problems/4sum/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(repeatedSubstringPattern("aabaaba"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("aaa"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {
	cat := ""

	for i := 0; i < len(s)/2; i++ {
		cat = s[:i+1]
		for j := i + 1; j < len(s); j++ {
			if cat[j%len(cat)] != s[j] {
				break
			} else if j == len(s)-1 && (j+1)%len(cat) == 0 {
				return true
			}
		}
	}

	return false
}
