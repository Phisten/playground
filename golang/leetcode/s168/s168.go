package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(53))
	fmt.Println(convertToTitle(701))

}
func convertToTitle(columnNumber int) string {
	colMapping := "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	//             01234567890123456789012345

	str := ""

	for i := columnNumber; i > 0; i = (i - 1) / 26 {
		curCharIndex := i % 26
		str = colMapping[curCharIndex:curCharIndex+1] + str

		if i <= 26 {
			break
		}
	}

	return str
}
