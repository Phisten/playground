package main

import "fmt"

func main() {
	// fmt.Printf("%v -> [a,b,c,d] \n", getWordsInLongestSubsequence(7, []string{"bab", "dab", "ceb", "a", "b", "c", "d"}, []int{1, 2, 2, 1, 2, 3, 4}))
	// fmt.Printf("%v -> [aaa,ada] \n", getWordsInLongestSubsequence(3, []string{"bdb", "aaa", "ada"}, []int{2, 1, 3}))
	fmt.Printf("%v -> [dc,dd,da] \n", getWordsInLongestSubsequence(9, []string{"bad", "dc", "bc", "ccd", "dd", "da", "cad", "dba", "aba"}, []int{9, 7, 1, 2, 6, 8, 3, 7, 2}))
}

// 改成dp由前往後掃描可接續數量
func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	ans := []string{}

	// check hamming dist and groups , if hamming[a][b]==1, 可拼接
	hamming := make([][]int, n)
	for i := 0; i < n; i++ {
		hamming[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dist := 0
			if len(words[i]) == len(words[j]) && groups[i] != groups[j] {
				for c := range words[i] {
					if words[i][c] != words[j][c] {
						dist++
					}
				}
			}
			hamming[i][j] = dist
		}
		// fmt.Printf("%v\t %v\n", words[i], hamming[i])
	}

	// 累進串連數
	dp := make([]int, n)
	// 最長串連序列
	dpList := make([][]string, n)
	for i := range dp {
		dp[i] = 1
		dpList[i] = []string{words[i]}
	}

	for i := 0; i < n; i++ {
		curList := make([]string, len(dpList[i]))
		copy(curList, dpList[i])

		for j := i + 1; j < n; j++ {
			if hamming[i][j] == 1 {
				if dp[j] < dp[i]+1 {
					dp[j] = dp[i] + 1
					dpList[j] = curList
					dpList[j] = append(dpList[j], words[j])
				}
			}
		}
	}

	for _, v := range dpList {
		if len(v) > len(ans) {
			ans = v
		}
	}

	return ans
}

// ----------------

// 要考量nums內的字串有多種長度
// 掃描過程把不同長度的分開算
// 字串還要進行比對 長度必須一樣且只能差一個字
// c2第一個元素開始抓一定是對的 這題必須考量字串間距離
// 組合會跳索引 總長度短 可以用遞迴解
// x 未完成解題
func getWordsInLongestSubsequence_x(n int, words []string, groups []int) []string {

	// ansByLen := [11][]string{}

	wordsByLen := [11][]string{}
	groupsByLen := [11][]int{}

	bestAns := []string{}
	// bestSubLen := 0

	// group by len
	for i, v := range groups {
		l := len(words[i])
		wordsByLen[l] = append(wordsByLen[l], words[i])
		groupsByLen[l] = append(groupsByLen[l], v)
	}
	hammingByLen := [11][][]int{}
	// hamming dist
	for i := range wordsByLen {
		l := len(wordsByLen[i])
		hammingByLen[i] = make([][]int, l)
		for j := 0; j < l; j++ {
			hammingByLen[i][j] = make([]int, l)
			for k := j + 1; k < l; k++ {
				dist := 0
				for c := 0; c < i; c++ {
					if wordsByLen[i][j][c] != wordsByLen[i][k][c] {
						dist++
					}
				}
				hammingByLen[i][j][k] = dist
			}
		}
	}

	fmt.Printf("%v \n", hammingByLen)
	fmt.Printf("%v \n", wordsByLen)
	fmt.Printf("%v \n", groupsByLen)

	for l := 1; l <= 10; l++ {
		curAns := recursion(groupsByLen[l], wordsByLen[l], hammingByLen[l], 0)
		if len(curAns) > len(bestAns) {
			bestAns = curAns
		}

		// curN := len(groupsByLen[l])
		// for start := 0; start < curN; start++ {
		// 	tempAnsByLen := []string{}
		// 	lastAppendIdx := -1
		// 	for i := start; i < curN; i++ {
		// 		if i == start {
		// 			tempAnsByLen = append(tempAnsByLen, wordsByLen[l][i])
		// 			lastAppendIdx = i
		// 			continue
		// 		}

		// 		if groupsByLen[l][i] != groupsByLen[l][lastAppendIdx] && hammingByLen[l][lastAppendIdx][i] == 1 {
		// 			tempAnsByLen = append(tempAnsByLen, wordsByLen[l][i])
		// 			lastAppendIdx = i
		// 		}
		// 	}

		// 	// if len is better, append to ans
		// 	if len(tempAnsByLen) > len(bestAns) {
		// 		bestAns = tempAnsByLen
		// 	}
		// }
	}

	return bestAns
}

func recursion(groupsByLenL []int, wordsByLenL []string, hammingByLenL [][]int, startIdx int) []string {
	bestAns := []string{}
	curN := len(groupsByLenL)
	for start := startIdx; start < curN; start++ {
		tempAnsByLen := []string{}
		// lastAppendIdx := -1
		for i := start; i < curN; i++ {
			curAns := recursion(groupsByLenL, wordsByLenL, hammingByLenL, i)

			if len(curAns) > len(tempAnsByLen) {
				tempAnsByLen = curAns
			}
		}

		// if i == start {
		// 	tempAnsByLen = append(tempAnsByLen, wordsByLenL[i])
		// 	lastAppendIdx = i
		// 	continue
		// }

		// if groupsByLenL[i] != groupsByLenL[lastAppendIdx] && hammingByLenL[lastAppendIdx][i] == 1 {
		// 	tempAnsByLen = append(tempAnsByLen, wordsByLenL[i])
		// 	lastAppendIdx = i
		// }

		// if len is better, append to ans
		if len(tempAnsByLen) > len(bestAns) {
			bestAns = tempAnsByLen
		}
	}
	return bestAns
}
