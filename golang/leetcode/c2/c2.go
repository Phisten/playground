package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> [a,b,c] \n", getWordsInLongestSubsequence(4, []string{"a", "b", "c", "d"}, []int{1, 0, 1, 1}))
	fmt.Printf("%v -> [a,b,c] \n", getWordsInLongestSubsequence(4, []string{"a", "b", "c", "d"}, []int{1, 1, 1, 0}))
}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	ans := []string{}
	for i, v := range groups {
		if i == 0 {
			ans = append(ans, words[0])
			continue
		}

		if v != groups[i-1] {
			ans = append(ans, words[i])

		}
	}

	return ans
}
