// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("aab"))
}

// 暴力解 O(N^2)：遍歷每個字元作為起點 看可以往後爬多長不碰到重複

func lengthOfLongestSubstring_701ms(s string) int {
	// cache 各個起點字元的各個連續長度
	// 以起點index為索引的array 儲存目前長度中是否已有重複字元
	dead := make([]bool, len(s)+1)
	maxLen := 1

	for subLen := 1; subLen <= len(s); subLen++ {

		if subLen == 1 {
			continue
		}
		for startIdx := 0; startIdx < len(s)-subLen+1; startIdx++ {

			lastIdx := startIdx + subLen - 1
			if !dead[startIdx] && s[lastIdx:lastIdx+1] != s[startIdx:startIdx+1] && !dead[startIdx+1] {
				maxLen = subLen
			} else {
				dead[startIdx] = true
				continue
			}
		}
	}

	//edge case
	if len(s) == 0 {
		return 0
	}
	return maxLen
}
