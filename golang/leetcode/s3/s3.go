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

func lengthOfLongestSubstring_250ms(s string) int {
	// 雙指針比對
	passMaxLen := 0
	l := 0
	r := 0

	for r < len(s) {
		r = l + passMaxLen
		if r >= len(s) {
			break
		}

		skip := false
		// 設定已出現字元的hashtable 值為 index
		logMap := make(map[byte]int)
		for i := l; i <= r; i++ {
			if lastIdx, exists := logMap[s[i]]; exists {
				// 發生重複則拋棄前段
				l = lastIdx + 1
				skip = true
				break
			}

			logMap[s[i]] = i
		}
		if skip {
			continue
		}

		//嘗試擴張r
		r++
		passMaxLen++
	}

	if len(s) == 0 {
		return 0
	}
	return passMaxLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	// O(n)單次巡覽 紀錄每個字元上次出現位置 碰到重複時計算距離-1為最大長度

	// abcca
	// 1
	//  2
	//   3
	//    1
	//     2

	// abcba
	// 1
	//  2
	//   3
	//    2
	//     3

	// pwwkew
	// 1
	//  2
	//   1
	//    2
	//     3
	//      1

	// 新字元時streamSize+1
	// 舊字元時 距離與streamSize+1取較短的

	streamSize := 0
	bestLen := 0

	// 設定已出現字元的hashtable 值為 index
	logMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if lastIdx, exists := logMap[s[i]]; exists {
			streamSize = min(i-lastIdx, streamSize+1)
		} else {
			streamSize++
		}
		bestLen = max(bestLen, streamSize)
		logMap[s[i]] = i
	}

	return bestLen
}
