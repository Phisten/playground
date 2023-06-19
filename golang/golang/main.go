package main

import (
	"fmt"
	"math/rand"
)

func main() {
	helloList := []string{"Hello", "你好", "早安"}

	r := rand.Intn(3)

	fmt.Println(len(helloList))
	fmt.Println(r)
	fmt.Println(helloList[r])
}
