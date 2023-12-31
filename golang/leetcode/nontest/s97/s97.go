package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> T \n", isInterleave("", "b", "b"))
	fmt.Printf("%v -> T \n", isInterleave("a", "b", "ab"))
	fmt.Printf("%v -> T \n", isInterleave("aa", "bb", "abab"))

	fmt.Printf("%v -> T \n", isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Printf("%v -> F \n", isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}

// *空間要求：  O(s2.length)

func isInterleave(s1 string, s2 string, s3 string) bool {

	// len check
	len1 := len(s1)
	len2 := len(s2)
	len3 := len(s3)
	fmt.Printf("s1:%v, s2:%v, s3:%v\n", s1, s2, s3)

	if len1+len2 != len3 {
		return false
	}

	dp := make([][]bool, len1+1)
	dp[0] = make([]bool, len2+1)
	dp[0][0] = true
	// init and base case
	for i := 0; i < len1; i++ {
		dp[i+1] = make([]bool, len2+1)
	}
	for i := 0; i < len1; i++ {
		if dp[i][0] && s1[i] == s3[i] {
			dp[i+1][0] = true
		}
	}
	for j := 0; j < len2; j++ {
		if dp[0][j] && s2[j] == s3[j] {
			dp[0][j+1] = true
		}
	}

	optStr := map[bool]string{
		false: "0",
		true:  "1",
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			dp[i+1][j+1] = (dp[i][j+1] && s3[i+j+1] == s1[i]) || (dp[i+1][j] && s3[i+j+1] == s2[j])
			fmt.Printf("%v ", optStr[dp[i+1][j+1]])
		}
		fmt.Printf("\n")
	}

	return dp[len1][len2]
}
