package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> T \n", canConstruct("aa", "aab"))
}

// hashtable紀錄字母出現次數, ransomNote逐字扣掉table存量
func canConstruct(ransomNote string, magazine string) bool {
	magHT := map[rune]int{}
	for _, v := range magazine {
		magHT[v] += 1
	}
	for _, v := range ransomNote {
		if magHT[v] == 0 {
			return false
		}
		magHT[v] -= 1
	}
	return true
}
