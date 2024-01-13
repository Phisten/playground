package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(1))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return version >= 0
}

// 要找到T和F的邊界, 輸出T (第一個壞的版本)
// 二分搜尋先找T在哪一側
// 發現後繼續二分搜尋直到
func firstBadVersion(n int) int {
	l := 0
	r := n - 1
	for {
		s := (l + r) / 2

		result := isBadVersion(s + 1)
		if result {
			r = s
		} else {
			l = s + 1
		}

		if l == r {
			return l + 1
		}
	}
}
