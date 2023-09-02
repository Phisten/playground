package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> true \n", checkStrings("abcdba", "cabdab"))
	fmt.Printf("%v -> f \n", checkStrings("abe", "bea"))
	fmt.Printf("%v -> f \n", checkStrings("abe", "bea"))
}

func checkStrings(s1 string, s2 string) bool {
	hashOdd := map[byte]int{}
	hashEven := map[byte]int{}

	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			hashEven[s1[i]]++
			hashEven[s2[i]]--
		} else {
			hashOdd[s1[i]]++
			hashOdd[s2[i]]--
		}
	}

	ansE := true
	ansO := true
	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			ansE = ansE && hashEven[s1[i]] == 0 && hashEven[s2[i]] == 0
		} else {
			ansO = ansO && hashOdd[s1[i]] == 0 && hashOdd[s2[i]] == 0
		}
	}

	return ansE && ansO
}
