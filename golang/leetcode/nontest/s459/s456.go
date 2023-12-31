// https://leetcode.com/problems/4sum/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern_logicRule("aabaaba"))
	fmt.Println(repeatedSubstringPattern_logicRule("aba"))
	fmt.Println(repeatedSubstringPattern_logicRule("aaa"))
	fmt.Println(repeatedSubstringPattern_logicRule("abcabcabcabc"))
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

func repeatedSubstringPattern_logicRule(s string) bool {
	ss := s + s

	return strings.Contains(ss[1:len(ss)-1], s)
}
