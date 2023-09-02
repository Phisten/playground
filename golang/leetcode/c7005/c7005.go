package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v -> true \n", checkStrings("abcdba", "cabdab"))
	fmt.Printf("%v -> f \n", checkStrings("abe", "bea"))
	fmt.Printf("%v -> f \n", checkStrings("abe", "bea"))
}
